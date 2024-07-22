package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		state int32 = 0
		wg    sync.WaitGroup
	)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				state++
			}
		}()
	}
	wg.Wait()
	fmt.Println(state)
}
