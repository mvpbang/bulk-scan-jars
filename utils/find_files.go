package utils

import (
	"log"
	"os"
	"path/filepath"
)

// 根据特征后缀找文件

func FindFiles(config Config) []string {
	var files []string

	err := filepath.Walk(config.Base.WarDir, func(path string, info os.FileInfo, err error) error {
		// 过滤后缀文件
		//log.Println(filepath.Ext(path))
		for _, ext := range config.Base.Exts {
			if filepath.Ext(path) == ext {
				files = append(files, path)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	//log.Println("++++", files)
	return files
}
