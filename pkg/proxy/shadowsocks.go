package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/zu1k/proxypool/pkg/tool"
)

var (
	ErrorNotSSLink = errors.New("not a correct ss link")
)

type Shadowsocks struct {
	Base
	Password   string                 `yaml:"password" json:"password"`
	Cipher     string                 `yaml:"cipher" json:"cipher"`
	Plugin     string                 `yaml:"plugin,omitempty" json:"plugin,omitempty"`
	PluginOpts map[string]interface{} `yaml:"plugin-opts,omitempty" json:"plugin-opts,omitempty"`
}

func (ss Shadowsocks) Identifier() string {
	return net.JoinHostPort(ss.Server, strconv.Itoa(ss.Port)) + ss.Password
}

func (ss Shadowsocks) String() string {
	data, err := json.Marshal(ss)
	if err != nil {
		return ""
	}
	return string(data)
}

func (ss Shadowsocks) ToClash() string {
	data, err := json.Marshal(ss)
	if err != nil {
		return ""
	}
	return "- " + string(data)
}

func (ss Shadowsocks) ToSurge() string {
	// node1 = ss, server, port, encrypt-method=, password=, obfs=, obfs-host=, udp-relay=false
	if ss.Plugin == "obfs" {
		text := fmt.Sprintf("%s = ss, %s, %d, encrypt-method=%s, password=%s, obfs=%s, udp-relay=false",
			ss.Name, ss.Server, ss.Port, ss.Cipher, ss.Password, ss.PluginOpts["mode"])
		if ss.PluginOpts["host"].(string) != "" {
			text += ", obfs-host=" + ss.PluginOpts["host"].(string)
		}
		return text
	} else {
		return fmt.Sprintf("%s = ss, %s, %d, encrypt-method=%s, password=%s, udp-relay=false",
			ss.Name, ss.Server, ss.Port, ss.Cipher, ss.Password)
	}
}

func (ss Shadowsocks) Clone() Proxy {
	return &ss
}

func ParseSSLink(link string) (*Shadowsocks, error) {
	if !strings.HasPrefix(link, "ss://") {
		return nil, ErrorNotSSRLink
	}

	uri, err := url.Parse(link)
	if err != nil {
		return nil, ErrorNotSSLink
	}

	cipher := ""
	password := ""
	if uri.User.String() == "" {
		// base64的情况
		infos, err := tool.Base64DecodeString(uri.Hostname())
		if err != nil {
			return nil, err
		}
		uri, err = url.Parse("ss://" + infos)
		if err != nil {
			return nil, err
		}
		cipher = uri.User.Username()
		password, _ = uri.User.Password()
	} else {
		cipherInfoString, err := tool.Base64DecodeString(uri.User.Username())
		if err != nil {
			return nil, ErrorPasswordParseFail
		}
		cipherInfo := strings.SplitN(cipherInfoString, ":", 2)
		if len(cipherInfo) < 2 {
			return nil, ErrorPasswordParseFail
		}
		cipher = strings.ToLower(cipherInfo[0])
		password = cipherInfo[1]
	}
	server := uri.Hostname()
	port, _ := strconv.Atoi(uri.Port())

	moreInfos := uri.Query()
	pluginString := moreInfos.Get("plugin")
	plugin := ""
	pluginOpts := make(map[string]interface{})
	if strings.Contains(pluginString, ";") {
		pluginInfos, err := url.ParseQuery(pluginString)
		if err == nil {
			if strings.Contains(pluginString, "obfs") {
				plugin = "obfs"
				pluginOpts["mode"] = pluginInfos.Get("obfs")
				pluginOpts["host"] = pluginInfos.Get("obfs-host")
			} else if strings.Contains(pluginString, "v2ray") {
				plugin = "v2ray-plugin"
				pluginOpts["mode"] = pluginInfos.Get("mode")
				pluginOpts["host"] = pluginInfos.Get("host")
				pluginOpts["tls"] = strings.Contains(pluginString, "tls")
			}
		}
	}
	if port == 0 || cipher == "" {
		return nil, ErrorNotSSLink
	}

	return &Shadowsocks{
		Base: Base{
			Name:   strconv.Itoa(rand.Int()),
			Server: server,
			Port:   port,
			Type:   "ss",
		},
		Password:   password,
		Cipher:     cipher,
		Plugin:     plugin,
		PluginOpts: pluginOpts,
	}, nil
}

var (
	ssPlainRe = regexp.MustCompile("ss://([A-Za-z0-9+/_&?=@:%.-])+")
)

func GrepSSLinkFromString(text string) []string {
	results := make([]string, 0)
	texts := strings.Split(text, "ss://")
	for _, text := range texts {
		results = append(results, ssPlainRe.FindAllString("ss://"+text, -1)...)
	}
	return results
}
