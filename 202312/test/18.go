package main

import (
	"fmt"
	"sync"
)

var j int

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < 1234; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			j++
		}()
	}
	wg.Wait()
	fmt.Println(j)
}
