package provider

import (
	"strings"

	"github.com/zu1k/proxypool/proxy"
)

type Surge struct {
	Proxies []proxy.Proxy `yaml:"proxies"`
}

func (s Surge) Provide() string {
	var resultBuilder strings.Builder

	for _, p := range s.Proxies {
		if checkClashSupport(p) {
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
		return true
	default:
		return false
	}
	return false
}
