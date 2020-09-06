package provider

import (
	"strings"

	"github.com/zu1k/proxypool/pkg/proxy"

	"github.com/zu1k/proxypool/pkg/tool"
)

type SSRSub struct {
	Base
}

func (sub SSRSub) Provide() string {
	sub.Types = "ssr,ss"
	sub.preFilter()
	var resultBuilder strings.Builder
	for _, p := range *sub.Proxies {
		if p.TypeName() == "ssr" {
			resultBuilder.WriteString(p.Link() + "\n")
		} else if p.TypeName() == "ss" {
			ssr, err := proxy.SS2SSR(p.(*proxy.Shadowsocks))
			if err == nil {
				resultBuilder.WriteString(ssr.Link() + "\n")
			}
		}
	}
	return tool.Base64EncodeString(resultBuilder.String(), false)
}
