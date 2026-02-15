package main

import (
	"fmt"
	"strings"
)

func checkRecord(s string) bool {
	return strings.Count(s, "A") < 2 && !strings.Contains(s, "LLL")
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
}
