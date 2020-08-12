package getter

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/zu1k/proxypool/proxy"
	"github.com/zu1k/proxypool/tool"
)

type Subscribe struct {
	Url string
}

func (s Subscribe) Get() []proxy.Proxy {
	resp, err := http.Get(s.Url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	nodesString, err := tool.Base64DecodeString(string(body))
	if err != nil {
		return nil
	}
	nodesString = strings.ReplaceAll(nodesString, "\t", "")

	nodes := strings.Split(nodesString, "\n")
	return StringArray2ProxyArray(nodes)
}

func NewSubscribe(url string) *Subscribe {
	return &Subscribe{
		Url: url,
	}
}
