package utils

import (
	"archive/zip"
	"log"
	"path/filepath"
)

// 查看压缩文件，记录jar路径

func ViewFiles(file string) FileJar {
	// 存放jar
	//jars := make([]string, 0)
	var jars []string
	var filejar FileJar

	// zip形式打开文件
	zipFile, err := zip.OpenReader(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer zipFile.Close()

	for _, f := range zipFile.File {
		// 选出文件jar的路径
		if !f.FileInfo().IsDir() {
			if filepath.Ext(f.Name) == ".jar" {
				//log.Println(f.Name)
				jars = append(jars, f.Name)
			}
		}
	}

	filejar.FilePath = file
	filejar.Jars = jars

	return filejar
}
