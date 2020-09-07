package provider

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/zu1k/proxypool/pkg/tool"

	"github.com/zu1k/proxypool/pkg/proxy"
)

type SSSub struct {
	Base
}

type ssJson struct {
	Remarks    string                 `json:"remarks"`
	Server     string                 `json:"server"`
	ServerPort string                 `json:"server_port"`
	Method     string                 `json:"method"`
	Password   string                 `json:"password"`
	Plugin     string                 `json:"plugin"`
	PluginOpts map[string]interface{} `json:"plugin_opts"`
}

func (sub SSSub) Provide() string {
	sub.Types = "ss"
	sub.preFilter()
	proxies := make([]ssJson, 0, sub.Proxies.Len())
	for _, p := range *sub.Proxies {
		pp := p.(*proxy.Shadowsocks)
		proxies = append(proxies, ssJson{
			Remarks:    pp.Name,
			Server:     pp.Server,
			ServerPort: strconv.Itoa(pp.Port),
			Method:     pp.Cipher,
			Password:   pp.Password,
			Plugin:     pp.Plugin,
			PluginOpts: pp.PluginOpts,
		})
	}
	text, err := json.Marshal(proxies)
	if err != nil {
		return ""
	}
	return string(text)
}

type SIP002Sub struct {
	Base
}

func (sub SIP002Sub) Provide() string {
	sub.Types = "ss"
	sub.preFilter()
	var resultBuilder strings.Builder
	for _, p := range *sub.Proxies {
		resultBuilder.WriteString(p.Link() + "\n")
	}
	return tool.Base64EncodeString(resultBuilder.String(), false)
}
