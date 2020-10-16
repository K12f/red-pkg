package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Plat struct {
	Wechat Config `json:"wehcat"`
	Feishu Config `json:"feishu"`
}
type Config struct {
	Red  ColorR `json:"red"`
	Open ColorR `json:"open"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) Get(plat string, path string) (Config, error) {
	var PlatConfig *Plat
	var config Config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config, errors.New("读取配置文件失败")
	}
	err = json.Unmarshal(data, &PlatConfig)
	if err != nil {
		return config, errors.New("读取配置文件失败")
	}

	switch plat {
	case "1":
		config = PlatConfig.Wechat
	case "2":
		config = PlatConfig.Feishu
	default:
		err = errors.New("not found plat")
	}
	return config, err
}
