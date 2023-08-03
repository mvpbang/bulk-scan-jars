package utils

import (
	"archive/zip"
	"log"
	"path/filepath"
)

// 存储war、ear、jar 路径及查看包内jar

type FileJar struct {
	FilePath string
	Jars     []string
}

// 查看压缩文件，根据jar筛选出路径

func ViewFiles(file string) FileJar {
	// 存放jar
	//var jars []string
	jars := make([]string, 0)

	var filejar FileJar

	zipFile, err := zip.OpenReader(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer zipFile.Close()

	for _, f := range zipFile.File {
		if !f.FileInfo().IsDir() {
			//info := f.FileInfo().Name()
			//log.Println(f.Name, info)
			//log.Println(f.Name)
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
