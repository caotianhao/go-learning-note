package main

import (
	"fmt"
	"strconv"
)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var sf, ss, st int64
	for firstWord[0] == 'a' {
		if len(firstWord) != 1 {
			firstWord = firstWord[1:]
		} else {
			sf = 0
			break
		}
	}
	for secondWord[0] == 'a' {
		if len(secondWord) != 1 {
			secondWord = secondWord[1:]
		} else {
			ss = 0
			break
		}
	}
	for targetWord[0] == 'a' {
		if len(targetWord) != 1 {
			targetWord = targetWord[1:]
		} else {
			st = 0
			break
		}
	}
	lf, ls, lt := len(firstWord), len(secondWord), len(targetWord)
	//f, s, t := []rune(firstWord), []rune(secondWord), []rune(targetWord)
	stringF, stringS, stringT := "", "", ""
	//fmt.Println("aaaa+", stringF)
	for i := 0; i < lf; i++ {
		//temp, _ = strconv.ParseInt(string(firstWord[i]), 10, 64)
		//temp -= 97
		stringF += fmt.Sprintf("%d", firstWord[i]-97)
	}
	//fmt.Println("aaaaa+", stringF)
	for i := 0; i < ls; i++ {
		stringS += fmt.Sprintf("%d", secondWord[i]-97)
	}
	for i := 0; i < lt; i++ {
		stringT += fmt.Sprintf("%d", targetWord[i]-97)
	}
	sf, _ = strconv.ParseInt(stringF, 10, 64)
	ss, _ = strconv.ParseInt(stringS, 10, 64)
	st, _ = strconv.ParseInt(stringT, 10, 64)
	//fmt.Println("ddd", sf, ss, st)
	if sf+ss == st {
		return true
	} else {
		return false
	}
}

func main() {
	firstWord := "bcd"
	secondWord := "iiba"
	targetWord := "aaa"
	fmt.Println(isSumEqual(firstWord, secondWord, targetWord))
}
