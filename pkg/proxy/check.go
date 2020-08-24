package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/ivpusic/grpool"

	"github.com/Dreamacro/clash/adapters/outbound"
)

const defaultURLTestTimeout = time.Second * 5

func testDelay(p Proxy) (delay uint16, err error) {
	pmap := make(map[string]interface{})
	err = json.Unmarshal([]byte(p.String()), &pmap)
	if err != nil {
		return
	}

	pmap["port"] = int(pmap["port"].(float64))
	if p.TypeName() == "vmess" {
		pmap["alterId"] = int(pmap["alterId"].(float64))
	}

	clashProxy, err := outbound.ParseProxy(pmap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultURLTestTimeout)
	delay, err = clashProxy.URLTest(ctx, "http://www.gstatic.com/generate_204")
	cancel()
	return delay, err
}

func CleanBadProxiesWithGrpool(proxies []Proxy) (cproxies []Proxy) {
	pool := grpool.NewPool(500, 200)

	c := make(chan checkResult)
	defer close(c)

	pool.WaitCount(len(proxies))
	go func() {
		for _, p := range proxies {
			pp := p
			pool.JobQueue <- func() {
				defer pool.JobDone()
				delay, err := testDelay(pp)
				if err == nil {
					c <- checkResult{
						name:  pp.Identifier(),
						delay: delay,
					}
				}
			}
		}
	}()
	done := make(chan struct{})
	defer close(done)

	go func() {
		pool.WaitAll()
		pool.Release()
		done <- struct{}{}
	}()

	okMap := make(map[string]struct{})
	for {
		select {
		case r := <-c:
			if r.delay > 0 {
				okMap[r.name] = struct{}{}
			}
		case <-done:
			cproxies = make(ProxyList, 0, 500)
			for _, p := range proxies {
				if _, ok := okMap[p.Identifier()]; ok {
					cproxies = append(cproxies, p.Clone())
				}
			}
			return
		}
	}
}

func CleanBadProxies(proxies []Proxy) (cproxies []Proxy) {
	c := make(chan checkResult, 40)
	wg := &sync.WaitGroup{}
	wg.Add(len(proxies))
	for _, p := range proxies {
		go testProxyDelayToChan(p, c, wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	okMap := make(map[string]struct{})
	for r := range c {
		if r.delay > 0 {
			okMap[r.name] = struct{}{}
		}
	}
	cproxies = make(ProxyList, 0, 500)
	for _, p := range proxies {
		if _, ok := okMap[p.Identifier()]; ok {
			cproxies = append(cproxies, p.Clone())
		}
	}
	return
}

type checkResult struct {
	name  string
	delay uint16
}

func testProxyDelayToChan(p Proxy, c chan checkResult, wg *sync.WaitGroup) {
	defer wg.Done()
	delay, err := testDelay(p)
	if err == nil {
		c <- checkResult{
			name:  p.Identifier(),
			delay: delay,
		}
	}
}
