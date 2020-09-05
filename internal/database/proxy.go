package database

import (
	"github.com/zu1k/proxypool/pkg/proxy"
	"gorm.io/gorm"
)

type Proxy struct {
	gorm.Model
	proxy.Base
	Link       string
	Identifier string `gorm:"unique"`
}

func InitTables() {
	if DB == nil {
		err := connect()
		if err != nil {
			return
		}
	}
	err := DB.AutoMigrate(&Proxy{})
	if err != nil {
		panic(err)
	}
}

func SaveProxyList(pl proxy.ProxyList) {
	if DB == nil {
		return
	}
	proxies := make([]Proxy, pl.Len())
	for i, p := range pl {
		proxies[i] = Proxy{
			Base:       *p.BaseInfo(),
			Link:       p.Link(),
			Identifier: p.Identifier(),
		}
	}
	DB.Create(&proxies)
}
