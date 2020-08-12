package getter

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
	"github.com/zu1k/proxypool/proxy"
)

type WebFanqiangdang struct {
	c         *colly.Collector
	NumNeeded int
	Results   []string
	Url       string
}

func NewWebFanqiangdangGetter(url string, numNeeded int) *WebFanqiangdang {
	if numNeeded <= 0 {
		numNeeded = 200
	}
	return &WebFanqiangdang{
		c:         colly.NewCollector(),
		NumNeeded: numNeeded,
		Results:   make([]string, 0),
		Url:       url,
	}
}

func (w WebFanqiangdang) Get() []proxy.Proxy {
	// 找到所有的文字消息
	w.c.OnHTML("td.t_f", func(e *colly.HTMLElement) {
		w.Results = append(w.Results, GrepLinksFromString(e.Text)...)
	})

	// 从订阅中取出每一页，因为是订阅，所以都比较新
	w.c.OnXML("//item//link", func(e *colly.XMLElement) {
		if len(w.Results) < w.NumNeeded {
			_ = e.Request.Visit(e.Text)
		}
	})

	w.Results = make([]string, 0)
	err := w.c.Visit(w.Url)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
	}

	return StringArray2ProxyArray(w.Results)
}

func (w WebFanqiangdang) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	wg.Add(1)
	nodes := w.Get()
	for _, node := range nodes {
		pc <- node
	}
	wg.Done()
}
