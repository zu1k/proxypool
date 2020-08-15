package getter

import (
	"errors"
	"strings"
	"sync"

	"github.com/zu1k/proxypool/proxy"
	"github.com/zu1k/proxypool/tool"
)

type Getter interface {
	Get() []proxy.Proxy
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
	default:
		return "", errors.New("type is not string")
	}
	return "", nil
}
