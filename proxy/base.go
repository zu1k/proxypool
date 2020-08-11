package proxy

import "C"

type Base struct {
	Name   string `yaml:"name" json:"name"`
	Server string `yaml:"server" json:"server"`
	Port   int    `yaml:"port" json:"port"`
	Type   string `yaml:"type" json:"type"`
	UDP    bool   `yaml:"udp,omitempty" json:"udp,omitempty"`
}

type Proxy interface {
	String() string
	ToClash() string
	Identifier() string
}

func Deduplication(src []Proxy) []Proxy {
	result := make([]Proxy, 0, len(src))
	temp := map[string]struct{}{}
	for _, item := range src {
		if _, ok := temp[item.Identifier()]; !ok {
			temp[item.Identifier()] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
