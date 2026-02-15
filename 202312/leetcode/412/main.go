package main

import "fmt"

//answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
//answer[i] == "Fizz" 如果 i 是 3 的倍数。
//answer[i] == "Buzz" 如果 i 是 5 的倍数。
//answer[i] == i （以字符串形式）如果上述条件全不满足。
func fizzBuzz(n int) []string {
	str := make([]string, 0)
	for i := 0; i < n; i++ {
		str = append(str, fmt.Sprintf("%d", i+1))
	}
	for i := 0; i <= n; i++ {
		if i%3 == 0 && i%5 != 0 {
			str[i-1] = "Fizz"
		} else if i%3 != 0 && i%5 == 0 {
			str[i-1] = "Buzz"
		} else if i%3 == 0 && i%5 == 0 && i != 0 {
			str[i-1] = "FizzBuzz"
		}
	}
	return str
}

func main() {
	fmt.Println(fizzBuzz(79))
}
