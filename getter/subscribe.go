package getter

import "github.com/zu1k/proxypool/proxy"

type Subscribe struct {
	NumNeeded int
	Results   []string
	Url       string
}

func (s Subscribe) Get() []proxy.Proxy {
	return nil
}
