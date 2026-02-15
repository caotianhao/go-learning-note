package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	if n <= 150 {
		fmt.Printf("%.1f", 0.4463*float64(n))
	} else if n > 150 && n < 401 {
		fmt.Printf("%.1f", 0.4463*150+0.4663*float64(n-150))
	} else {
		fmt.Printf("%.1f", 0.4463*150+0.4663*250+0.5663*float64(n-400))
	}
}
