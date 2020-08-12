package getter

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/zu1k/proxypool/proxy"
	"github.com/zu1k/proxypool/tool"
)

func init() {
	Register("web-lucnorg", NewWebLucnorg)
}

const lucnorgSsrLink = "https://lncn.org/api/ssrList"

type WebLucnOrg struct {
}

func NewWebLucnorg(options tool.Options) Getter {
	return &WebLucnOrg{}
}

func (w *WebLucnOrg) Get() []proxy.Proxy {
	resp, err := http.Post(lucnorgSsrLink, "", nil)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	response := struct {
		Code string `json:"code"`
		Ssrs string `json:"ssrs"`
	}{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}

	dec := decryptAesForLucn(response.Code, response.Ssrs)
	if dec == nil {
		return nil
	}

	type node struct {
		Url string `json:"url"`
	}
	ssrs := make([]node, 0)
	err = json.Unmarshal(dec, &ssrs)
	if err != nil {
		return nil
	}

	result := make([]string, 0)
	for _, node := range ssrs {
		result = append(result, node.Url)
	}
	return StringArray2ProxyArray(result)
}

func (w *WebLucnOrg) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	defer wg.Done()
	nodes := w.Get()
	for _, node := range nodes {
		pc <- node
	}
}

func decryptAesForLucn(code string, c string) []byte {
	if code == "" {
		code = "abclnv561cqqfg30"
	}
	cipher, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil
	}
	result := tool.AesEcbDecryptWithPKCS7Unpadding(cipher, []byte(code))
	return result
}
