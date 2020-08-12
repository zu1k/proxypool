package getter

import (
	"io/ioutil"
	"net/http"

	"github.com/zu1k/proxypool/proxy"
)

type WebFuzz struct {
	Url string
}

func (w WebFuzz) Get() []proxy.Proxy {
	resp, err := http.Get(w.Url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	text := string(body)

	results := proxy.GrepSSRLinkFromString(text)
	results = append(results, proxy.GrepVmessLinkFromString(text)...)

	return StringArray2ProxyArray(results)
}

func NewWebFuzz(url string) *WebFuzz {
	return &WebFuzz{Url: url}
}
