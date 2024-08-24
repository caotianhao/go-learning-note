package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for i := 0; i < 2; i++ {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		t := s.Text()
		o := strings.Fields(t)
		fmt.Println(o)
	}
	//663 66 9 3
	//6613 66 9 3
}
