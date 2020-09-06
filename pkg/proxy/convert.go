package proxy

import (
	"errors"

	"github.com/zu1k/proxypool/pkg/tool"
)

func SS2SSR(ss *Shadowsocks) (ssr *ShadowsocksR, err error) {
	if ss == nil {
		return nil, errors.New("ss is nil")
	}
	if !tool.CheckInList(SSRCipherList, ss.Cipher) {
		return nil, errors.New("cipher not support")
	}
	base := ss.Base
	base.Type = "ssr"
	return &ShadowsocksR{
		Base:     base,
		Password: ss.Password,
		Cipher:   ss.Cipher,
		Protocol: "origin",
		Obfs:     "plain",
		Group:    "proxy.tgbot.co",
	}, nil
}

func SSR2SS(ssr *ShadowsocksR) (ss *Shadowsocks, err error) {
	if ssr == nil {
		return nil, errors.New("ssr is nil")
	}
	if !tool.CheckInList(SSCipherList, ssr.Cipher) {
		return nil, errors.New("cipher not support")
	}
	if ssr.Protocol != "origin" || ssr.Obfs != "plain" {
		return nil, errors.New("protocol or obfs not allowed")
	}
	base := ssr.Base
	base.Type = "ss"
	return &Shadowsocks{
		Base:       base,
		Password:   ssr.Password,
		Cipher:     ssr.Cipher,
		Plugin:     "",
		PluginOpts: nil,
	}, nil
}

var SSRCipherList = []string{
	"aes-128-cfb",
	"aes-192-cfb",
	"aes-256-cfb",
	"aes-128-ctr",
	"aes-192-ctr",
	"aes-256-ctr",
	"aes-128-ofb",
	"aes-192-ofb",
	"aes-256-ofb",
	"des-cfb",
	"bf-cfb",
	"cast5-cfb",
	"rc4-md5",
	"chacha20-ietf",
	"salsa20",
	"camellia-128-cfb",
	"camellia-192-cfb",
	"camellia-256-cfb",
	"idea-cfb",
	"rc2-cfb",
	"seed-cfb",
}

var SSCipherList = []string{
	"aes-128-gcm",
	"aes-192-gcm",
	"aes-256-gcm",
	"aes-128-cfb",
	"aes-192-cfb",
	"aes-256-cfb",
	"aes-128-ctr",
	"aes-192-ctr",
	"aes-256-ctr",
	"rc4-md5",
	"chacha20-ietf",
	"xchacha20",
	"chacha20-ietf-poly1305",
	"xchacha20-ietf-poly1305",
}
