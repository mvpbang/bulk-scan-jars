package utils

import (
	"regexp"
)

type JarInfo struct {
	ParPath string
	JarPath string
	Name    string
	Ver     string
}

func ExtractNameVer(filejar FileJar) []JarInfo {

	//func ExtractNameVer(jarpath string) JarInfo {
	//jarPath := "WEB-INF/lib/BondeReverseDesensitize-0.0.25-SNAPSHOT.jar"
	//jarPath := "WEB-INF/lib/nacos-spring-context-0.0.2.jar"
	//jarPath := []string{
	//	"WEB-INF/lib/nacos-spring-context-0.0.2.jar",
	//	"WEB-INF/lib/BondeReverseDesensitize-0.0.25-SNAPSHOT.jar",
	//	"WEB-INF/lib/nacos-spring-context-0.0.2-20230718-SNAPSHOT.jar",
	//	"WEB-INF/lib/x-spring-context1-0.0.2-20230718-SNAPSHOT.jar",
	//}

	var jarinfo JarInfo
	var jarinfos = make([]JarInfo, 0)

	rJar := `.*/(?P<jarname>.*)`
	rName := `^(?P<name>.*?)-\d`
	rVer := `\b(?:\d+\.?)+\-?\d+|\d`

	rjar := regexp.MustCompile(rJar)
	rname := regexp.MustCompile(rName)
	rver := regexp.MustCompile(rVer)

	for _, filepath := range filejar.Jars {
		rjarv := rjar.FindStringSubmatch(filepath)
		rnamev := rname.FindStringSubmatch(rjarv[1])
		rverv := rver.FindStringSubmatch(rjarv[1])

		jarinfo = JarInfo{
			ParPath: filejar.FilePath,
			JarPath: filepath,
			Name:    rnamev[1],
			Ver:     rverv[0],
		}
		jarinfos = append(jarinfos, jarinfo)

		//log.Println("+++++", jarinfo[i])
	}

	/*	for _, jarpath := range jarPath {
		rjarv := rjar.FindStringSubmatch(jarpath)
		rnamev := rname.FindStringSubmatch(rjarv[1])
		rverv := rver.FindStringSubmatch(rjarv[1])
		log.Println(rnamev[1], rverv)
	}*/
	return jarinfos
}
