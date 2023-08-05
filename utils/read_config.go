package utils

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

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
}
