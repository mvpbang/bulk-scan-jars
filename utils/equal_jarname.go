package utils

import (
	"regexp"
	"strings"
)

func EqualJarName(filejar FileJar, config Config) (jf FileJar) {

	for _, jarpath := range filejar.Jars {
		for _, eqjarname := range config.Eq {
			eqstr := strings.Split(eqjarname, ":")[0]
			// 匹配到则加到结果集
			if regexp.MustCompile(eqstr).MatchString(jarpath) {
				jf.Jars = append(jf.Jars, jarpath)
			}
		}
		for _, lejarname := range config.Le {
			lestr := strings.Split(lejarname, ":")[0]
			// 匹配到则加到结果集
			if regexp.MustCompile(lestr).MatchString(jarpath) {
				jf.Jars = append(jf.Jars, jarpath)
			}
		}
	}
	// 去重 jars,待实现

	jf.FilePath = filejar.FilePath
	//log.Println("++++++", jf)
	return jf
}
