package app

import (
	"log"
	"sync"
	"time"

	"github.com/zu1k/proxypool/internal/cache"
	"github.com/zu1k/proxypool/pkg/provider"
	"github.com/zu1k/proxypool/pkg/proxy"
)

var location, _ = time.LoadLocation("PRC")

func CrawlGo() {
	wg := &sync.WaitGroup{}
	var pc = make(chan proxy.Proxy)
	for _, g := range Getters {
		wg.Add(1)
		go g.Get2Chan(pc, wg)
	}
	proxies := cache.GetProxies("proxies")
	go func() {
		wg.Wait()
		close(pc)
	}()
	for node := range pc {
		if node != nil {
			proxies = append(proxies, node)
		}
	}
	// 节点去重
	proxies = proxies.Deduplication()
	log.Println("CrawlGo node count:", len(proxies))
	proxies = provider.Clash{Proxies: proxies}.CleanProxies()
	proxies.NameAddCounrty().Sort().NameAddIndex()
	cache.SetProxies("allproxies", proxies)
	cache.AllProxiesCount = proxies.Len()
	cache.SSProxiesCount = proxies.TypeLen("ss")
	cache.SSRProxiesCount = proxies.TypeLen("ssr")
	cache.VmessProxiesCount = proxies.TypeLen("vmess")
	cache.TrojanProxiesCount = proxies.TypeLen("trojan")
	cache.LastCrawlTime = time.Now().In(location).Format("2006-01-02 15:04:05")

	// 可用性检测
	proxies = proxy.CleanBadProxies(proxies)
	log.Println("CrawlGo clash useable node count:", len(proxies))
	proxies.NameAddCounrty().Sort().NameAddIndex()
	cache.SetProxies("proxies", proxies)
	cache.UsefullProxiesCount = proxies.Len()

	cache.SetString("clashproxies", provider.Clash{Proxies: proxies}.Provide())
	cache.SetString("surgeproxies", provider.Surge{Proxies: proxies}.Provide())
}
