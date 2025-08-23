package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

func main() {
	fmt.Println(defangIPaddr("255.100.50.0"))
}
