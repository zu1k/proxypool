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
