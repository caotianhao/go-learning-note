package main

import "fmt"

func toGoatLatin(sentence string) (slice824 string) {
	tmp := divide824(sentence)
	l := len(tmp)
	for i := 0; i < l; i++ {
		if tmp[i][0] == 'A' || tmp[i][0] == 'E' || tmp[i][0] == 'I' || tmp[i][0] == 'O' ||
			tmp[i][0] == 'U' || tmp[i][0] == 'a' || tmp[i][0] == 'e' || tmp[i][0] == 'i' ||
			tmp[i][0] == 'o' || tmp[i][0] == 'u' {
			tmp[i] += "ma"
		} else {
			tmp[i] = reverse824(tmp[i])
			tmp[i] += "ma"
		}
		for j := 1; j <= i+1; j++ {
			tmp[i] += "a"
		}
		slice824 += tmp[i]
		slice824 += " "
	}
	return slice824[:len(slice824)-1]
}

func divide824(sentence string) (slice824 []string) {
	str := ""
	for _, v := range sentence {
		if v == ' ' {
			slice824 = append(slice824, str)
			str = ""
			continue
		}
		str += string(v)
	}
	slice824 = append(slice824, str)
	return
}

func reverse824(str string) string {
	return str[1:] + string(str[0])
}

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(reverse824("s"))
}
