package getter

import (
	"fmt"
	"sync"

	"github.com/zu1k/proxypool/tool"

	"github.com/gocolly/colly"
	"github.com/zu1k/proxypool/proxy"
)

func init() {
	Register("tgchannel", NewTGChannelGetter)
}

type TGChannelGetter struct {
	c         *colly.Collector
	NumNeeded int
	results   []string
	Url       string
}

func NewTGChannelGetter(options tool.Options) Getter {
	num, found := options["num"]
	if !found || int(num.(float64)) <= 0 {
		num = 200
	}
	url, found := options["channel"]
	if found {
		return &TGChannelGetter{
			c:         colly.NewCollector(),
			NumNeeded: int(num.(float64)),
			Url:       "https://t.me/s/" + url.(string),
		}
	}
	return nil
}

func (g *TGChannelGetter) Get() []proxy.Proxy {
	g.results = make([]string, 0)
	// 找到所有的文字消息
	g.c.OnHTML("div.tgme_widget_message_text", func(e *colly.HTMLElement) {
		g.results = append(g.results, GrepLinksFromString(e.Text)...)
	})

	// 找到之前消息页面的链接，加入访问队列
	g.c.OnHTML("link[rel=prev]", func(e *colly.HTMLElement) {
		if len(g.results) < g.NumNeeded {
			_ = e.Request.Visit(e.Attr("href"))
		}
	})

	g.results = make([]string, 0)
	err := g.c.Visit(g.Url)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
	}

	return StringArray2ProxyArray(g.results)
}

func (g *TGChannelGetter) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	defer wg.Done()
	nodes := g.Get()
	for _, node := range nodes {
		pc <- node
	}
}
