package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var state int32 = 0
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				//state++                    //不行
				atomic.AddInt32(&state, 1) //行
			}
		}()
	}
	wg.Wait()
	fmt.Println(state)
}
