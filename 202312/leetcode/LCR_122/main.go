package main

import (
	"fmt"
	"strings"
)

func pathEncryption(path string) string {
	return strings.Replace(path, ".", " ", -1)
}

func main() {
	fmt.Println(pathEncryption("a.aef.qerf.bb"))
}
