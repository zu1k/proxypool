package app

import "github.com/zu1k/proxypool/proxy"

var (
	GeoIp       proxy.GeoIP
	ProjectName = "proxypool"
)

func init() {
	GeoIp = proxy.NewGeoIP("assets/GeoLite2-City.mmdb")
}
