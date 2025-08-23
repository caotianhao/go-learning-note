package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	return len(strings.Fields(s))
}

func main() {
	fmt.Println(countSegments("  Hello,   my name is John  "))
}
