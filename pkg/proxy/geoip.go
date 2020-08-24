package proxy

import (
	"log"
	"net"
	"os"

	"github.com/oschwald/geoip2-golang"
)

var geoIp GeoIP

func InitGeoIpDB() {
	geoIp = NewGeoIP("assets/GeoLite2-City.mmdb")
}

// GeoIP2
type GeoIP struct {
	db *geoip2.Reader
}

// new geoip from db file
func NewGeoIP(filePath string) (geoip GeoIP) {
	// 判断文件是否存在
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		log.Println("文件不存在，请自行下载 Geoip2 City库，并保存在", filePath)
		os.Exit(1)
	} else {
		db, err := geoip2.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		geoip = GeoIP{db: db}
	}
	return
}

// find ip info
func (g GeoIP) Find(ipORdomain string) (ip, country string, err error) {
	ips, err := net.LookupIP(ipORdomain)
	if err != nil {
		return "", "", err
	}
	ipData := net.ParseIP(ips[0].String())
	record, err := g.db.City(ipData)
	if err != nil {
		return "", "", err
	}
	return ips[0].String(), record.Country.IsoCode, nil
}
