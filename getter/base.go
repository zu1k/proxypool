package getter

import "github.com/zu1k/proxypool/proxy"

type Getter interface {
	Get() []proxy.Proxy
}
