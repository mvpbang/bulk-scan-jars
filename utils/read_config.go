package utils

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Base *Base    `yml:"base"`
	Eq   []string `yml:"eq"`
	Le   []string `yml:"le"`
}

type Base struct {
	WarDir string   `yml:"wardir"`
	Result string   `yml:"result"`
	Exts   []string `yml:"exts"`
}

// ReadConfig 读取配置
func ReadConfig(config *Config) {
	// 读取yml
	ymlFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Println(err)
		return
	}

	// 解码
	err = yaml.Unmarshal(ymlFile, &config)
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println(config.Base.WarDir)
	//log.Println(config)
}
