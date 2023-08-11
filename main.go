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
		headstr := "代码,判断类型,成果路径,jar路径,jar包,比较结果,jar包安全版本\n"
		logfile.WriteString(headstr)
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
		// 过滤jar后缀
		filejar := utils.ViewFiles(file)
		// 再次过滤仅仅选出符合搜索的jar
		fj := utils.EqualJarName(filejar, config)
		// 提取jarname、jarver
		jarinfo := utils.ExtractNameVer(fj)
		for _, jf := range jarinfo {
			// 判断比较
			utils.CompareVer(config, jf)
		}
	}
}
