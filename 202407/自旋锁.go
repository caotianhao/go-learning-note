package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// SpinLock 自旋锁
type SpinLock struct {
	locked uint32
}

// Lock 尝试获取锁，如果锁已经被占用，则忙等待直到锁可用
func (s *SpinLock) Lock() {
	retries := 0
	for !atomic.CompareAndSwapUint32(&s.locked, 0, 1) {

		// 最开始特别多goroutine争抢，导致全是retrying，采用循环10次输出一下
		// fmt.Println("SpinLock: Lock failed, retrying...")

		if retries < 10 { // 10次重试，输出信息
			fmt.Println("SpinLock: Lock failed, retrying...")
			retries++
		}
		runtime.Gosched()
	}
	fmt.Println("SpinLock: Lock acquired successfully.")
}

func (s *SpinLock) Unlock() {
	atomic.StoreUint32(&s.locked, 0)
	fmt.Println("SpinLock: Unlock done.")
}

func main() {

	spinLock := &SpinLock{}

	// 使用多个goroutine尝试同时获取锁
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			spinLock.Lock()
			defer spinLock.Unlock()

			fmt.Printf("Goroutine %d is starting...\n", i)
			// 模拟执行逻辑相关
			time.Sleep(100 * time.Millisecond)

			fmt.Printf("Goroutine %d is done.\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines have finished.")
}
