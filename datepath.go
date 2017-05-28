package main

import (
	"log"
	"strconv"
	"time"
)

func intToStr(a int) string {
	if a < 10 {
		return "0" + strconv.Itoa(a)
	}
	return strconv.Itoa(a)
}

func timepath(y int, m int, d int) string {
	return intToStr(y) + intToStr(m) + "/" + intToStr(d)
}

func GenerateDateList(d int) []string {
	var datalist []string
	starttime := time.Now().Add(time.Hour * time.Duration(d) * -24)
	for {
		if starttime.After(time.Now()) {
			break
		}
		datalist = append(datalist, timepath(starttime.Year(), int(starttime.Month()), starttime.Day()))
		starttime = starttime.Add(time.Hour * 24)
	}
	return datalist
}

func Preproducor(d int, n *NSQD, config *Config) {
	for _, datepath := range GenerateDateList(d) {
		path := config.PrefixPath + datepath
		log.Printf("make " + path)
		select {
		case n.PreChan <- path:
			break
		case <-n.ExitChan:
			break
		}
	}
}
