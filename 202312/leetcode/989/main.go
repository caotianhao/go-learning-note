package main

import "fmt"

func addToArrayForm(num []int, k int) []int {
	//if num
	kk, nn := divide989(k), reverse989(num)
	ln, lk := len(nn), len(kk)
	if ln > lk {
		for i := 0; i < ln-lk; i++ {
			kk = append(kk, 0)
		}
	} else if lk > ln {
		for i := 0; i < lk-ln; i++ {
			nn = append(nn, 0)
		}
	}
	l := len(nn)
	//fmt.Println(nn, kk)
	ret := make([]int, l+1)
	for i := 0; i < l; i++ {
		ret[i] += nn[i] + kk[i]
		if ret[i] > 9 {
			ret[i] -= 10
			ret[i+1]++
		}
	}
	if ret[l] == 0 {
		ret = ret[:l]
	}
	return reverse989(ret)
}

func divide989(n int) (sl []int) {
	for n != 0 {
		sl = append(sl, n%10)
		n /= 10
	}
	return
}

func reverse989(arr []int) []int {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}
	return arr
}

func main() {
	//fmt.Println(reverse989([]int{1, 2, 3, 4, 5}))
	//fmt.Println(divide989(7890))
	fmt.Println(addToArrayForm([]int{2, 7, 4}, 981))
	fmt.Println(addToArrayForm([]int{9, 9, 9}, 1))
	fmt.Println(addToArrayForm([]int{9, 9}, 1))
	fmt.Println(addToArrayForm([]int{0}, 1233))
}
