package main

import "fmt"

func isRunningString(s1, s2 string) string {
	l := len(s1)
	for i := 0; i < l; i++ {
		if s1[i:]+s1[:i] == s2 {
			return "Yes"
		}
	}
	return "No"
}

func main() {
	t := 0
	_, _ = fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		s1, s2 := " ", " "
		_, _ = fmt.Scanf("%s", &s1)
		_, _ = fmt.Scanf("%s", &s2)
		fmt.Printf("%s\n", isRunningString(s1, s2))
	}
}
