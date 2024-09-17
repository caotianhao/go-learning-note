package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomLetter() byte {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.NewSource(time.Now().UnixNano())
	return letters[rand.Intn(52)]
}

func main() {
	target := "HelloWorld"              // 目标字符串
	result := make([]byte, len(target)) // 用于存放生成结果的切片
	currentIndex := 0                   // 当前匹配的字母索引

	for currentIndex < len(target) {
		letter := randomLetter()

		// 直接将生成的字母放入结果中
		result[currentIndex] = letter
		fmt.Printf("%s\n", string(result[:currentIndex+1]))

		// 检查是否和目标字母匹配
		if letter == target[currentIndex] {
			currentIndex++ // 匹配成功，继续下一个字母
		}
	}

	//fmt.Println("成功生成字符串:", target)
}
