package provider

import (
	"encoding/json"
	"strconv"

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
	sub.Types = "ss,ssr"
	sub.preFilter()
	proxies := make([]ssJson, 0, sub.Proxies.Len())
	for _, p := range *sub.Proxies {
		var pp *proxy.Shadowsocks

		if p.TypeName() == "ssr" {
			var err error
			pp, err = proxy.SSR2SS(p.(*proxy.ShadowsocksR))
			if err != nil {
				continue
			}
		} else if p.TypeName() == "ss" {
			pp = p.(*proxy.Shadowsocks)
		}

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
