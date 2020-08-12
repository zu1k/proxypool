package app

import (
	"math/rand"
	"strconv"

	"github.com/zu1k/proxypool/provider"

	"github.com/zu1k/proxypool/app/cache"

	"github.com/zu1k/proxypool/getter"
	"github.com/zu1k/proxypool/proxy"
)

func CrawlTGChannel() {
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

	proxies = append(proxies, cache.GetProxies()...)
	proxies = proxy.Deduplication(proxies)

	num := len(proxies)
	for i := 0; i < num; i++ {
		proxies[i].SetName("@tgbotlist_" + strconv.Itoa(rand.Int()))
	}
	cache.SetProxies(proxies)
	cache.SetString("clashproxies", provider.Clash{Proxies: proxies}.Provide())
}
