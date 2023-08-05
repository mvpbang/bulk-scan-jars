package main

import (
	"bulk-scan-jars/utils"
	"io"
	"log"
	"os"
)

func init() {
	logfile, err := os.OpenFile(".//check.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
	}
	// 多写(console、file)
	mw := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(mw)
}

func main() {
	var config utils.Config

	// 读取配置
	utils.ReadConfig(&config)

	// 符合后缀的文件
	files := utils.FindFiles(config)

	// 遍历比较
	for _, file := range files {
		filejar := utils.ViewFiles(file)
		jarinfo := utils.ExtractNameVer(filejar)
		for _, jf := range jarinfo {
			utils.CompareVer(config, jf)
		}
	}
}
