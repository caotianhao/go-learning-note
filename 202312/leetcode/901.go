package main

import "fmt"

type StockSpanner struct {
	stock []int
}

func Constructor901() StockSpanner {
	return StockSpanner{stock: []int{}}
}

func (ss *StockSpanner) Next(price int) int {
	cnt := 0
	ss.stock = append(ss.stock, price)
	for i := len(ss.stock) - 1; i >= 0; i-- {
		if ss.stock[i] > price {
			break
		} else {
			cnt++
		}
	}
	return cnt
}

func main() {
	//["StockSpanner","next","next","next","next","next","next","next"],
	//[      [],      [100],  [80],  [60],   [70],  [60],  [75],  [85]]
	//	输出：[null,1,1,1,2,1,4,6]
	s := Constructor901()
	fmt.Println(s.Next(100))
	fmt.Println(s.Next(80))
	fmt.Println(s.Next(60))
	fmt.Println(s.Next(70))
	fmt.Println(s.Next(60))
	fmt.Println(s.Next(75))
	fmt.Println(s.Next(85))
}
