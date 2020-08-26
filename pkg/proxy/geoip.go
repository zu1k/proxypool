package proxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/oschwald/geoip2-golang"
	bingeoip "github.com/zu1k/proxypool/internal/bindata/geoip"
)

var geoIp GeoIP

func InitGeoIpDB() {
	err := bingeoip.RestoreAsset("", "assets/GeoLite2-City.mmdb")
	if err != nil {
		panic(err)
	}
	geoIp = NewGeoIP("assets/GeoLite2-City.mmdb")
}

// GeoIP2
type GeoIP struct {
	db    *geoip2.Reader
	flags cclist
}

type CountryCode struct {
	Code     string `json:"code"`
	Emoji    string `json:"emoji"`
	Unicode  string `json:"unicode`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Dialcode string `json:"dialCode`
}

// new geoip from db file
func NewGeoIP(filePath string) (geoip GeoIP) {
	var countrycodes cclist
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
		data, err := ioutil.ReadFile("assets/flags.json")
		if err != nil {
			log.Fatal(err)
			return
		}

		json.Unmarshal(data, &countrycodes)
		geoip = GeoIP{db: db, flags: countrycodes}
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

	countrycode := record.Country.IsoCode
	if countrycode == "" {
		// handle unknown country code
		return ips[0].String(), fmt.Sprintf("🏁 ZZ"), nil
	}
	countryflag := g.getFlag(countrycode)
	country = fmt.Sprintf("%v %v", countryflag, countrycode)

	return ips[0].String(), country, nil
}

// getFlag method take country code as input, return its corresponding country/region flag
func (g GeoIP) getFlag(countrycode string) string {
	result := find(g.flags, countrycode)
	return result.Emoji
}

type cclist []CountryCode

func (c cclist) in(code string, left bool) bool {
	length := len(c)
	if left {
		return c[length-1].Code >= code
	}
	return c[0].Code <= code
}

// find will find corresponding country flag emoji from flags for a given country code
func find(list cclist, target string) CountryCode {
	var result CountryCode

	length := len(list)
	if length == 1 {
		if list[0].Code == target {
			return list[0]
		}
		return CountryCode{Emoji: "🏁"}
	}
	split := length / 2
	left, right := list[:split], list[split:]

	if left.in(target, true) {
		result = find(left, target)
	}
	if right.in(target, false) {
		result = find(right, target)
	}

	return result
}
