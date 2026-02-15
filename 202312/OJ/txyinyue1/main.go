package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param n int整型
 * @return int整型
 */
func fun(n int) (cnt int) {
	for n != 0 {
		t := n % 10
		n /= 10
		if t == 5 || t == 0 {
			return
		}
		cnt++
	}
	return -1
}

func main() {
	fmt.Println(fun(52))
}
