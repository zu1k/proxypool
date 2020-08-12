package app

import (
	"log"
	"math/rand"
	"strconv"
	"sync"

	"github.com/zu1k/proxypool/provider"

	"github.com/zu1k/proxypool/app/cache"

	"github.com/zu1k/proxypool/proxy"
)

func Crawl() {
	proxies := make([]proxy.Proxy, 0)
	for _, g := range Getters {
		proxies = append(proxies, g.Get()...)
	}
	proxies = append(proxies, cache.GetProxies()...)
	proxies = proxy.Deduplication(proxies)

	num := len(proxies)
	for i := 0; i < num; i++ {
		proxies[i].SetName(strconv.Itoa(rand.Int()))
	}
	log.Println("Crawl node count:", num)
	cache.SetProxies(proxies)
	cache.SetString("clashproxies", provider.Clash{Proxies: proxies}.Provide())
}

func CrawlGo() {
	wg := &sync.WaitGroup{}
	var pc = make(chan proxy.Proxy)
	for _, g := range Getters {
		wg.Add(1)
		go g.Get2Chan(pc, wg)
	}
	proxies := cache.GetProxies()
	go func() {
		wg.Wait()
		close(pc)
	}()
	for node := range pc {
		if node != nil {
			proxies = append(proxies, node)
		}
	}
	proxies = proxy.Deduplication(proxies)

	num := len(proxies)
	for i := 0; i < num; i++ {
		proxies[i].SetName(strconv.Itoa(rand.Int()))
	}
	log.Println("CrawlGo node count:", num)
	cache.SetProxies(proxies)
	cache.SetString("clashproxies", provider.Clash{Proxies: proxies}.Provide())
}
