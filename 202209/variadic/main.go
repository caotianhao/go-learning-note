package main

import "fmt"

//支持可变参数，所以Go不支持函数的重载
//实现求0-n个数的和
func sum(args ...int) (sum int) {
	sum = 0
	for i := range args {
		sum += args[i]
	}
	return
}

//实现求1-n个数的和
//如果形参列表有可变参数，则其应在形参列表最后
func sum1(n1 float64, args ...float64) (sum float64) {
	sum = n1
	for i := range args {
		sum += args[i]
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, -3, 4, 5, -6, 7, 8))
	fmt.Println(sum1(1.1, 2.5, 3.47, 4.57, -5.8587, 6.3334656, 7, 8.1, 9.33222))
}
