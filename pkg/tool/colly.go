package tool

import (
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly"
)

func GetColly() *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(UserAgent),
	)
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second, // 超时时间
			KeepAlive: 10 * time.Second, // keepAlive 超时时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       20 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 10 * time.Second,
	})
	return c
}
