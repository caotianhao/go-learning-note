package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getPrice(initialPrice float64, couponInfos string) int {
	info1 := strings.Split(couponInfos, ";")
	allInfo := make([][]string, 0)
	for _, info := range info1 {
		i := strings.Split(info, ",")
		a := make([]string, 0)
		for _, v := range i {
			ii := strings.Split(v, ":")
			a = append(a, ii[1])
			//fmt.Println(i, ii)
		}
		allInfo = append(allInfo, a)
	}
	now := make([]float64, 0)
	for _, aInfo := range allInfo {
		initialPrice1 := initialPrice
		if aInfo[len(aInfo)-1] == "false" {
			if len(aInfo) == 4 {
				man, _ := strconv.Atoi(aInfo[1])
				jian, _ := strconv.Atoi(aInfo[2])
				if initialPrice1 >= float64(man) {
					initialPrice1 -= float64(jian)
				}
			} else {
				if aInfo[0] == "discountCoupon" {
					dis, _ := strconv.ParseFloat(aInfo[1], 64)
					initialPrice1 *= dis
				} else {
					jian, _ := strconv.Atoi(aInfo[1])
					initialPrice1 -= float64(jian)
				}
			}
			now = append(now, initialPrice1)
		}
	}
	sort.Float64s(now)
	if len(now) != 0 {
		initialPrice = now[0]
	}
	for _, aInfo := range allInfo {
		if aInfo[len(aInfo)-1] == "true" {
			if len(aInfo) == 4 {
				man, _ := strconv.Atoi(aInfo[1])
				jian, _ := strconv.Atoi(aInfo[2])
				if initialPrice >= float64(man) {
					initialPrice -= float64(jian)
				}
			} else {
				if aInfo[0] == "discountCoupon" {
					dis, _ := strconv.ParseFloat(aInfo[1], 64)
					initialPrice *= dis
				} else {
					jian, _ := strconv.Atoi(aInfo[1])
					initialPrice -= float64(jian)
				}
			}
		}
	}
	return int(initialPrice)
}

func main() {
	fmt.Println(getPrice(100, "type:fullReductionCoupon,basePrice:500,min"+
		"usPrice:50,isStackable:true;type:discountCoupon,discount:0.5,isStackable:false"))
}
