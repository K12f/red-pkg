package rekpkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Red  ColorR
	Open ColorR
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Parse(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("读取配置文件失败", err)
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		log.Fatal("解析配置文件失败", err)
	}
}
