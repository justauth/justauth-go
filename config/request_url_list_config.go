package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	requestUrlListConfig = RequestUrlListConfig{}
)

type RequestUrlListConfig struct {
	RequestUrlList []PlatformUrl `yaml:"RequestUrlList"`
}

type PlatformUrl struct {
	Platform       string `yaml:"Platform"`
	AuthorizeUrl   string `yaml:"AuthorizeUrl"`
	AccessTokenUrl string `yaml:"AccessTokenUrl"`
	UserInfoUrl    string `yaml:"UserInfoUrl"`
	RefreshUrl     string `yaml:"RefreshUrl"`
}

type PlatformStr string

const (
	PlatformStrGitee    PlatformStr = "gitee"
	PlatformStrWechatMp PlatformStr = "wechat_mp"
	PlatformStrTwitter  PlatformStr = "twitter"
	PlatformStrAliPay   PlatformStr = "alipay"
	PlatformStrHuaWei   PlatformStr = "huawei"
)

type UserGender string

const (
	UserGenderMan     UserGender = "男"
	UserGenderWoman   UserGender = "女"
	UserGenderUnknown UserGender = "未知"
)

func InitConfig() error {
	fileContent, err := ioutil.ReadFile("config/request_url.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(fileContent, &requestUrlListConfig)
	if err != nil {
		return err
	}
	return nil
}

func GetRequestUrlListConfig() (s *RequestUrlListConfig) {
	return &requestUrlListConfig
}

func (config *RequestUrlListConfig) GetUrlByPlatform(platformStr PlatformStr) (bool, PlatformUrl) {
	for _, platform := range config.RequestUrlList {
		if platform.Platform == string(platformStr) {
			return true, platform
		}
	}
	return false, PlatformUrl{}
}
