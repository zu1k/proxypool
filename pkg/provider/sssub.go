package provider

import (
	"strings"

	"github.com/zu1k/proxypool/pkg/tool"
)

type SSSub struct {
	Base
}

func (sub SSSub) Provide() string {
	sub.Types = "ss"
	sub.preFilter()
	var resultBuilder strings.Builder
	for _, p := range *sub.Proxies {
		resultBuilder.WriteString(p.Link() + "\n")
	}
	return tool.Base64EncodeString(resultBuilder.String())
}
