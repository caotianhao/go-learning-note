package main

import "fmt"

const c5723 = 1160

var prime5723 []int

func init() {
	prime5723 = append(prime5723, 2)
	for i := 3; i < c5723; i += 2 {
		if isPrime5723(i) {
			prime5723 = append(prime5723, i)
		}
	}
}

func isPrime5723(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(prime5723)
	k := 0
	_, _ = fmt.Scan(&k)
	sum, i, cnt, a := 0, 0, 0, make([]int, 0)
	for {
		if sum > k {
			break
		}
		sum += prime5723[i]
		a = append(a, prime5723[i])
		i++
		cnt++
	}
	a = a[:len(a)-1]
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println(cnt - 1)
}
