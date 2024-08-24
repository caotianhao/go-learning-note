package main

import (
	"fmt"
	"strings"
)

func change(s string) string {
	k := strings.Count(s, "0")
	if k < 3 {
		return strings.Replace(s, "0", "00", -1)
	} else {
		return strings.Replace(s, "0", "00", 3)
	}
}

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)
	fmt.Printf("%s", change(s))
}
