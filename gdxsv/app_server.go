package main

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/golang/glog"
)

type Server struct {
	app  *App
	conn *net.TCPConn
}

func NewServer(app *App) *Server {
	return &Server{
		app: app,
	}
}

func (s *Server) ListenAndServe(addr string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	listner, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	for {
		tcpConn, err := listner.AcceptTCP()
		if err != nil {
			glog.Errorln(err)
			continue
		}
		glog.Infoln("A new tcp connection open.", tcpConn.RemoteAddr())
		conn := NewConn(tcpConn)
		conn.peer = s.app.NewPeer(conn)
		go conn.serve()
	}
}

type GameConn struct {
	conn *net.TCPConn
	peer *AppPeer

	chWrite    chan bool
	chDispatch chan bool
	chQuit     chan interface{}

	mOutbuf sync.Mutex
	outbuf  []byte

	mInbuf sync.Mutex
	inbuf  []byte
}

func NewConn(conn *net.TCPConn) *GameConn {
	return &GameConn{
		conn:       conn,
		chWrite:    make(chan bool, 1),
		chDispatch: make(chan bool, 1),
		outbuf:     make([]byte, 0, 1024),
		inbuf:      make([]byte, 0, 1024),
	}
}

func (c *GameConn) serve() {
	defer c.conn.Close()
	defer c.peer.OnClose()

	ctx, cancel := context.WithCancel(context.Background())

	go c.dispatchLoop(ctx, cancel)
	go c.writeLoop(ctx, cancel)
	go c.readLoop(ctx, cancel)

	c.peer.OnOpen()
	<-ctx.Done()
}

func (c *GameConn) SendMessage(msg *Message) {
	glog.V(2).Infof("\t->%v %v \n", c.Address(), msg)
	c.mOutbuf.Lock()
	c.outbuf = append(c.outbuf, msg.Serialize()...)
	c.mOutbuf.Unlock()
	select {
	case c.chWrite <- true:
	default:
	}
}

func (c *GameConn) Address() string {
	return c.conn.RemoteAddr().String()
}

func (c *GameConn) readLoop(ctx context.Context, cancel func()) {
	defer cancel()

	buf := make([]byte, 4096)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		c.conn.SetReadDeadline(time.Now().Add(time.Minute * 30))
		n, err := c.conn.Read(buf)
		if err != nil {
			glog.Infoln("TCP conn error:", err)
			return
		}
		if n == 0 {
			glog.Infoln("TCP read zero")
			return
		}
		c.mInbuf.Lock()
		c.inbuf = append(c.inbuf, buf[:n]...)
		c.mInbuf.Unlock()

		select {
		case c.chDispatch <- true:
		default:
		}
	}
}

func (c *GameConn) writeLoop(ctx context.Context, cancel func()) {
	defer cancel()

	buf := make([]byte, 0, 128)
	for {
		select {
		case <-ctx.Done():
			return
		case <-c.chWrite:
			c.mOutbuf.Lock()
			if len(c.outbuf) == 0 {
				c.mOutbuf.Unlock()
				continue
			}
			buf = append(buf, c.outbuf...)
			c.outbuf = c.outbuf[:0]
			c.mOutbuf.Unlock()

			sum := 0
			size := len(buf)
			for sum < size {
				c.conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
				n, err := c.conn.Write(buf[sum:])
				if err != nil {
					glog.Errorf("%v write error: %v\n", c.Address(), err)
					break
				}
				sum += n
			}
			buf = buf[:0]
		}
	}
}

func (c *GameConn) dispatchLoop(ctx context.Context, cancel func()) {
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.chDispatch:
			c.mInbuf.Lock()
			for len(c.inbuf) >= HeaderSize {
				n, msg := Deserialize(c.inbuf)
				if n == 0 {
					// not enough data comming
					break
				}

				c.inbuf = c.inbuf[n:]
				if msg != nil {
					glog.V(2).Infof("%v %v\n", c.Address(), msg)
					c.peer.OnMessage(msg)
				}
			}
			c.mInbuf.Unlock()
		}
	}
}
