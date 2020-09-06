package provider

import (
	"strings"

	"github.com/zu1k/proxypool/pkg/tool"

	"github.com/zu1k/proxypool/pkg/proxy"
)

type Clash struct {
	Base
}

func (c Clash) CleanProxies() (proxies proxy.ProxyList) {
	proxies = make(proxy.ProxyList, 0)
	for _, p := range *c.Proxies {
		if checkClashSupport(p) {
			proxies = append(proxies, p)
		}
	}
	return
}

func (c Clash) Provide() string {
	c.preFilter()

	var resultBuilder strings.Builder
	resultBuilder.WriteString("proxies:\n")
	for _, p := range *c.Proxies {
		if checkClashSupport(p) {
			resultBuilder.WriteString(p.ToClash() + "\n")
		}
	}
	return resultBuilder.String()
}

func checkClashSupport(p proxy.Proxy) bool {
	switch p.TypeName() {
	case "ssr":
		ssr := p.(*proxy.ShadowsocksR)
		if tool.CheckInList(proxy.SSRCipherList, ssr.Cipher) && tool.CheckInList(ssrProtocolList, ssr.Protocol) && tool.CheckInList(ssrObfsList, ssr.Obfs) {
			return true
		}
	case "vmess":
		vmess := p.(*proxy.Vmess)
		if tool.CheckInList(vmessCipherList, vmess.Cipher) {
			return true
		}
	case "ss":
		ss := p.(*proxy.Shadowsocks)
		if tool.CheckInList(proxy.SSCipherList, ss.Cipher) {
			return true
		}
	case "trojan":
		return true
	default:
		return false
	}
	return false
}

var ssrObfsList = []string{
	"plain",
	"http_simple",
	"http_post",
	"random_head",
	"tls1.2_ticket_auth",
	"tls1.2_ticket_fastauth",
}

var ssrProtocolList = []string{
	"origin",
	"verify_deflate",
	"verify_sha1",
	"auth_sha1",
	"auth_sha1_v2",
	"auth_sha1_v4",
	"auth_aes128_md5",
	"auth_aes128_sha1",
	"auth_chain_a",
	"auth_chain_b",
}

var vmessCipherList = []string{
	"auto",
	"aes-128-gcm",
	"chacha20-poly1305",
	"none",
}
