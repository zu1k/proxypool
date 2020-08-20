package cache

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zu1k/proxypool/pkg/proxy"
)

var c = cache.New(cache.NoExpiration, 10*time.Minute)

func GetProxies(key string) proxy.ProxyList {
	result, found := c.Get(key)
	if found {
		log.Println("found cache for:", key, "length:", len(result.(proxy.ProxyList)))
		return result.(proxy.ProxyList)
	}
	log.Println("cache not found:", key)
	return nil
}

func SetProxies(key string, proxies proxy.ProxyList) {
	c.Set(key, proxies, cache.NoExpiration)
}

func SetString(key, value string) {
	c.Set(key, value, cache.NoExpiration)
}

func GetString(key string) string {
	result, found := c.Get(key)
	if found {
		return result.(string)
	}
	return ""
}
