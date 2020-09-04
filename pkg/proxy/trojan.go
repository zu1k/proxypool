package proxy

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrorNotTrojanink = errors.New("not a correct trojan link")
)

type Trojan struct {
	Base
	Password       string   `yaml:"password" json:"password"`
	ALPN           []string `yaml:"alpn,omitempty" json:"alpn,omitempty"`
	SNI            string   `yaml:"sni,omitempty" json:"sni,omitempty"`
	SkipCertVerify bool     `yaml:"skip-cert-verify,omitempty" json:"skip-cert-verify,omitempty"`
	UDP            bool     `yaml:"udp,omitempty" json:"udp,omitempty"`
}

/**
  - name: "trojan"
    type: trojan
    server: server
    port: 443
    password: yourpsk
    # udp: true
    # sni: example.com # aka server name
    # alpn:
    #   - h2
    #   - http/1.1
    # skip-cert-verify: true
*/

func (t Trojan) Identifier() string {
	return net.JoinHostPort(t.Server, strconv.Itoa(t.Port)) + t.Password
}

func (t Trojan) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(data)
}

func (t Trojan) ToClash() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return "- " + string(data)
}

func (t Trojan) ToSurge() string {
	return ""
}

func (t Trojan) Clone() Proxy {
	return &t
}

// https://p4gefau1t.github.io/trojan-go/developer/url/
func (t Trojan) Link() (link string) {
	query := url.Values{}
	if t.SNI != "" {
		query.Set("sni", url.QueryEscape(t.SNI))
	}

	uri := url.URL{
		Scheme:   "trojan",
		User:     url.User(url.QueryEscape(t.Password)),
		Host:     net.JoinHostPort(t.Server, strconv.Itoa(t.Port)),
		RawQuery: query.Encode(),
		Fragment: t.Name,
	}

	return uri.String()
}

func ParseTrojanLink(link string) (*Trojan, error) {
	if !strings.HasPrefix(link, "trojan://") && !strings.HasPrefix(link, "trojan-go://") {
		return nil, ErrorNotTrojanink
	}

	/**
	trojan-go://
	    $(trojan-password)
	    @
	    trojan-host
	    :
	    port
	/?
	    sni=$(tls-sni.com)&
	    type=$(original|ws|h2|h2+ws)&
	        host=$(websocket-host.com)&
	        path=$(/websocket/path)&
	    encryption=$(ss;aes-256-gcm;ss-password)&
	    plugin=$(...)
	#$(descriptive-text)
	*/

	uri, err := url.Parse(link)
	if err != nil {
		return nil, ErrorNotSSLink
	}

	password := uri.User.Username()
	password, _ = url.QueryUnescape(password)

	server := uri.Hostname()
	port, _ := strconv.Atoi(uri.Port())

	moreInfos := uri.Query()
	sni := moreInfos.Get("sni")
	sni, _ = url.QueryUnescape(sni)
	transformType := moreInfos.Get("type")
	transformType, _ = url.QueryUnescape(transformType)
	host := moreInfos.Get("host")
	host, _ = url.QueryUnescape(host)
	path := moreInfos.Get("path")
	path, _ = url.QueryUnescape(path)

	alpn := make([]string, 0)
	if transformType == "h2" {
		alpn = append(alpn, "h2")
	}

	if port == 0 {
		return nil, ErrorNotTrojanink
	}

	return &Trojan{
		Base: Base{
			Name:   strconv.Itoa(rand.Int()),
			Server: server,
			Port:   port,
			Type:   "trojan",
		},
		Password:       password,
		ALPN:           alpn,
		UDP:            true,
		SNI:            host,
		SkipCertVerify: true,
	}, nil
}

var (
	trojanPlainRe = regexp.MustCompile("trojan(-go)?://([A-Za-z0-9+/_&?=@:%.-])+")
)

func GrepTrojanLinkFromString(text string) []string {
	results := make([]string, 0)
	texts := strings.Split(text, "trojan://")
	for _, text := range texts {
		results = append(results, trojanPlainRe.FindAllString("trojan://"+text, -1)...)
	}
	return results
}
