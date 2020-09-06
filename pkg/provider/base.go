package provider

import (
	"strings"

	"github.com/zu1k/proxypool/pkg/proxy"
)

type Provider interface {
	Provide() string
}

type Base struct {
	Proxies    *proxy.ProxyList `yaml:"proxies"`
	Types      string           `yaml:"type"`
	Country    string           `yaml:"country"`
	NotCountry string           `yaml:"not_country"`
}

func (b *Base) preFilter() {
	proxies := make(proxy.ProxyList, 0)

	needFilterType := true
	needFilterCountry := true
	needFilterNotCountry := true
	if b.Types == "" || b.Types == "all" {
		needFilterType = false
	}
	if b.Country == "" || b.Country == "all" {
		needFilterCountry = false
	}
	if b.NotCountry == "" {
		needFilterNotCountry = false
	}
	types := strings.Split(b.Types, ",")
	countries := strings.Split(b.Country, ",")
	notCountries := strings.Split(b.NotCountry, ",")

	bProxies := *b.Proxies
	for _, p := range bProxies {
		if needFilterType {
			typeOk := false
			for _, t := range types {
				if p.TypeName() == t {
					typeOk = true
					break
				}
			}
			if !typeOk {
				goto exclude
			}
		}

		if needFilterNotCountry {
			for _, c := range notCountries {
				if strings.Contains(p.BaseInfo().Name, c) {
					goto exclude
				}
			}
		}

		if needFilterCountry {
			countryOk := false
			for _, c := range countries {
				if strings.Contains(p.BaseInfo().Name, c) {
					countryOk = true
					break
				}
			}
			if !countryOk {
				goto exclude
			}
		}

		proxies = append(proxies, p)
	exclude:
	}

	b.Proxies = &proxies
}
