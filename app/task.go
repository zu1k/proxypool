package app

import (
	"fmt"

	"github.com/zu1k/proxypool/getter"
	"github.com/zu1k/proxypool/proxy"
)

func CrawlTGChannel() {
	node := make([]proxy.Proxy, 0)
	node = append(node, getter.NewTGSsrlistGetter("https://t.me/s/ssrList", 200).Get()...)
	node = append(node, getter.NewTGSsrlistGetter("https://t.me/s/SSRSUB", 200).Get()...)
	node = append(node, getter.NewTGSsrlistGetter("https://t.me/s/FreeSSRNode", 200).Get()...)
	node = append(node, getter.NewTGSsrlistGetter("https://t.me/s/ssrlists", 200).Get()...)

	node = append(node, GetProxies()...)
	node = proxy.Deduplication(node)
	fmt.Println(len(node))
	SetProxies(node)
}
