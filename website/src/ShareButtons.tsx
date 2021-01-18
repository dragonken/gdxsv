import React from 'react';
import { useIntl } from 'react-intl';
import { WebsiteUrl } from './Const';

interface ILink {
    href: string;
    ariaLabel: string;
    children: React.ReactNode;
}

const Link = ({ href, ariaLabel, children }: ILink) => (
    <a
        href={href}
        target="_blank"
        aria-label={ariaLabel}
        rel="noopener noreferrer nofollow"
        style={{ margin: '0 0.2rem' }}
    >
        {children}
    </a>
);

const TwitterButton = () => {
    const intl = useIntl();
    const text = intl.formatMessage({ id: "footer.tweet.data-text" });
    const url = escape(WebsiteUrl);
    const hashTags = ([
        'gdxsv',
        '連ジ',
        'ガンダム',
        'Gundam',
    ].join(','));
    const href = `https://twitter.com/share?text=${text}&url=${url}&hashtags=${hashTags}`;

    return (
        <Link href={href} ariaLabel="Share on Twitter">
            <svg style={{ width: '2rem' }} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M12 0C5.38 0 0 5.38 0 12s5.38 12 12 12 12-5.38 12-12S18.62 0 12 0zm5.26 9.38v.34c0 3.48-2.64 7.5-7.48 7.5-1.48 0-2.87-.44-4.03-1.2 1.37.17 2.77-.2 3.9-1.08-1.16-.02-2.13-.78-2.46-1.83.38.1.8.07 1.17-.03-1.2-.24-2.1-1.3-2.1-2.58v-.05c.35.2.75.32 1.18.33-.7-.47-1.17-1.28-1.17-2.2 0-.47.13-.92.36-1.3C7.94 8.85 9.88 9.9 12.06 10c-.04-.2-.06-.4-.06-.6 0-1.46 1.18-2.63 2.63-2.63.76 0 1.44.3 1.92.82.6-.12 1.95-.27 1.95-.27-.35.53-.72 1.66-1.24 2.04z" /></svg>
        </Link>
    );
};

const FacebookButton = () => {
    const intl = useIntl();
    const href = `https://facebook.com/sharer/sharer.php?u=${escape(WebsiteUrl)}`;
    return (
        <Link href={href} ariaLabel={intl.formatMessage({ id: "footer.facebook.title" })}>
            <svg style={{ width: '2rem' }} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M12 0C5.38 0 0 5.38 0 12s5.38 12 12 12 12-5.38 12-12S18.62 0 12 0zm3.6 11.5h-2.1v7h-3v-7h-2v-2h2V8.34c0-1.1.35-2.82 2.65-2.82h2.35v2.3h-1.4c-.25 0-.6.13-.6.66V9.5h2.34l-.24 2z" /></svg>
        </Link>
    );
};

const MeweButton = () => {
    const intl = useIntl();
    const text = encodeURI(intl.formatMessage({ id: "footer.tweet.data-text" }));
    const url = escape(WebsiteUrl);
    const href = `https://mewe.com/share?link=${url}%20${text}`;
    return (
        <Link href={href} ariaLabel="Share on Twitter">
            <svg style={{ width: '2rem' }} xmlns="http://www.w3.org/2000/svg" viewBox="-3 -3 30 30">
                <clipPath id="circle">
                    <circle cx="12" cy="12" r="15" />
                </clipPath>
                <circle cx="12" cy="12" r="15" fill="black" />
                <path id="mewe" d="M3.029 7.712a.768.768 0 100 1.536.768.768 0 000-1.536zm7.883 0a.768.768 0 100 1.536.768.768 0 000-1.536zm4.414 0a.768.768 0 100 1.536.768.768 0 000-1.536zm4.418 0a.768.768 0 100 1.537.768.768 0 000-1.537zm-4.447 2.307c-.27 0-.453.175-.54.428l-1.378 4.175-1.403-4.193a.546.546 0 00-.532-.392c-.296 0-.55.227-.55.523 0 .07.018.148.044.21l1.805 5.056c.105.279.296.444.558.444h.105c.261 0 .461-.165.557-.444l1.36-3.985 1.36 3.985c.096.279.288.444.55.444h.105c.261 0 .462-.174.558-.444l1.804-5.057a.617.617 0 00.052-.218.525.525 0 00-.531-.514c-.262 0-.445.166-.523.383l-1.404 4.202-1.377-4.175c-.079-.253-.27-.428-.541-.428zM.54 10.045a.538.538 0 00-.54.541v5.117c0 .297.227.523.523.523a.52.52 0 00.532-.523V12.05l1.482 2.224c.113.174.253.279.444.279.201 0 .34-.105.454-.28l1.49-2.24v3.661a.54.54 0 00.533.532.53.53 0 00.54-.532v-5.108a.538.538 0 00-.54-.54h-.114a.54.54 0 00-.488.278L3 13.227 1.15 10.333a.565.565 0 00-.497-.288zm8.368 1.439c-1.316 0-2.25 1.081-2.25 2.397v.018c0 1.42 1.03 2.389 2.363 2.389.715 0 1.239-.236 1.657-.61a.46.46 0 00.14-.323.415.415 0 00-.419-.427.391.391 0 00-.27.096 1.655 1.655 0 01-1.09.4c-.68 0-1.212-.418-1.325-1.168h2.885c.27 0 .497-.209.497-.505 0-1.064-.723-2.267-2.188-2.267zm12.904 0c-1.317 0-2.25 1.081-2.25 2.397v.018c0 1.42 1.029 2.389 2.363 2.389.715 0 1.238-.236 1.657-.61a.46.46 0 00.139-.323.415.415 0 00-.418-.427.392.392 0 00-.271.096 1.652 1.652 0 01-1.09.4c-.68 0-1.211-.418-1.325-1.168h2.886c.27 0 .497-.209.497-.505 0-1.064-.724-2.267-2.188-2.267zm-12.913.863c.698 0 1.099.532 1.169 1.212H7.705c.096-.715.549-1.212 1.194-1.212zm12.904 0c.697 0 1.098.532 1.168 1.212h-2.363c.096-.715.55-1.212 1.195-1.212z" />
                <use clip-path="url(#circle)" xlinkHref="#mewe" fill="white" />
            </svg>
        </Link>
    );
};

const LineButton = () => {
    const url = escape(WebsiteUrl);
    const href = `line://msg/text/${url}`;
    return (
        <Link href={href} ariaLabel="Share on LINE">
            <svg style={{ width: '2rem' }} xmlns="http://www.w3.org/2000/svg" viewBox="0 -6 128 128">
                <clipPath id="circle">
                    <circle cx="64" cy="58" r="64" />
                </clipPath>
                <circle cx="64" cy="58" r="64" fill="black" />
                <path id="line" xmlns="http://www.w3.org/2000/svg" d="M 103.699997 70.959999 C 104.324051 70.965523 104.834427 70.464104 104.839951 69.840042 C 104.839981 69.836693 104.839996 69.833344 104.839996 69.829994 L 104.839996 65.599998 C 104.834511 64.974304 104.325714 64.469971 103.699997 64.470001 L 92.360001 64.470001 L 92.360001 60.070000 L 103.720001 60.070000 C 104.340195 60.070023 104.844513 59.570171 104.849998 58.950001 L 104.849998 54.720001 C 104.844513 54.094307 104.335716 53.589977 103.709999 53.590000 L 92.360001 53.590000 L 92.360001 49.209999 L 103.720001 49.209999 C 104.344086 49.209999 104.849998 48.704082 104.849998 48.079998 L 104.849998 43.849998 C 104.849998 43.220394 104.339600 42.709999 103.709999 42.709999 L 87.000000 42.709999 C 86.370392 42.709999 85.860001 43.220394 85.860001 43.849998 L 85.860001 69.820000 C 85.865486 70.445694 86.374283 70.950027 87.000000 70.949997 L 103.709999 70.949997 Z M 41.880001 70.959999 C 42.505722 70.960022 43.014511 70.455696 43.020000 69.830002 L 43.020000 65.599998 C 43.014511 64.974304 42.505722 64.469971 41.880001 64.470001 L 30.520000 64.470001 L 30.520000 43.860001 C 30.520000 43.230396 30.009605 42.720001 29.380001 42.720001 L 25.160000 42.720001 C 24.534304 42.725491 24.029976 43.234280 24.030001 43.860001 L 24.030001 69.830002 C 24.035431 70.451820 24.538183 70.954575 25.160000 70.959999 L 41.880001 70.959999 M 51.939999 42.730000 L 47.720001 42.730000 C 47.098183 42.735432 46.595432 43.238182 46.590000 43.860001 L 46.590000 69.830002 C 46.595432 70.451820 47.098183 70.954575 47.720001 70.959999 L 51.939999 70.959999 C 52.565720 70.960022 53.074509 70.455696 53.079998 69.830002 L 53.080002 43.860001 C 53.080002 43.230396 52.569607 42.720001 51.940002 42.720001 M 80.669998 42.730000 L 76.440002 42.730000 C 75.818184 42.735432 75.315430 43.238182 75.310005 43.860001 L 75.309998 59.279999 L 63.430000 43.240002 C 63.398373 43.198612 63.365009 43.158577 63.330002 43.120003 C 63.309639 43.090878 63.286167 43.064053 63.260002 43.040001 L 63.240002 43.020000 L 63.180000 42.980000 L 63.150002 42.950001 C 63.130768 42.935551 63.110737 42.922195 63.090000 42.910000 C 63.078320 42.897964 63.064823 42.887840 63.049999 42.880001 L 62.990002 42.860001 L 62.950001 42.840000 L 62.889999 42.799999 C 62.876759 42.802238 62.863239 42.802238 62.849998 42.799999 C 62.830002 42.779999 62.810001 42.779999 62.779999 42.770000 L 62.740002 42.759998 C 62.740002 42.750000 62.700001 42.759998 62.680000 42.740002 C 62.660000 42.720001 62.639999 42.740002 62.630001 42.740002 C 62.606850 42.735886 62.583153 42.735886 62.560001 42.740002 C 62.541187 42.730194 62.520935 42.723442 62.500000 42.720001 L 58.270000 42.720001 C 57.640396 42.720001 57.130001 43.230396 57.130001 43.860001 L 57.130001 69.830002 C 57.135490 70.455696 57.644279 70.960030 58.270000 70.959999 L 62.490002 70.959999 C 63.115723 70.960022 63.624512 70.455696 63.630001 69.830002 L 63.630001 54.400002 L 75.529999 70.459999 C 75.606644 70.577583 75.705078 70.679413 75.820000 70.760002 L 75.830002 70.760002 C 75.851761 70.778770 75.875191 70.795509 75.900002 70.810005 L 75.930000 70.809998 C 75.951637 70.820686 75.971809 70.834137 75.989998 70.849998 C 76.010002 70.849998 76.019997 70.849998 76.040001 70.870003 C 76.053200 70.872696 76.066803 70.872696 76.080002 70.870003 C 76.104523 70.887260 76.131485 70.900742 76.160004 70.910004 C 76.166473 70.912292 76.173531 70.912292 76.180000 70.910004 C 76.277779 70.936607 76.378670 70.950066 76.480003 70.950005 L 80.680000 70.949997 C 81.301819 70.944565 81.804573 70.441818 81.809998 69.820000 L 81.809998 43.860001 C 81.810020 43.234280 81.305695 42.725491 80.680000 42.720001" stroke="none" stroke-width="1.000000" stroke-dasharray="" stroke-dashoffset="1.000000" fill="white" fill-rule="evenodd" fill-opacity="1.000000"/>
                <use clip-path="url(#circle)" xlinkHref="#line" />
            </svg>
        </Link>
    );
}

export { TwitterButton, FacebookButton, MeweButton, LineButton };
