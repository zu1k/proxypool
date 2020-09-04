package proxy

import (
	"fmt"
	"testing"
)

func TestSSLink(t *testing.T) {
	ss, err := ParseSSLink("ss://YWVzLTI1Ni1jZmI6ZUlXMERuazY5NDU0ZTZuU3d1c3B2OURtUzIwMXRRMERAMTcyLjEwNC4xNjEuNTQ6ODA5OQ==#翻墙党223.13新加坡")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ss)
	fmt.Println(ss.Link())
	ss, err = ParseSSLink(ss.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ss)
}

func TestSSRLink(t *testing.T) {
	ssr, err := ParseSSRLink("ssr://MTcyLjEwNC4xNjEuNTQ6ODA5OTpvcmlnaW46YWVzLTI1Ni1jZmI6cGxhaW46WlVsWE1FUnVhelk1TkRVMFpUWnVVM2QxYzNCMk9VUnRVekl3TVhSUk1FUT0vP29iZnNwYXJhbT0mcHJvdG9wYXJhbT0mcmVtYXJrcz01Ny03NWFLWjVZV2FNakl6TGpFejVwYXc1WXFnNVoyaCZncm91cD01cGF3NVlxZzVaMmg=")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ssr)
	fmt.Println(ssr.Link())
	ssr, err = ParseSSRLink(ssr.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ssr)
}

func TestTrojanLink(t *testing.T) {
	trojan, err := ParseTrojanLink("trojan://65474277@sqcu.hostmsu.ru:55551?allowinsecure=0&peer=mza.hkfq.xyz&mux=1&ws=0&wspath=&wshost=&ss=0&ssmethod=aes-128-gcm&sspasswd=&group=#%E9%A6%99%E6%B8%AFCN2-MZA%E8%8A%82%E7%82%B9-%E5%AE%BF%E8%BF%81%E8%81%94%E9%80%9A%E4%B8%AD%E8%BD%AC")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(trojan)
	fmt.Println(trojan.Link())
	trojan, err = ParseTrojanLink(trojan.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(trojan)
}

func TestVmessLink(t *testing.T) {
	v, err := ParseVmessLink("vmess://ew0KICAidiI6ICIyIiwNCiAgInBzIjogIuW+ruS/oeWFrOS8l+WPtyDlpJrlvannmoTlpKfljYPkuJbnlYwiLA0KICAiYWRkIjogInMyNzEuc25vZGUueHl6IiwNCiAgInBvcnQiOiAiNDQzIiwNCiAgImlkIjogIjZhOTAwZDYzLWNiOTItMzVhMC1hZWYwLTNhMGMxMWFhODUyMyIsDQogICJhaWQiOiAiMSIsDQogICJuZXQiOiAid3MiLA0KICAidHlwZSI6ICJub25lIiwNCiAgImhvc3QiOiAiczI3MS5zbm9kZS54eXoiLA0KICAicGF0aCI6ICIvcGFuZWwiLA0KICAidGxzIjogInRscyINCn0=")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
	fmt.Println(v.Link())
	v, err = ParseVmessLink(v.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
