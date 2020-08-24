package proxy

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/zu1k/proxypool/pkg/tool"
)

var (
	ErrorNotSSRLink             = errors.New("not a correct ssr link")
	ErrorPasswordParseFail      = errors.New("password parse failed")
	ErrorPathNotComplete        = errors.New("path not complete")
	ErrorMissingQuery           = errors.New("link missing query")
	ErrorProtocolParamParseFail = errors.New("protocol param parse failed")
	ErrorObfsParamParseFail     = errors.New("obfs param parse failed")
)

type ShadowsocksR struct {
	Base
	Password      string `yaml:"password" json:"password"`
	Cipher        string `yaml:"cipher" json:"cipher"`
	Protocol      string `yaml:"protocol" json:"protocol"`
	ProtocolParam string `yaml:"protocol-param,omitempty" json:"protocol_param,omitempty"`
	Obfs          string `yaml:"obfs" json:"obfs"`
	ObfsParam     string `yaml:"obfs-param,omitempty" json:"obfs_param,omitempty"`
	Group         string `yaml:"group,omitempty" json:"group,omitempty"`
}

func (ssr ShadowsocksR) Identifier() string {
	return net.JoinHostPort(ssr.Server, strconv.Itoa(ssr.Port)) + ssr.Password + ssr.ProtocolParam
}

func (ssr ShadowsocksR) String() string {
	data, err := json.Marshal(ssr)
	if err != nil {
		return ""
	}
	return string(data)
}

func (ssr ShadowsocksR) ToClash() string {
	data, err := json.Marshal(ssr)
	if err != nil {
		return ""
	}
	return "- " + string(data)
}

func (ssr ShadowsocksR) ToSurge() string {
	return ""
}

func (ssr ShadowsocksR) Clone() Proxy {
	return &ssr
}

func ParseSSRLink(link string) (*ShadowsocksR, error) {
	if !strings.HasPrefix(link, "ssr") {
		return nil, ErrorNotSSRLink
	}

	ssrmix := strings.SplitN(link, "://", 2)
	if len(ssrmix) < 2 {
		return nil, ErrorNotSSRLink
	}
	linkPayloadBase64 := ssrmix[1]
	payload, err := tool.Base64DecodeString(linkPayloadBase64)
	if err != nil {
		return nil, ErrorMissingQuery
	}

	infoPayload := strings.SplitN(payload, "/?", 2)
	if len(infoPayload) < 2 {
		return nil, ErrorNotSSRLink
	}
	ssrpath := strings.Split(infoPayload[0], ":")
	if len(ssrpath) < 6 {
		return nil, ErrorPathNotComplete
	}
	// base info
	server := strings.ToLower(ssrpath[0])
	port, _ := strconv.Atoi(ssrpath[1])
	protocol := strings.ToLower(ssrpath[2])
	cipher := strings.ToLower(ssrpath[3])
	obfs := strings.ToLower(ssrpath[4])
	password, err := tool.Base64DecodeString(ssrpath[5])
	if err != nil {
		return nil, ErrorPasswordParseFail
	}

	moreInfo, _ := url.ParseQuery(infoPayload[1])

	// remarks
	remarks := moreInfo.Get("remarks")
	remarks, err = tool.Base64DecodeString(remarks)
	if err != nil {
		remarks = ""
		err = nil
	}
	if strings.ContainsAny(remarks, "\t\r\n ") {
		remarks = strings.ReplaceAll(remarks, "\t", "")
		remarks = strings.ReplaceAll(remarks, "\r", "")
		remarks = strings.ReplaceAll(remarks, "\n", "")
		remarks = strings.ReplaceAll(remarks, " ", "")
	}

	// protocol param
	protocolParam, err := tool.Base64DecodeString(moreInfo.Get("protoparam"))
	if err != nil {
		return nil, ErrorProtocolParamParseFail
	}
	if tool.ContainChineseChar(protocolParam) {
		protocolParam = ""
	}
	if strings.HasSuffix(protocol, "_compatible") {
		protocol = strings.ReplaceAll(protocol, "_compatible", "")
	}

	// obfs param
	obfsParam, err := tool.Base64DecodeString(moreInfo.Get("obfsparam"))
	if err != nil {
		return nil, ErrorObfsParamParseFail
	}
	if tool.ContainChineseChar(obfsParam) {
		obfsParam = ""
	}
	if strings.HasSuffix(obfs, "_compatible") {
		obfs = strings.ReplaceAll(obfs, "_compatible", "")
	}

	//group, err := tool.Base64DecodeString(moreInfo.Get("group"))
	//if err != nil {
	//	group = ""
	//}
	group := ""

	return &ShadowsocksR{
		Base: Base{
			Name:   remarks + "_" + strconv.Itoa(rand.Int()),
			Server: server,
			Port:   port,
			Type:   "ssr",
		},
		Password:      password,
		Cipher:        cipher,
		Protocol:      protocol,
		ProtocolParam: protocolParam,
		Obfs:          obfs,
		ObfsParam:     obfsParam,
		Group:         group,
	}, nil
}

var (
	ssrPlainRe = regexp.MustCompile("ssr://([A-Za-z0-9+/_-])+")
)

func GrepSSRLinkFromString(text string) []string {
	results := make([]string, 0)
	texts := strings.Split(text, "ssr://")
	for _, text := range texts {
		results = append(results, ssrPlainRe.FindAllString("ssr://"+text, -1)...)
	}
	return results
}
