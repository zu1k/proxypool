package proxy

import (
	"encoding/json"
	"fmt"
)

func ParseSurgeLink(link string) (Proxy, error) {
	link = link[2:]
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(link), &m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	switch(m["type"]) {
	case "ss":
		return &Shadowsocks{
			Base: Base{
				Name: m["name"].(string),
				Server: m["server"].(string),
				Port: int(m["port"].(float64)),
				Type: m["type"].(string),
				Country: m["country"].(string),
			},

			Password: m["password"].(string),
			Cipher: m["cipher"].(string),

		}, nil
	case "vmess":
		fmt.Println(m)
		wsHeaders := make(map[string]string)
		if _, ok := m["ws-headers"]; ok {
			tmp := m["ws-headers"].(map[string]interface{})
			for k, v := range tmp {
				wsHeaders[k] = v.(string)
			}
		}
		servername := ""
		if _, ok := m["servername"]; ok {
			servername = m["servername"].(string)
		}
		tls := false
		if _, ok := m["tls"]; ok {
			tls = m["tls"].(bool)
		}
		cipher := ""
		if _, ok := m["cipher"]; ok {
			cipher = m["cipher"].(string)
		}
		alterId := 0
		if _, ok := m["alterId"]; ok {
			alterId = int(m["alterId"].(float64))
		}

		return &Vmess{
			Base: Base{
				Name: m["name"].(string),
				Server: m["server"].(string),
				Port: int(m["port"].(float64)),
				Type: m["type"].(string),
				Country: m["country"].(string),
				UDP: false,
			},
			UUID:           m["uuid"].(string),
			AlterID:        alterId,
			Cipher:         cipher,
			TLS:            tls,
			Network:        m["network"].(string),
			HTTPOpts:       HTTPOptions{},
			WSPath:         m["ws-path"].(string),
			WSHeaders:      wsHeaders,
			SkipCertVerify: m["skip-cert-verify"].(bool),
			ServerName:     servername,
		}, nil
	}
	return nil, nil
}
