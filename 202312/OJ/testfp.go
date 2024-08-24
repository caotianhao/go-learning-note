package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var i = 0
	var mu sync.Mutex
	timer := time.NewTicker(time.Millisecond)

	//done := make(chan struct{})
	ctx, c := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-timer.C:
				mu.Lock()
				i++
				mu.Unlock()
			//case <-done:
			case <-ctx.Done():
				fmt.Println("任务结束")
				return
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	timer.Stop()

	//done <- struct{}{}
	c()
	mu.Lock()
	fmt.Println(i)
	mu.Unlock()

	fmt.Println("主程序结束")
}
