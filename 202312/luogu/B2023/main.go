package main

import "fmt"

func main() {
	var (
		ch  byte
		i   int
		f32 float32
		f64 float64
	)
	_, _ = fmt.Scanf("%c\n%d\n%f\n%f", &ch, &i, &f32, &f64)
	fmt.Printf("%c %d %.6f %.6f", ch, i, f32, f64)
}
