package getter

import (
	"io/ioutil"
	"sync"

	"github.com/zu1k/proxypool/proxy"
	"github.com/zu1k/proxypool/tool"
)

func init() {
	Register("webfuzz", NewWebFuzzGetter)
}

type WebFuzz struct {
	Url string
}

func (w *WebFuzz) Get() []proxy.Proxy {
	resp, err := tool.GetHttpClient().Get(w.Url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return FuzzParseProxyFromString(string(body))
}

func (w *WebFuzz) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	defer wg.Done()
	nodes := w.Get()
	for _, node := range nodes {
		pc <- node
	}
}

func NewWebFuzzGetter(options tool.Options) Getter {
	url, found := options["url"]
	if found {
		return &WebFuzz{Url: url.(string)}
	}
	return nil
}
