package getter

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/zu1k/proxypool/pkg/proxy"
	"github.com/zu1k/proxypool/pkg/tool"
)

func init() {
	Register("web-fanqiangdang-rss", NewWebFanqiangdangRSSGetter)
	Register("web-fanqiangdang", NewWebFanqiangdangGetter)
}

type WebFanqiangdang struct {
	c       *colly.Collector
	Url     string
	results proxy.ProxyList
}

func NewWebFanqiangdangGetter(options tool.Options) (getter Getter, err error) {
	urlInterface, found := options["url"]
	if found {
		url, err := AssertTypeStringNotNull(urlInterface)
		if err != nil {
			return nil, err
		}
		return &WebFanqiangdang{
			c:   colly.NewCollector(),
			Url: url,
		}, nil
	}
	return nil, ErrorUrlNotFound
}

func (w *WebFanqiangdang) Get() proxy.ProxyList {
	w.results = make(proxy.ProxyList, 0)
	w.c.OnHTML("td.t_f", func(e *colly.HTMLElement) {
		w.results = append(w.results, FuzzParseProxyFromString(e.Text)...)
		subUrls := urlRe.FindAllString(e.Text, -1)
		for _, url := range subUrls {
			w.results = append(w.results, (&Subscribe{Url: url}).Get()...)
		}
	})

	w.c.OnHTML("th.new>a[href]", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if strings.HasPrefix(url, "https://fanqiangdang.com/thread") {
			_ = e.Request.Visit(url)
		}
	})

	w.results = make(proxy.ProxyList, 0)
	err := w.c.Visit(w.Url)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
	}

	return w.results
}

func (w *WebFanqiangdang) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	defer wg.Done()
	nodes := w.Get()
	for _, node := range nodes {
		pc <- node
	}
}

type WebFanqiangdangRSS struct {
	c       *colly.Collector
	Url     string
	results []string
}

func NewWebFanqiangdangRSSGetter(options tool.Options) (getter Getter, err error) {
	urlInterface, found := options["url"]
	if found {
		url, err := AssertTypeStringNotNull(urlInterface)
		if err != nil {
			return nil, err
		}
		return &WebFanqiangdangRSS{
			c:   tool.GetColly(),
			Url: url,
		}, nil
	}
	return nil, ErrorUrlNotFound
}

func (w *WebFanqiangdangRSS) Get() proxy.ProxyList {
	w.results = make([]string, 0)
	w.c.OnHTML("td.t_f", func(e *colly.HTMLElement) {
		w.results = append(w.results, GrepLinksFromString(e.Text)...)
	})

	w.c.OnXML("//item//link", func(e *colly.XMLElement) {
		_ = e.Request.Visit(e.Text)
	})

	w.results = make([]string, 0)
	err := w.c.Visit(w.Url)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
	}

	return StringArray2ProxyArray(w.results)
}

func (w *WebFanqiangdangRSS) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	defer wg.Done()
	nodes := w.Get()
	for _, node := range nodes {
		pc <- node
	}
}
