package main

import (
	"bulk-scan-jars/utils"
	"io"
	"log"
	"os"
)

func init() {
	var logfile *os.File
	filename := ".//no_pass.txt"

	// 文件不存在
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		logfile, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Panicln(err)
		}
		logfile.WriteString("代码,判断类型,成果路径,jar路径,jar包,比较结果,jar包安全版本\n")
	} else {
		logfile, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Panicln(err)
		}
	}

	//logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	log.Panicln(err)
	//}

	// 多写(console、file)
	mw := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(mw)
}

func main() {
	//log.Println("嘎嘎")
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
