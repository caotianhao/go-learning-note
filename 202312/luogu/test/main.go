package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	s := make([]string, 0)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		s = append(s, scanner.Text())
	}
	fmt.Println(s)
}
