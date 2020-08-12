package getter

import (
	"strings"

	"github.com/zu1k/proxypool/proxy"
)

type Getter interface {
	Get() []proxy.Proxy
}

func String2Proxy(link string) proxy.Proxy {
	var err error
	var data proxy.Proxy
	if strings.HasPrefix(link, "ssr://") {
		data, err = proxy.ParseSSRLink(link)
	} else if strings.HasPrefix(link, "vmess://") {
		data, err = proxy.ParseVmessLink(link)
	} else if strings.HasPrefix(link, "ss://") {
		data, err = proxy.ParseSSLink(link)
	}
	if err != nil {
		return nil
	}
	return data
}

func StringArray2ProxyArray(origin []string) []proxy.Proxy {
	results := make([]proxy.Proxy, 0)
	for _, link := range origin {
		results = append(results, String2Proxy(link))
	}
	return results
}

func GrepLinksFromString(text string) []string {
	results := proxy.GrepSSRLinkFromString(text)
	results = append(results, proxy.GrepVmessLinkFromString(text)...)
	results = append(results, proxy.GrepSSLinkFromString(text)...)
	return results
}

func FuzzParseProxyFromString(text string) []proxy.Proxy {
	return StringArray2ProxyArray(GrepLinksFromString(text))
}
