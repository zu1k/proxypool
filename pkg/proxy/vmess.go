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
	ErrorNotVmessLink          = errors.New("not a correct vmess link")
	ErrorVmessPayloadParseFail = errors.New("vmess link payload parse failed")
)

type Vmess struct {
	Base
	UUID           string            `yaml:"uuid" json:"uuid"`
	AlterID        int               `yaml:"alterId" json:"alterId"`
	Cipher         string            `yaml:"cipher" json:"cipher"`
	TLS            bool              `yaml:"tls,omitempty" json:"tls,omitempty"`
	Network        string            `yaml:"network,omitempty" json:"network,omitempty"`
	HTTPOpts       HTTPOptions       `yaml:"http-opts,omitempty" json:"http-opts,omitempty"`
	WSPath         string            `yaml:"ws-path,omitempty" json:"ws-path,omitempty"`
	WSHeaders      map[string]string `yaml:"ws-headers,omitempty" json:"ws-headers,omitempty"`
	SkipCertVerify bool              `yaml:"skip-cert-verify,omitempty" json:"skip-cert-verify,omitempty"`
	ServerName     string            `yaml:"servername,omitempty" json:"servername,omitempty"`
}

type HTTPOptions struct {
	Method  string              `yaml:"method,omitempty" json:"method,omitempty"`
	Path    []string            `yaml:"path,omitempty" json:"path,omitempty"`
	Headers map[string][]string `yaml:"headers,omitempty" json:"headers,omitempty"`
}

func (v Vmess) Identifier() string {
	return net.JoinHostPort(v.Server, strconv.Itoa(v.Port)) + v.Cipher + v.UUID
}

func (v Vmess) String() string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

func (v Vmess) ToClash() string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return "- " + string(data)
}

func (v Vmess) ToSurge() string {
	// node2 = vmess, server, port, username=, ws=true, ws-path=, ws-headers=
	if v.Network == "ws" {
		wsHeasers := ""
		for k, v := range v.WSHeaders {
			if wsHeasers == "" {
				wsHeasers = k + ":" + v
			} else {
				wsHeasers += "|" + k + ":" + v
			}
		}
		text := fmt.Sprintf("%s = vmess, %s, %d, username=%s, ws=true, tls=%t, ws-path=%s",
			v.Name, v.Server, v.Port, v.UUID, v.TLS, v.WSPath)
		if wsHeasers != "" {
			text += ", ws-headers=" + wsHeasers
		}
		return text
	} else {
		return fmt.Sprintf("%s = vmess, %s, %d, username=%s, tls=%t",
			v.Name, v.Server, v.Port, v.UUID, v.TLS)
	}
}

func (v Vmess) Clone() Proxy {
	return &v
}

type vmessLinkJson struct {
	Add  string      `json:"add"`
	V    string      `json:"v"`
	Ps   string      `json:"ps"`
	Port interface{} `json:"port"`
	Id   string      `json:"id"`
	Aid  string      `json:"aid"`
	Net  string      `json:"net"`
	Type string      `json:"type"`
	Host string      `json:"host"`
	Path string      `json:"path"`
	Tls  string      `json:"tls"`
}

func ParseVmessLink(link string) (*Vmess, error) {
	if !strings.HasPrefix(link, "vmess") {
		return nil, ErrorNotVmessLink
	}

	vmessmix := strings.SplitN(link, "://", 2)
	if len(vmessmix) < 2 {
		return nil, ErrorNotVmessLink
	}
	linkPayload := vmessmix[1]
	if strings.Contains(linkPayload, "?") {
		// 使用第二种解析方法
		var infoPayloads []string
		if strings.Contains(linkPayload, "/?") {
			infoPayloads = strings.SplitN(linkPayload, "/?", 2)
		} else {
			infoPayloads = strings.SplitN(linkPayload, "?", 2)
		}
		if len(infoPayloads) < 2 {
			return nil, ErrorNotVmessLink
		}

		baseInfo, err := tool.Base64DecodeString(infoPayloads[0])
		if err != nil {
			return nil, ErrorVmessPayloadParseFail
		}
		baseInfoPath := strings.Split(baseInfo, ":")
		if len(baseInfoPath) < 3 {
			return nil, ErrorPathNotComplete
		}
		// base info
		cipher := baseInfoPath[0]
		mixInfo := strings.SplitN(baseInfoPath[1], "@", 2)
		if len(mixInfo) < 2 {
			return nil, ErrorVmessPayloadParseFail
		}
		uuid := mixInfo[0]
		server := mixInfo[1]
		portStr := baseInfoPath[2]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, ErrorVmessPayloadParseFail
		}

		moreInfo, _ := url.ParseQuery(infoPayloads[1])
		remarks := moreInfo.Get("remarks")
		obfs := moreInfo.Get("obfs")
		network := "tcp"
		if obfs == "websocket" {
			network = "ws"
		}
		//obfsParam := moreInfo.Get("obfsParam")
		path := moreInfo.Get("path")
		if path == "" {
			path = "/"
		}
		tls := moreInfo.Get("tls") == "1"

		wsHeaders := make(map[string]string)
		return &Vmess{
			Base: Base{
				Name:   remarks + "_" + strconv.Itoa(rand.Int()),
				Server: server,
				Port:   port,
				Type:   "vmess",
				UDP:    false,
			},
			UUID:           uuid,
			AlterID:        0,
			Cipher:         cipher,
			TLS:            tls,
			Network:        network,
			HTTPOpts:       HTTPOptions{},
			WSPath:         path,
			WSHeaders:      wsHeaders,
			SkipCertVerify: true,
			ServerName:     server,
		}, nil
	} else {
		payload, err := tool.Base64DecodeString(linkPayload)
		if err != nil {
			return nil, ErrorVmessPayloadParseFail
		}
		vmessJson := vmessLinkJson{}
		err = json.Unmarshal([]byte(payload), &vmessJson)
		if err != nil {
			return nil, err
		}
		port := 443
		portInterface := vmessJson.Port
		if i, ok := portInterface.(int); ok {
			port = i
		} else if s, ok := portInterface.(string); ok {
			port, _ = strconv.Atoi(s)
		}

		alterId, err := strconv.Atoi(vmessJson.Aid)
		if err != nil {
			alterId = 0
		}
		tls := vmessJson.Tls == "tls"

		wsHeaders := make(map[string]string)
		if vmessJson.Host != "" {
			wsHeaders["HOST"] = vmessJson.Host
		}

		if vmessJson.Path == "" {
			vmessJson.Path = "/"
		}
		return &Vmess{
			Base: Base{
				Name:   vmessJson.Ps + "_" + strconv.Itoa(rand.Int()),
				Server: vmessJson.Add,
				Port:   port,
				Type:   "vmess",
				UDP:    false,
			},
			UUID:           vmessJson.Id,
			AlterID:        alterId,
			Cipher:         "auto",
			TLS:            tls,
			Network:        vmessJson.Net,
			HTTPOpts:       HTTPOptions{},
			WSPath:         vmessJson.Path,
			WSHeaders:      wsHeaders,
			SkipCertVerify: true,
			ServerName:     vmessJson.Host,
		}, nil
	}
}

var (
	vmessPlainRe = regexp.MustCompile("vmess://([A-Za-z0-9+/_?&=-])+")
)

func GrepVmessLinkFromString(text string) []string {
	results := make([]string, 0)
	texts := strings.Split(text, "vmess://")
	for _, text := range texts {
		results = append(results, vmessPlainRe.FindAllString("vmess://"+text, -1)...)
	}
	return results
}
