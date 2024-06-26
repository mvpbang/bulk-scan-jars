package main

import (
	"log"
	"strings"
)

var eqs = []string{"xxl:1.4.0", "x-spring-context1:0.0.2-20230718"}

func main() {
	str1 := "xxl"
	str2 := "-1.7.8"
	for _, v := range eqs {
		if str1 == strings.Split(v, ":")[0] {
			log.Println(v, str1)
		}
	}

	log.Println(strings.TrimLeft(str2, "-"))

}
