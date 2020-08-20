package getter

import (
	"errors"
	"strings"
	"sync"

	"github.com/zu1k/proxypool/pkg/proxy"
	"github.com/zu1k/proxypool/pkg/tool"
)

type Getter interface {
	Get() proxy.ProxyList
	Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup)
}

type creator func(options tool.Options) (getter Getter, err error)

var creatorMap = make(map[string]creator)

func Register(sourceType string, c creator) {
	creatorMap[sourceType] = c
}

func NewGetter(sourceType string, options tool.Options) (getter Getter, err error) {
	c, ok := creatorMap[sourceType]
	if ok {
		return c(options)
	}
	return nil, ErrorCreaterNotSupported
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
	} else if strings.HasPrefix(link, "trojan://") {
		data, err = proxy.ParseTrojanLink(link)
	}
	if err != nil {
		return nil
	}
	return data
}

func StringArray2ProxyArray(origin []string) proxy.ProxyList {
	results := make(proxy.ProxyList, 0)
	for _, link := range origin {
		results = append(results, String2Proxy(link))
	}
	return results
}

func GrepLinksFromString(text string) []string {
	results := proxy.GrepSSRLinkFromString(text)
	results = append(results, proxy.GrepVmessLinkFromString(text)...)
	results = append(results, proxy.GrepSSLinkFromString(text)...)
	results = append(results, proxy.GrepTrojanLinkFromString(text)...)
	return results
}

func FuzzParseProxyFromString(text string) proxy.ProxyList {
	return StringArray2ProxyArray(GrepLinksFromString(text))
}

var (
	ErrorUrlNotFound         = errors.New("url should be specified")
	ErrorCreaterNotSupported = errors.New("type not supported")
)

func AssertTypeStringNotNull(i interface{}) (str string, err error) {
	switch i.(type) {
	case string:
		str = i.(string)
		if str == "" {
			return "", errors.New("string is null")
		}
		return str, nil
	default:
		return "", errors.New("type is not string")
	}
	return "", nil
}
