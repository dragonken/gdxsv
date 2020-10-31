package main

import (
	"fmt"
)

const (
	PingLimitTh = 50
)

var (
	lbsLobbySettings map[uint16]*LobbySetting
)

// PS2 LobbyID: 1-23
// DC2 LobbyID: 2, 4-6, 9-17, 19-22

type LobbySetting struct {
	Name                string
	McsRegion           string
	Comment             string
	EnableExtraCostCmd  bool
	EnableForceStartCmd bool
	TeamShuffle         bool
	PingLimit           bool
}

func init() {
	lbsLobbySettings = map[uint16]*LobbySetting{
		// earth lobbies
		1: {
			// PS2 Only
			Name:                "",
			McsRegion:           "",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "1",
		},
		2: {
			Name:                "タクラマカン砂漠",
			McsRegion:           "",
			PingLimit:           true,
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			TeamShuffle:         true,
			Comment:             "Default Server",
		},
		3: {
			// PS2 Only
			Name:                "",
			McsRegion:           "",
			PingLimit:           true,
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "1",
		},
		4: {
			Name:                "黒海南岸森林地帯",
			McsRegion:           "asia-east2",
			PingLimit:           true,
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
		},
		5: {
			Name:        "オデッサ",
			McsRegion:   "asia-east2",
			TeamShuffle: true,
			PingLimit:   true,
			Comment:     "TeamShuffle",
		},
		6: {
			Name:        "ベルファスト",
			McsRegion:   "asia-northeast1",
			TeamShuffle: true,
			PingLimit:   true,
			Comment:     "TeamShuffle",
		},
		7: {
			// PS2 Only
			Name:                "",
			McsRegion:           "",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "1",
		},
		8: {
			// PS2 Only
			Name:                "",
			McsRegion:           "",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "1",
		},
		9: {
			Name:                "ニューヤーク",
			McsRegion:           "asia-northeast1",
			PingLimit:           true,
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
		},
		10: {
			Name:                "グレートキャニオン",
			McsRegion:           "asia-northeast3",
			PingLimit:           true,
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
		},
		11: {
			Name:                "ジャブロー",
			McsRegion:           "asia-northeast2",
			PingLimit:           true,
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
		},
		12: {
			Name:                "地下基地",
			McsRegion:           "asia-east1",
			PingLimit:           true,
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "For JP vs HK",
		},

		// space lobbies
		13: {
			Name:                "ソロモン",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Free Lobby",
		},
		14: {
			Name:                "ソロモン宙域",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Free Lobby",
		},
		15: {
			Name:                "ア・バオア・クー宙域",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Private Room",
		},
		16: {
			Name:                "ア・バオア・クー外部",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Private Room",
		},
		17: {
			Name:                "ア・バオア・クー内部",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Private Room",
		},
		18: {
			// PS2 Only
			Name:                "",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Free Lobby",
		},
		19: {
			Name:                "衛星軌道1",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Free Lobby",
		},
		20: {
			Name:                "衛星軌道2",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Free Lobby",
		},
		21: {
			Name:                "サイド6宙域",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Free Lobby",
		},
		22: {
			Name:                "サイド7内部",
			McsRegion:           "best",
			EnableForceStartCmd: true,
			EnableExtraCostCmd:  true,
			Comment:             "Free Lobby",
		},
	}

	for _, x := range lbsLobbySettings {
		locName, ok := gcpLocationName[x.McsRegion]
		if !ok {
			locName = "Default Server"
		}
		if x.McsRegion == "best" {
			locName = "Best Server [Auto Detection]"
		}
		x.Comment = fmt.Sprintf("<B>%s<B><BR><B>%s<END>", locName, x.Comment)
	}
}
