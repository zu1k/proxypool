package app

import (
	"log"
	"math/rand"
	"strconv"
	"sync"

	"github.com/zu1k/proxypool/provider"

	"github.com/zu1k/proxypool/app/cache"

	"github.com/zu1k/proxypool/getter"
	"github.com/zu1k/proxypool/proxy"
)

func Crawl() {
	proxies := make([]proxy.Proxy, 0)

	// tg上各种节点分享频道
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/ssrList", 200).Get()...)
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/SSRSUB", 200).Get()...)
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/FreeSSRNode", 200).Get()...)
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/ssrlists", 200).Get()...)
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/ssrshares", 200).Get()...)
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/V2List", 200).Get()...)
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/ssrtool", 200).Get()...)
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/vmessr", 200).Get()...)
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/FreeSSR666", 200).Get()...)
	proxies = append(proxies, getter.NewTGChannelGetter("https://t.me/s/fanqiang666", 200).Get()...)

	// 各种网站上公开的
	proxies = append(proxies, getter.WebFreessrXyz{}.Get()...)
	proxies = append(proxies, getter.WebLucnOrg{}.Get()...)

	// 从web页面模糊获取
	proxies = append(proxies, getter.NewWebFuzz("https://zfjvpn.gitbook.io/").Get()...)
	proxies = append(proxies, getter.NewWebFuzz("https://www.freefq.com/d/file/free-ssr/20200811/1f3e9d0d0064f662457062712dcf1b66.txt").Get()...)
	proxies = append(proxies, getter.NewWebFuzz("https://merlinblog.xyz/wiki/freess.html").Get()...)
	// 翻墙党
	proxies = append(proxies, getter.NewWebFanqiangdangGetter("https://fanqiangdang.com/forum.php?mod=rss&fid=50&auth=0", 200).Get()...)
	proxies = append(proxies, getter.NewWebFanqiangdangGetter("https://fanqiangdang.com/forum.php?mod=rss&fid=2&auth=0", 200).Get()...)
	proxies = append(proxies, getter.NewWebFanqiangdangGetter("https://fanqiangdang.com/forum.php?mod=rss&fid=36&auth=0", 200).Get()...)

	// 订阅链接
	proxies = append(proxies, getter.NewSubscribe("https://raw.githubusercontent.com/ssrsub/ssr/master/v2ray").Get()...)
	proxies = append(proxies, getter.NewSubscribe("https://raw.githubusercontent.com/ssrsub/ssr/master/ssrsub").Get()...)
	proxies = append(proxies, getter.NewSubscribe("https://raw.githubusercontent.com/ssrsub/ssr/master/ss-sub").Get()...)

	proxies = append(proxies, cache.GetProxies()...)
	proxies = proxy.Deduplication(proxies)

	num := len(proxies)
	for i := 0; i < num; i++ {
		proxies[i].SetName(strconv.Itoa(rand.Int()))
	}
	cache.SetProxies(proxies)
	cache.SetString("clashproxies", provider.Clash{Proxies: proxies}.Provide())
}

func CrawlGo() {
	wg := sync.WaitGroup{}
	var pc = make(chan proxy.Proxy)
	// tg上各种节点分享频道
	go getter.NewTGChannelGetter("https://t.me/s/ssrList", 200).Get2Chan(pc, &wg)
	go getter.NewTGChannelGetter("https://t.me/s/SSRSUB", 200).Get2Chan(pc, &wg)
	go getter.NewTGChannelGetter("https://t.me/s/FreeSSRNode", 200).Get2Chan(pc, &wg)
	go getter.NewTGChannelGetter("https://t.me/s/ssrlists", 200).Get2Chan(pc, &wg)
	go getter.NewTGChannelGetter("https://t.me/s/ssrshares", 200).Get2Chan(pc, &wg)
	go getter.NewTGChannelGetter("https://t.me/s/V2List", 200).Get2Chan(pc, &wg)
	go getter.NewTGChannelGetter("https://t.me/s/ssrtool", 200).Get2Chan(pc, &wg)
	go getter.NewTGChannelGetter("https://t.me/s/vmessr", 200).Get2Chan(pc, &wg)
	go getter.NewTGChannelGetter("https://t.me/s/FreeSSR666", 200).Get2Chan(pc, &wg)
	go getter.NewTGChannelGetter("https://t.me/s/fanqiang666", 200).Get2Chan(pc, &wg)

	// 各种网站上公开的
	go getter.WebFreessrXyz{}.Get2Chan(pc, &wg)
	go getter.WebLucnOrg{}.Get2Chan(pc, &wg)

	// 从web页面模糊获取
	go getter.NewWebFuzz("https://zfjvpn.gitbook.io/").Get2Chan(pc, &wg)
	go getter.NewWebFuzz("https://www.freefq.com/d/file/free-ssr/20200811/1f3e9d0d0064f662457062712dcf1b66.txt").Get2Chan(pc, &wg)
	go getter.NewWebFuzz("https://merlinblog.xyz/wiki/freess.html").Get2Chan(pc, &wg)
	// 翻墙党
	go getter.NewWebFanqiangdangGetter("https://fanqiangdang.com/forum.php?mod=rss&fid=50&auth=0", 200).Get2Chan(pc, &wg)
	go getter.NewWebFanqiangdangGetter("https://fanqiangdang.com/forum.php?mod=rss&fid=2&auth=0", 200).Get2Chan(pc, &wg)
	go getter.NewWebFanqiangdangGetter("https://fanqiangdang.com/forum.php?mod=rss&fid=36&auth=0", 200).Get2Chan(pc, &wg)

	// 订阅链接
	go getter.NewSubscribe("https://raw.githubusercontent.com/ssrsub/ssr/master/v2ray").Get2Chan(pc, &wg)
	go getter.NewSubscribe("https://raw.githubusercontent.com/ssrsub/ssr/master/ssrsub").Get2Chan(pc, &wg)
	go getter.NewSubscribe("https://raw.githubusercontent.com/ssrsub/ssr/master/ss-sub").Get2Chan(pc, &wg)

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
