package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Dreamacro/clash/adapters/outbound"
)

const defaultURLTestTimeout = time.Second * 15

func Check(p Proxy) (delay uint16, err error) {
	pmap := make(map[string]interface{})
	err = json.Unmarshal([]byte(p.String()), &pmap)
	if err != nil {
		return
	}

	pmap["port"] = int(pmap["port"].(float64))
	if p.Type() == "vmess" {
		pmap["alterId"] = int(pmap["alterId"].(float64))
	}

	clashProxy, err := outbound.ParseProxy(pmap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultURLTestTimeout)
	defer cancel()
	delay, err = clashProxy.URLTest(ctx, "http://www.gstatic.com/generate_204")
	return delay, err
}

func CleanProxies(proxies []Proxy) (cproxies []Proxy) {
	c := make(chan checkResult, 40)
	for _, p := range proxies {
		go checkProxyWithChan(p, c)
	}
	okMap := make(map[string]struct{})
	size := len(proxies)
	for i := 0; i < size; i++ {
		r := <-c
		if r.delay > 0 {
			okMap[r.name] = struct{}{}
		}
	}
	cproxies = make([]Proxy, 0)
	for _, p := range proxies {
		if _, ok := okMap[p.Identifier()]; ok {
			cproxies = append(cproxies, p)
		}
	}
	return
}

type checkResult struct {
	name  string
	delay uint16
}

func checkProxyWithChan(p Proxy, c chan checkResult) {
	delay, err := Check(p)
	if err != nil {
		c <- checkResult{
			name:  p.Identifier(),
			delay: 0,
		}
	}
	c <- checkResult{
		name:  p.Identifier(),
		delay: delay,
	}
}
