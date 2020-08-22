package provider

import (
	"strings"

	"github.com/zu1k/proxypool/pkg/proxy"
)

type Clash struct {
	Proxies proxy.ProxyList `yaml:"proxies"`
	Types   string          `yaml:"type"`
	Country string          `yaml:"country"`
}

func (c Clash) CleanProxies() (proxies proxy.ProxyList) {
	proxies = make(proxy.ProxyList, 0)
	for _, p := range c.Proxies {
		if checkClashSupport(p) {
			proxies = append(proxies, p)
		}
	}
	return
}

func (c Clash) Provide() string {
	var resultBuilder strings.Builder
	resultBuilder.WriteString("proxies:\n")

	noNeedFilterType := false
	noNeedFilterCountry := false
	if c.Types == "" || c.Types == "all" {
		noNeedFilterType = true
	}
	if c.Country == "" || c.Country == "all" {
		noNeedFilterCountry = true
	}
	types := strings.Split(c.Types, ",")
	countries := strings.Split(c.Country, ",")

	for _, p := range c.Proxies {
		if !checkClashSupport(p) {
			continue
		}

		typeOk := false
		countryOk := false
		if !noNeedFilterType {
			for _, t := range types {
				if p.TypeName() == t {
					typeOk = true
					break
				}
			}
		}
		if !noNeedFilterCountry {
			for _, c := range countries {
				if strings.HasPrefix(p.BaseInfo().Name, c) {
					countryOk = true
					break
				}
			}
		}

		if (noNeedFilterType || typeOk) && (noNeedFilterCountry || countryOk) {
			resultBuilder.WriteString(p.ToClash() + "\n")
		}
	}

	return resultBuilder.String()
}

func checkClashSupport(p proxy.Proxy) bool {
	switch p.TypeName() {
	case "ssr":
		ssr := p.(*proxy.ShadowsocksR)
		if checkInList(ssrCipherList, ssr.Cipher) && checkInList(ssrProtocolList, ssr.Protocol) && checkInList(ssrObfsList, ssr.Obfs) {
			return true
		}
	case "vmess":
		vmess := p.(*proxy.Vmess)
		if checkInList(vmessCipherList, vmess.Cipher) {
			return true
		}
	case "ss":
		ss := p.(*proxy.Shadowsocks)
		if checkInList(ssCipherList, ss.Cipher) {
			return true
		}
	case "trojan":
		return true
	default:
		return false
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

var ssCipherList = []string{
	"aes-128-gcm",
	"aes-192-gcm",
	"aes-256-gcm",
	"aes-128-cfb",
	"aes-192-cfb",
	"aes-256-cfb",
	"aes-128-ctr",
	"aes-192-ctr",
	"aes-256-ctr",
	"rc4-md5",
	"chacha20-ietf",
	"xchacha20",
	"chacha20-ietf-poly1305",
	"xchacha20-ietf-poly1305",
}
