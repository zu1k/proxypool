package app

import (
	"time"

	"github.com/zu1k/proxypool/proxy"

	"github.com/patrickmn/go-cache"
)

var c = cache.New(cache.NoExpiration, 10*time.Minute)

func GetProxies() []proxy.Proxy {
	result, found := c.Get("proxies")
	if found {
		return result.([]proxy.Proxy)
	}
	return nil
}

func SetProxies(proxies []proxy.Proxy) {
	c.Set("proxies", proxies, cache.NoExpiration)
}
