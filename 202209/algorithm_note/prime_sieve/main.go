package main

import (
	"fmt"
	"time"
)

//埃拉托色尼素数筛法
//从 2 开始，划掉他的倍数，然后从头找没被划掉的，再划掉其倍数
func eratosthenes(n int) (slicePrime []int) {
	isNotPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !isNotPrime[i] {
			slicePrime = append(slicePrime, i)
			for j := 2 * i; j <= n; j += i {
				isNotPrime[j] = true
			}
		}
	}
	return
}

//正常遍历
func findPrime(n int) (slicePrime []int) {
	slicePrime = append(slicePrime, 2)
	for i := 3; i <= n; i += 2 {
		if isPrimeSSS(i) {
			slicePrime = append(slicePrime, i)
		}
	}
	return
}

func isPrimeSSS(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var k int
	fmt.Print("Prime from 2 to ")
	_, _ = fmt.Scanln(&k)

	time0 := time.Now()
	_ = findPrime(k)
	time1 := time.Since(time0)
	fmt.Println("normal:", time1)
	//fmt.Println(a)

	time2 := time.Now()
	_ = eratosthenes(k)
	time3 := time.Since(time2)
	fmt.Println("eratosthenes:", time3)
	//fmt.Println(b)
}
