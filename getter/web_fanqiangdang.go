package getter

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
	"github.com/zu1k/proxypool/proxy"
	"github.com/zu1k/proxypool/tool"
)

func init() {
	Register("web-fanqiangdang", NewWebFanqiangdangGetter)
}

type WebFanqiangdang struct {
	c         *colly.Collector
	NumNeeded int
	Url       string
	results   []string
}

func NewWebFanqiangdangGetter(options tool.Options) Getter {
	num, found := options["num"]

	t := 200
	switch num.(type) {
	case int:
		t = num.(int)
	case float64:
		t = int(num.(float64))
	}

	if !found || t <= 0 {
		t = 200
	}
	url, found := options["url"]
	if found {
		return &WebFanqiangdang{
			c:         colly.NewCollector(),
			NumNeeded: t,
			Url:       url.(string),
		}
	}
	return nil
}

func (w *WebFanqiangdang) Get() []proxy.Proxy {
	w.results = make([]string, 0)
	// 找到所有的文字消息
	w.c.OnHTML("td.t_f", func(e *colly.HTMLElement) {
		w.results = append(w.results, GrepLinksFromString(e.Text)...)
	})

	// 从订阅中取出每一页，因为是订阅，所以都比较新
	w.c.OnXML("//item//link", func(e *colly.XMLElement) {
		if len(w.results) < w.NumNeeded {
			_ = e.Request.Visit(e.Text)
		}
	})

	w.results = make([]string, 0)
	err := w.c.Visit(w.Url)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
	}

	return StringArray2ProxyArray(w.results)
}

func (w *WebFanqiangdang) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	defer wg.Done()
	nodes := w.Get()
	for _, node := range nodes {
		pc <- node
	}
}
