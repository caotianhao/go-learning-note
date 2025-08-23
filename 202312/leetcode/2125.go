package main

import "fmt"

func numberOfBeams(bank []string) (res int) {
	beam, light := make([]int, len(bank)), make([]int, 0)
	for i, v := range bank {
		for _, s := range v {
			if s == '1' {
				beam[i]++
			}
		}
	}
	for _, v := range beam {
		if v != 0 {
			light = append(light, v)
		}
	}
	if len(light) < 2 {
		return 0
	}
	for i := 0; i < len(light)-1; i++ {
		res += light[i] * light[i+1]
	}
	return
}

func main() {
	fmt.Println(numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
}
