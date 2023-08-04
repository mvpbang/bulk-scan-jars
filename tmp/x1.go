package main

import (
	"log"
	"regexp"
)

func main() {
	slice_str1 := []string{
		"BondeReverseDesensitize-0.0.25-SNAPSHOT.jar",
		"nacos-spring-context-0.0.2.jar",
	}

	//r1 := `(?P<id>\d)\.(?P=id)`
	r1 := `(?P<id>\d+)`
	rc1 := regexp.MustCompile(r1)
	log.Println(rc1.String())
	//result := rc1.FindString([...]slice_str1)
	for _, v := range slice_str1 {
		result := rc1.FindAllString(v, -1)
		log.Println(result)
	}

}
