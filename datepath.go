package main

import (
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
	starttime := time.Now()
	nowtime := starttime
	for {
		if !nowtime.Add(time.Hour * time.Duration(d) * -24).Before(starttime) {
			break
		}
		datalist = append(datalist, timepath(starttime.Year(), int(starttime.Month()), starttime.Day()))
		starttime = starttime.Add(time.Hour * -24)
	}
	return datalist
}
