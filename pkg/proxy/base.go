package proxy

import (
	"errors"
	"strings"
)

type Base struct {
	Name    string `yaml:"name" json:"name" gorm:"index"`
	Server  string `yaml:"server" json:"server" gorm:"index"`
	Port    int    `yaml:"port" json:"port" gorm:"index"`
	Type    string `yaml:"type" json:"type" gorm:"index"`
	UDP     bool   `yaml:"udp,omitempty" json:"udp,omitempty"`
	Country string `yaml:"country,omitempty" json:"country,omitempty" gorm:"index"`
	Useable bool   `yaml:"useable,omitempty" json:"useable,omitempty" gorm:"index"`
}

func (b *Base) TypeName() string {
	if b.Type == "" {
		return "unknown"
	}
	return b.Type
}

func (b *Base) SetName(name string) {
	b.Name = name
}

func (b *Base) SetIP(ip string) {
	b.Server = ip
}

func (b *Base) BaseInfo() *Base {
	return b
}

func (b *Base) Clone() Base {
	c := *b
	return c
}

func (b *Base) SetUseable(useable bool) {
	b.Useable = useable
}

func (b *Base) SetCountry(country string) {
	b.Country = country
}

type Proxy interface {
	String() string
	ToClash() string
	ToSurge() string
	Link() string
	Identifier() string
	SetName(name string)
	SetIP(ip string)
	TypeName() string
	BaseInfo() *Base
	Clone() Proxy
	SetUseable(useable bool)
	SetCountry(country string)
}

func ParseProxyFromLink(link string) (p Proxy, err error) {
	if strings.HasPrefix(link, "ssr://") {
		p, err = ParseSSRLink(link)
	} else if strings.HasPrefix(link, "vmess://") {
		p, err = ParseVmessLink(link)
	} else if strings.HasPrefix(link, "ss://") {
		p, err = ParseSSLink(link)
	} else if strings.HasPrefix(link, "trojan://") {
		p, err = ParseTrojanLink(link)
	} else if strings.HasPrefix(link, "- {") {
		p, err = ParseSurgeLink(link)
	}
	// else if strings.HasPrefix(link, "[{") && strings.HasSuffix(link, "}]") {
	// 	fmt.Println("$$$$$$ " + link + " ######")
	// 	p = nil
	// }

	if err != nil || p == nil {
		return nil, errors.New("link parse failed")
	}
	ip, country, err := geoIp.Find(p.BaseInfo().Server)
	if err != nil {
		country = "🏁 ZZ"
	}
	p.SetCountry(country)
	// trojan依赖域名？
	if p.TypeName() != "trojan" {
		p.SetIP(ip)
	}
	return
}
