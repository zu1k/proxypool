package getter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/zu1k/proxypool/pkg/proxy"
	"github.com/zu1k/proxypool/pkg/tool"
)

func init() {
	Register("web-freessrxyz", NewWebFreessrxyzGetter)
}

const (
	freessrxyzSsrLink   = "https://api.free-ssr.xyz/ssr"
	freessrxyzV2rayLink = "https://api.free-ssr.xyz/v2ray"
)

type WebFreessrXyz struct {
}

func NewWebFreessrxyzGetter(options tool.Options) (getter Getter, err error) {
	return &WebFreessrXyz{}, nil
}

func (w *WebFreessrXyz) Get() proxy.ProxyList {
	results := freessrxyzFetch(freessrxyzSsrLink)
	results = append(results, freessrxyzFetch(freessrxyzV2rayLink)...)
	return results
}

func (w *WebFreessrXyz) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	defer wg.Done()
	nodes := w.Get()
	log.Printf("STATISTIC: FreeSSRxyz\tcount=%d\turl=%s\n", len(nodes), "api.free-ssr.xyz")
	for _, node := range nodes {
		pc <- node
	}
}

func freessrxyzFetch(link string) proxy.ProxyList {
	resp, err := tool.GetHttpClient().Get(link)
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
