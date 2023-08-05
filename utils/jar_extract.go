package utils

import (
	"log"
	"regexp"
	"strings"
)

func ExtractNameVer(filejar FileJar) []JarInfo {

	var jarinfo JarInfo
	//var jarinfos = make([]JarInfo, 0)
	var jarinfos []JarInfo
	var rn, rv string

	// 正则提取规则
	rJar := `.*/(?P<jarname>.*)`
	rName := `^(?P<name>.*?)-\d`
	//rVer := `\b(?:\d+\.?)+\-?\d+|\b\d`
	rVer := `-\b(?:\d+\.?)+\-?\d+|-\d`

	rjar := regexp.MustCompile(rJar)
	rname := regexp.MustCompile(rName)
	rver := regexp.MustCompile(rVer)

	for _, filepath := range filejar.Jars {
		rjarv := rjar.FindStringSubmatch(filepath)
		rnamev := rname.FindStringSubmatch(rjarv[1])
		rverv := rver.FindStringSubmatch(rjarv[1])

		//log.Println(rjarv, rverv)
		// 对特殊jar提取名字 + 版本 异常处理
		if len(rverv) == 0 || len(rnamev) == 0 {
			log.Printf("警告警告: %v,%v,%v,%v,%v", filejar.FilePath, filepath, rjarv, rnamev, rverv)
			rn = "gaga"
			rv = "9.9.9"
		} else {
			rv = strings.TrimLeft(rverv[0], "-")
			rn = rnamev[1]
		}

		jarinfo = JarInfo{
			ParPath: filejar.FilePath,
			JarPath: filepath,
			Name:    rn,
			Ver:     rv,
		}

		//log.Println("+++", jarinfo)
		jarinfos = append(jarinfos, jarinfo)
	}

	return jarinfos
}
