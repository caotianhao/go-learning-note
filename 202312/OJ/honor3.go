package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getTotalMoney(useInfo string) int {
	tmp := strings.Split(useInfo, "|")
	teleSecond, _ := strconv.Atoi(tmp[1])
	telePerMinute, _ := strconv.Atoi(tmp[2])
	dataInKB, _ := strconv.Atoi(tmp[3])
	dataPer10KB, _ := strconv.Atoi(tmp[4])

	tele, data := 0, 0

	teleMinute := int(math.Ceil(float64(teleSecond) / 60.0))
	tele25 := 6.0 * telePerMinute / 10.0
	tele610 := tele25 + telePerMinute/2.0
	tele1120 := tele610 + telePerMinute/2.0
	if teleMinute <= 1 {
		tele = 0
	} else if teleMinute <= 5 {
		tele = (teleMinute - 1) * 3.0 * telePerMinute / 20.0
	} else if teleMinute <= 10 {
		tele = (teleMinute-5)*telePerMinute/10.0 + tele25
	} else if teleMinute <= 20 {
		tele = (teleMinute-10)*telePerMinute/20.0 + tele610
	} else {
		tele = (teleMinute-20)*telePerMinute/5.0/10.0 + tele1120
	}

	dataInMB := int(math.Ceil(float64(dataInKB) / 1024.0))
	data100 := 1024.0 * dataPer10KB
	data200 := 2.0 * data100
	if dataInMB <= 100 {
		data = 1024.0 * dataInMB * dataPer10KB / 10.0 / 10.0
	} else if dataInMB <= 125 {
		data = data100
	} else if dataInMB <= 225 {
		data = 1024.0*(dataInMB-125)/10.0*dataPer10KB/10.0 + data100
	} else if dataInMB <= 325 {
		data = data200
	} else if dataInMB <= 725 {
		data = 1024.0*(dataInMB-325)/10.0*dataPer10KB/10.0 + data200
	} else {
		ps := int(math.Floor(float64(dataInMB-625) / 10.0 / 10.0))
		if ps%2 == 0 {
			data = (6 + ps/2) * data100
		} else {
			data = 1024.0*(dataInMB-(625+(ps+1)*100.0))/10.0*dataPer10KB/10.0 + (6+(ps-1)/2)*data100
		}
	}

	return tele + data
}

func main() {
	var useInfo string
	_, _ = fmt.Scanf("%s", &useInfo)
	fmt.Println(getTotalMoney(useInfo))
}
