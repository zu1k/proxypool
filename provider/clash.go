package provider

import (
	"strings"

	"github.com/zu1k/proxypool/proxy"
)

type Clash struct {
	Proxies []proxy.Proxy `yaml:"proxies"`
}

func (c Clash) Provide() string {
	var resultBuilder strings.Builder

	resultBuilder.WriteString("proxies:\n")
	for _, p := range c.Proxies {
		if checkClashSupport(p) {
			resultBuilder.WriteString(p.ToClash() + "\n")
		}
	}

	return resultBuilder.String()
}

func checkClashSupport(p proxy.Proxy) bool {
	switch p.(type) {
	case proxy.ShadowsocksR:
		ssr := p.(proxy.ShadowsocksR)
		if checkInList(ssrCipherList, ssr.Cipher) && checkInList(ssrProtocolList, ssr.Protocol) && checkInList(ssrObfsList, ssr.Obfs) {
			return true
		}
	case proxy.Vmess:
		vmess := p.(proxy.Vmess)
		if checkInList(vmessCipherList, vmess.Cipher) {
			return true
		}
	default:
		return false
	}
	return false
}

func checkInList(list []string, item string) bool {
	for _, i := range list {
		if item == i {
			return true
		}
	}
	return false
}

var ssrCipherList = []string{
	"aes-128-cfb",
	"aes-192-cfb",
	"aes-256-cfb",
	"aes-128-ctr",
	"aes-192-ctr",
	"aes-256-ctr",
	"aes-128-ofb",
	"aes-192-ofb",
	"aes-256-ofb",
	"des-cfb",
	"bf-cfb",
	"cast5-cfb",
	"rc4-md5",
	"chacha20-ietf",
	"salsa20",
	"camellia-128-cfb",
	"camellia-192-cfb",
	"camellia-256-cfb",
	"idea-cfb",
	"rc2-cfb",
	"seed-cfb",
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
