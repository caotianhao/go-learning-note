package main

import "fmt"

//func convertTemperature(celsius float64) (slice2469 []float64) {
//	slice2469 = append(slice2469, celsius+273.15, celsius*1.80+32.00)
//	return
//}

//func convertTemperature(celsius float64) []float64 {
//	ans := make([]float64, 2)
//	ans[0], ans[1] = celsius+273.15, celsius*1.8+32
//	return ans
//}

func convertTemperature(c float64) []float64 {
	return []float64{c + 273.15, c*1.8 + 32}
}

func main() {
	fmt.Println(convertTemperature(12.22))
}
