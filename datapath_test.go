package main

import (
	"testing"
)

func TestPrintDataPath(t *testing.T) {
	dateList := GenerateDateList(10)
	for _, date := range dateList {
		t.Log(date)
	}
}
