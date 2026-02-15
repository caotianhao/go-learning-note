package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(1<<6 - 1 + http.StatusAccepted)
	fmt.Println(1<<20 | 1<<18 | 1)
	for i := 3; i < 16; i++ {
		fmt.Println(i, i^1)
	}
}
