package getter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zu1k/proxypool/proxy"
)

const (
	freessrxyzSsrLink   = "https://api.free-ssr.xyz/ssr"
	freessrxyzV2rayLink = "https://api.free-ssr.xyz/v2ray"
)

type WebFreessrXyz struct {
}

func (w WebFreessrXyz) Get() []proxy.Proxy {
	results := freessrxyzFetch(freessrxyzSsrLink)
	results = append(results, freessrxyzFetch(freessrxyzV2rayLink)...)
	return results
}

func freessrxyzFetch(link string) []proxy.Proxy {
	resp, err := http.Get(link)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	type node struct {
		Url string `json:"url"`
	}
	ssrs := make([]node, 0)
	err = json.Unmarshal(body, &ssrs)
	if err != nil {
		return nil
	}

	result := make([]string, 0)
	for _, node := range ssrs {
		result = append(result, node.Url)
	}

	return StringArray2ProxyArray(result)
}
