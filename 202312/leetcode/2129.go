package main

import (
	"fmt"
	"strings"
)

func capitalizeTitle(title string) string {
	div := strings.Fields(title)
	res := make([]string, 0)
	for _, v := range div {
		if len(v) < 3 {
			res = append(res, strings.ToLower(v))
		} else {
			res = append(res, strings.ToUpper(string(v[0]))+strings.ToLower(v[1:]))
		}
	}
	return strings.Join(res, " ")
}

func main() {
	fmt.Println(capitalizeTitle("First leTTeR of EACH Word"))
}
