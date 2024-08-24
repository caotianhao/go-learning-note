package main

import (
	"fmt"
	"math"
)

func main() {
	var s, v int
	_, _ = fmt.Scanf("%d %d", &s, &v)
	minute := int(math.Ceil(float64(s)/float64(v))) + 10
	if minute <= 480 {
		fmt.Printf("%02d:%02d", (480-minute)/60, (480-minute)%60)
	} else {
		fmt.Printf("%02d:%02d", (1920-minute)/60, (1920-minute)%60)
	}
}
