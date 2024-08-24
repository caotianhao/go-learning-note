package main

import "fmt"

const (
	queue      int = 1e5 + 1
	initNumber int = 1
)

var initQueue [queue]int

func init() {
	initQueue[0] = initNumber
	for i := 1; i < queue; i++ {
		initQueue[i] = initQueue[i-1] + i
	}
}

func main() {
	length, maxValue := 0, 0
	_, _ = fmt.Scanf("%d %d", &length, &maxValue)
	pq := initQueue[:length]
	cnt := 0
	if pq[length-1] > maxValue {
		for i := 0; i < length; i++ {
			if pq[i] > maxValue {
				cnt++
			}
		}
	}
	for i := 0; i < length; i++ {

	}
	fmt.Println(initQueue[:length])
}
