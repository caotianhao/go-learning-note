package main

import "fmt"

var globalPath []string

var myPhoneMap = map[string]string{
	"2": "abc", "3": "def",
	"4": "ghi", "5": "jkl", "6": "mno",
	"7": "pqrs", "8": "tuv", "9": "wxyz",
}

func phoneWithNumber(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var dfs func(string, string, int)
	dfs = func(digits, path string, idx int) {
		if len(digits) == idx {
			globalPath = append(globalPath, path)
		} else {
			d := string(digits[idx])
			letter := myPhoneMap[d]
			for i := 0; i < len(letter); i++ {
				dfs(digits, path+string(letter[i]), idx+1)
			}
		}
	}
	dfs(digits, "", 0)
	return globalPath
}

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)
	ans := phoneWithNumber(s)
	fmt.Print("[")
	for i := 0; i < len(ans)-1; i++ {
		fmt.Printf("\"%s\",", ans[i])
	}
	fmt.Printf("\"%s\"", ans[len(ans)-1])
	fmt.Print("]")
}
