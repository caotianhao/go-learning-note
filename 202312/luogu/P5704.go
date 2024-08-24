package main

import "fmt"

func main() {
	var ch byte
	_, _ = fmt.Scanf("%c", &ch)
	fmt.Printf("%c", ch-32)
}
