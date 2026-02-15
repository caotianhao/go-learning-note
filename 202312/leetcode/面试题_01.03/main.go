package main

import (
	"fmt"
	"strings"
)

func replaceSpaces(S string, length int) string {
	return strings.Replace(S[:length], " ", "%20", -1)
}

func main() {
	fmt.Println(replaceSpaces("Mr John Smith    ", 13))
}
