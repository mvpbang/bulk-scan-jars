package main

import (
	"log"
	"regexp"
)

func main() {
	//jarPath := "WEB-INF/lib/BondeReverseDesensitize-0.0.25-SNAPSHOT.jar"
	//jarPath := "WEB-INF/lib/nacos-spring-context-0.0.2.jar"
	jarPath := []string{
		"WEB-INF/lib/nacos-spring-context-0.0.2.jar",
		"WEB-INF/lib/BondeReverseDesensitize-0.0.25-SNAPSHOT.jar",
		"WEB-INF/lib/nacos-spring-context-0.0.2-20230718-SNAPSHOT.jar",
		"WEB-INF/lib/x-spring-context1-0.0.2-20230718-SNAPSHOT.jar",
	}

	//jarname := strings.Split(jarpath, "/")
	//jarname := strings.SplitAfter(jarpath, "/")
	//log.Println(jarname)

	rJar := `.*/(?P<jarname>.*)`
	rName := `^(?P<name>.*?)-\d`
	rVer := `\b(?:\d+\.?)+\-?\d+`

	rjar := regexp.MustCompile(rJar)
	rname := regexp.MustCompile(rName)
	rver := regexp.MustCompile(rVer)

	for _, jarpath := range jarPath {
		rjarv := rjar.FindStringSubmatch(jarpath)
		rnamev := rname.FindStringSubmatch(rjarv[1])
		rverv := rver.FindStringSubmatch(rjarv[1])
		log.Println(rnamev[1], rverv)
	}

}
