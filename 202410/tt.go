package main

import (
	"fmt"
	"math/rand"
	"time"
)

const song = 803

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	fmt.Println("平平无奇的苹")
	fmt.Println(r.Intn(song) + 1)
}
