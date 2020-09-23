package rekpkg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Config struct {
	Red  ColorR `json:"red"`
	Open ColorR `json:"open"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) Get(path string) (*Config, error) {
	var config *Config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config, errors.New("读取配置文件失败")
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, errors.New("读取配置文件失败")
	}
	return config, err
}
