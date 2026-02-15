package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	// cnt 计算可以放多少花
	l, cnt := len(flowerbed), 0
	// 两次遍历，把旁边是1的0全变成2
	for i := 1; i < l; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 1 {
			flowerbed[i] = 2
		}
	}
	for i := 0; i < l-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i+1] == 1 {
			flowerbed[i] = 2
		}
	}
	// 左右指针，遍历一次即可
	// 本质上是计算处理过的数组的连续0
	p, q := 0, 0
	// 模拟
	for q < l {
		if flowerbed[p] != 0 && flowerbed[q] != 0 {
			p++
			q++
		} else if flowerbed[p] == 0 && flowerbed[q] == 0 {
			q++
		} else if flowerbed[p] == 0 && flowerbed[q] != 0 {
			if (q-p)%2 == 0 {
				// 有偶数个连续0，那就是num/2朵花
				cnt += (q - p) / 2
			} else {
				// 有奇数个连续0，那就是num/2+1朵花
				cnt += (q-p)/2 + 1
			}
			p = q
		} else {
			p = q
		}
	}
	// 注意滑动窗口最后一步也要判断
	if (q-p)%2 == 0 {
		cnt += (q - p) / 2
	} else {
		cnt += (q-p)/2 + 1
	}
	return cnt >= n
}

func main() {
	//fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 1}, 5))
	//fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	//fmt.Println(canPlaceFlowers([]int{0}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0, 0}, 3))
	fmt.Println(canPlaceFlowers([]int{0}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 0}, 2))
	//fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1}, 3))
}
