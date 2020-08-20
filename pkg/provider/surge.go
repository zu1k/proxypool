package provider

import (
	"strings"

	"github.com/zu1k/proxypool/pkg/proxy"
)

type Surge struct {
	Proxies proxy.ProxyList `yaml:"proxies"`
}

func (s Surge) Provide() string {
	var resultBuilder strings.Builder

	for _, p := range s.Proxies {
		if checkSurgeSupport(p) {
			resultBuilder.WriteString(p.ToSurge() + "\n")
		}
	}

	return resultBuilder.String()
}

func checkSurgeSupport(p proxy.Proxy) bool {
	switch p.(type) {
	case *proxy.ShadowsocksR:
		return false
	case *proxy.Vmess:
		return true
	case *proxy.Shadowsocks:
		ss := p.(*proxy.Shadowsocks)
		if checkInList(ssCipherList, ss.Cipher) {
			return true
		}
	default:
		return false
	}
	return false
}
