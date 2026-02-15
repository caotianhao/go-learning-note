package main

import "fmt"

func main() {
	var m, h float64
	_, _ = fmt.Scanf("%f %f", &m, &h)
	bmi := m / h / h
	if bmi < 18.5 {
		fmt.Printf("Underweight")
	} else if bmi >= 18.5 && bmi < 24 {
		fmt.Printf("Normal")
	} else {
		fmt.Printf("%.6g\nOverweight", bmi)
	}
}
