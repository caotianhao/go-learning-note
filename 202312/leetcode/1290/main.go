package main

import "fmt"

type ListNode1290 struct {
	Val  int
	Next *ListNode1290
}

func getDecimalValue(head *ListNode1290) int {
	//二进制链表转整数
	//初始化结果
	ans := 0
	//只要 head 的指针域非空，就继续进行循环
	for head != nil {
		//假设当前就是最低位，如果还有下一位的话就乘 2
		ans = 2*ans + head.Val
		//指针后移一位
		head = head.Next
	}
	return ans
}

func main() {
	fmt.Println(getDecimalValue(nil))
}
