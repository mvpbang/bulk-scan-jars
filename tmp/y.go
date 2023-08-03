package main

import (
	"log"
	"strings"
)

func main() {
	str1 := "1.2.3"
	//str2 := "2.0.0.0"
	//str2 := "1.2.3.1"
	str2 := "1.2.3-20230802"

	log.Println(str2 >= str1)

	str3 := "BondeGate:2.1.28"
	log.Println(strings.Split(str3, ":"))
}
