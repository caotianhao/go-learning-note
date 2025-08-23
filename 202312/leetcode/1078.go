package main

import "fmt"

func findOcurrences(text string, first string, second string) (slice1078 []string) {
	tmp := divide1078(text)
	l := len(tmp)
	if l < 3 {
		return
	}
	for i := 2; i < l; i++ {
		if tmp[i-2] == first && tmp[i-1] == second {
			slice1078 = append(slice1078, tmp[i])
		}
	}
	return
}

func divide1078(text string) (slice1078 []string) {
	str := ""
	for _, v := range text {
		if v == ' ' {
			slice1078 = append(slice1078, str)
			str = ""
			continue
		}
		str += string(v)
	}
	slice1078 = append(slice1078, str)
	return
}

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
	fmt.Println(findOcurrences("we will", "we", "will"))
	//fmt.Println(divide1078("alice is a good girl she is a good student"))
}
