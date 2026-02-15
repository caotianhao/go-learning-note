package main

import (
	"fmt"
	"reflect"
)

func backspaceCompare(s string, t string) bool {
	return reflect.DeepEqual(string2Arr844(s), string2Arr844(t))
}

func string2Arr844(s string) string {
	str := make([]rune, 0)
	for _, v := range s {
		if v != '#' {
			str = append(str, v)
		} else {
			if len(str) != 0 {
				str = str[:len(str)-1]
			}
		}
	}
	return string(str)
}

func main() {
	fmt.Println(backspaceCompare("ab#dc", "ad#c"))
	fmt.Println(backspaceCompare("####ab#dc", "ad#dc"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a#b##cd", "####c##d#c"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
	fmt.Println(backspaceCompare("xywrrmp", "xywrrmu#p"))
	fmt.Println(string2Arr844("ab#dc"))
	fmt.Println(string2Arr844("a#dc"))
	fmt.Println(string2Arr844("##dc"))
	fmt.Println(string2Arr844("####ab#dc"))
}
