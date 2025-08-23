package main

import (
	"fmt"
	"math/bits"
)

//给一个 uint32 型整数，统计它的二进制里有几个 1

//法1：用 1,10,100,1000,10000... 二进制数去和 num 做 & 运算
//如果结果是 0，就说明 1 对应的那位是 0，反之亦然，这样下去统计 32 位后就是结果
//func hammingWeight(num uint32) int {
//	one := 0
//	for i := 0; i < 32; i++ {
//		if (1<<i)&num != 0 {
//			one++
//		}
//	}
//	return one
//}

//法2：直接转成 int 再转二进制 slice 然后统计 1 的个数
//func hammingWeight(num uint32) int {
//	one := 0
//	for _, v := range dec2Bin191(int(num)) {
//		if v == 1 {
//			one++
//		}
//	}
//	return one
//}
//
//func dec2Bin191(num int) (slice191 []int) {
//	if num == 0 {
//		return
//	}
//	for num != 0 {
//		slice191 = append(slice191, num%2)
//		num /= 2
//	}
//	return
//}

//法3：调库...学到了学到了
func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}

//两种方法在 leetcode 上效率差不多.....
func main() {
	fmt.Println(hammingWeight(3)) //0011
	fmt.Println(hammingWeight(4294967000))
}
