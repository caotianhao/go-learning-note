package main

import (
	"fmt"
	"math"
)

// query       0  1  2  3  4     print
// ---------------------------------------
// Q 0 4       5  5  5  5  5  -> 5
// B 0 2 3     2  2  5  5  5     ok
// B 3 4 2     2  2  5  3  5     ok
// Q 2 3       2  2  5  3  5     5
// B 1 4 3     2  2  5  3  5     fail
// B 1 4 2     2  0  3  1  5     ok
// Q 0 4       2  0  3  1  5     0
// R 0 3 1     2  0  3  1  5     fail
// R 1 4 5     2  0  3  1  5     fail
// Q 2 4       2  0  3  1  5     1

// 过了 90% 剩下的没找到问题
// 同样的代码，有时候 90% 有时候 85%
func main() {
	// 站点数，座位数，查询数
	station, seat, query := 0, 0, 0
	// 查询标志
	var char byte
	// 出发地，目的地，涉及的票数
	src, dst, numOfTicket := 0, 0, 0

	_, _ = fmt.Scanf("%d %d %d", &station, &seat, &query)

	// 记录每个站的票数
	sliceOfSeat := make([]int, station)
	// 记录从 a 到 b 卖出了几张票
	sold := map[[2]int]int{}
	// 初始化全票都在
	for i := 0; i < station; i++ {
		sliceOfSeat[i] = seat
	}
	for i := 0; i < query; i++ {
		// 读取查询标志
		_, _ = fmt.Scanf("%c", &char)
		if char == 'Q' {
			// Q 查询所经过的地方的最少票
			_, _ = fmt.Scanf("%d %d", &src, &dst)
			minN := math.MaxInt64
			for index := src; index < dst; index++ {
				if sliceOfSeat[index] < minN {
					minN = sliceOfSeat[index]
				}
			}
			fmt.Printf("%d\n", minN)
		} else if char == 'B' {
			// B 查询所经过的地方的最少票
			_, _ = fmt.Scanf("%d %d %d", &src, &dst, &numOfTicket)
			minN := math.MaxInt64
			for index := src; index < dst; index++ {
				if sliceOfSeat[index] < minN {
					minN = sliceOfSeat[index]
				}
			}
			// 若余票最小值大于等于所要购买的，卖出
			if minN >= numOfTicket {
				for index := src; index < dst; index++ {
					sliceOfSeat[index] -= numOfTicket
				}
				fmt.Printf("OK!\n")
				// 同时记录卖出的票
				sold[[2]int{src, dst}] += numOfTicket
			} else {
				fmt.Printf("Fail!\n")
			}
		} else {
			_, _ = fmt.Scanf("%d %d %d", &src, &dst, &numOfTicket)
			if soldOut, ok := sold[[2]int{src, dst}]; !ok {
				// 若没卖出过从 a 到 b 的票直接 fail
				fmt.Printf("Fail!\n")
			} else {
				// 若卖出过，退的比卖的还多，也 fail
				if soldOut < numOfTicket {
					fmt.Printf("Fail!\n")
				} else {
					// 退回
					fmt.Printf("OK!\n")
					for index := src; index < dst; index++ {
						sliceOfSeat[index] += numOfTicket
					}
				}
			}
		}
	}
}
