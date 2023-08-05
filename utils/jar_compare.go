package utils

import (
	"log"
	"strings"
)

func comareEq(config Config, jarinfo JarInfo, ch chan int) {
	// 迭代比较
	for _, eqjar := range config.Eq {
		if jarinfo.Name == strings.Split(eqjar, ":")[0] {
			if jarinfo.Ver != strings.Split(eqjar, ":")[1] {
				log.Printf("相等判断: %v %v %v:%v 不等于 %v", jarinfo.ParPath, jarinfo.JarPath, jarinfo.Name, jarinfo.Ver, eqjar)
			}
		}
	}

	ch <- 1
}

func comareLe(config Config, jarinfo JarInfo, ch chan int) {
	for _, eqjar := range config.Le {
		if jarinfo.Name == strings.Split(eqjar, ":")[0] {
			//log.Println("+++++", jarinfo)
			if jarinfo.Ver < strings.Split(eqjar, ":")[1] {
				log.Printf("大于等于判断: %v %v %v:%v 不大于等于 %v", jarinfo.ParPath, jarinfo.JarPath, jarinfo.Name, jarinfo.Ver, eqjar)
			}
		}
	}
	ch <- 1
}

// 比对坐标

func CompareVer(config Config, jarinfo JarInfo) {
	ch := make(chan int, 0)
	//log.Println(jarinfo)
	go comareEq(config, jarinfo, ch)
	go comareLe(config, jarinfo, ch)
	<-ch
	<-ch
}
