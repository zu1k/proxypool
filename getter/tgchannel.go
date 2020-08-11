package getter

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/zu1k/proxypool/proxy"
)

type TGChannelGetter struct {
	c         *colly.Collector
	NumNeeded int
	Results   []string
	Url       string
}

func NewTGSsrlistGetter(url string, numNeeded int) *TGChannelGetter {
	if numNeeded <= 0 {
		numNeeded = 200
	}
	return &TGChannelGetter{
		c:         colly.NewCollector(),
		NumNeeded: numNeeded,
		Results:   make([]string, 0),
		Url:       url,
	}
}

func (g TGChannelGetter) Get() []proxy.Proxy {
	// 找到所有的文字消息
	g.c.OnHTML("div.tgme_widget_message_text", func(e *colly.HTMLElement) {
		g.Results = append(g.Results, proxy.GrepSSRLinkFromString(e.Text)...)
	})

	// 找到之前消息页面的链接，加入访问队列
	g.c.OnHTML("link[rel=prev]", func(e *colly.HTMLElement) {
		if len(g.Results) < g.NumNeeded {
			_ = e.Request.Visit(e.Attr("href"))
		}
	})

	g.Results = make([]string, 0)
	err := g.c.Visit(g.Url)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
	}

	results := make([]proxy.Proxy, 0)
	for _, link := range g.Results {
		data, err := proxy.ParseSSRLink(link)
		if err != nil {
			continue
		}
		results = append(results, *data)
	}
	return results
}
