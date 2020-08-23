package config

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/zu1k/proxypool/pkg/tool"
)

var configFilePath = "config.yaml"

type ConfigOptions struct {
	Domain      string   `json:"domain" yaml:"domain"`
	CFEmail     string   `json:"cf_email" yaml:"cf_email"`
	CFKey       string   `json:"cf_key" yaml:"cf_key"`
	SourceFiles []string `json:"source-files" yaml:"source-files"`
}

// Config 配置
var Config ConfigOptions

// Parse 解析配置文件，支持本地文件系统和网络链接
func Parse(path string) error {
	if path == "" {
		path = configFilePath
	} else {
		configFilePath = path
	}
	fileData, err := ReadFile(path)
	if err != nil {
		return err
	}
	Config = ConfigOptions{}
	err = yaml.Unmarshal(fileData, &Config)
	if err != nil {
		return err
	}

	// 部分配置环境变量优先
	if domain := os.Getenv("DOMAIN"); domain != "" {
		Config.Domain = domain
	}
	if cfEmail := os.Getenv("CF_API_EMAIL"); cfEmail != "" {
		Config.CFEmail = cfEmail
	}
	if cfKey := os.Getenv("CF_API_KEY"); cfKey != "" {
		Config.CFKey = cfKey
	}

	return nil
}

// 从本地文件或者http链接读取配置文件内容
func ReadFile(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := tool.GetHttpClient().Get(path)
		if err != nil {
			return nil, errors.New("config file http get fail")
		}
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	} else {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return nil, err
		}
		return ioutil.ReadFile(path)
	}
}
