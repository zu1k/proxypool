package app

import (
	"github.com/zu1k/proxypool/getter"
	"github.com/zu1k/proxypool/proxy"
)

func CrawlTGChannel() {
	node := make([]proxy.Proxy, 0)

	// tg上各种节点分享频道
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/ssrList", 200).Get()...)
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/SSRSUB", 200).Get()...)
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/FreeSSRNode", 200).Get()...)
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/ssrlists", 200).Get()...)
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/ssrshares", 200).Get()...)
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/V2List", 200).Get()...)
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/ssrtool", 200).Get()...)
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/vmessr", 200).Get()...)
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/FreeSSR666", 200).Get()...)
	node = append(node, getter.NewTGChannelGetter("https://t.me/s/fanqiang666", 200).Get()...)

	// 各种网站上公开的
	node = append(node, getter.WebFreessrXyz{}.Get()...)
	node = append(node, getter.WebLucnOrg{}.Get()...)

	// 从web页面模糊获取
	node = append(node, getter.NewWebFuzz("https://zfjvpn.gitbook.io/").Get()...)
	node = append(node, getter.NewWebFuzz("https://www.freefq.com/d/file/free-ssr/20200811/1f3e9d0d0064f662457062712dcf1b66.txt").Get()...)
	node = append(node, getter.NewWebFuzz("https://merlinblog.xyz/wiki/freess.html").Get()...)

	node = append(node, GetProxies()...)
	node = proxy.Deduplication(node)
	SetProxies(node)
}
