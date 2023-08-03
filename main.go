package main

import (
	"bulk-scan-jars/utils"
	"io"
	"log"
	"os"
)

func init() {
	logFile2, err := os.OpenFile(".//check.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile2)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(mw)
}

func main() {
	var config utils.Config

	//读取配置反序列化yml
	utils.ReadConfig(&config)
	//log.Println(config)

	files := utils.FindFiles(config)

	// 记录war、jar、ear 和对应jar的关系
	for _, file := range files {
		filejar := utils.ViewFiles(file)
		jarinfo := utils.ExtractNameVer(filejar)
		for _, jf := range jarinfo {
			utils.CompareVer(config, jf)
		}
	}
}
