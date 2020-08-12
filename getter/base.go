package getter

import (
	"strings"

	"github.com/zu1k/proxypool/proxy"
)

type Getter interface {
	Get() []*proxy.Proxy
}

func String2Proxy(link string) proxy.Proxy {
	var err error
	var data proxy.Proxy
	if strings.HasPrefix(link, "ssr://") {
		data, err = proxy.ParseSSRLink(link)
	} else if strings.HasPrefix(link, "vmess://") {
		data, err = proxy.ParseVmessLink(link)
	}
	if err != nil {
		return nil
	}
	return data
}

func StringArray2ProxyArray(origin []string) []proxy.Proxy {
	var err error
	results := make([]proxy.Proxy, 0)
	for _, link := range origin {
		var data proxy.Proxy
		if strings.HasPrefix(link, "ssr://") {
			data, err = proxy.ParseSSRLink(link)
		} else if strings.HasPrefix(link, "vmess://") {
			data, err = proxy.ParseVmessLink(link)
		}
		if err != nil {
			continue
		}
		results = append(results, data)
	}
	return results
}
