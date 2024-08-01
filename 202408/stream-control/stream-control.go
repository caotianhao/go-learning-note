package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// SlowQuery 模拟一个慢查询
func SlowQuery() error {
	rand.NewSource(time.Now().UnixNano())
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return nil
}

// TimeoutWithTime 使用time.After实现超时控制
func TimeoutWithTime(timeout time.Duration) {
	done := make(chan error, 1)
	start := time.Now()

	go func() {
		done <- SlowQuery()
	}()

	select {
	case err := <-done:
		if err == nil {
			fmt.Printf("Execution time: %v\n", time.Since(start))
		}
	case <-time.After(timeout):
		fmt.Println("timeout occur")
	}
}

// TimeoutWithContext 使用context.WithTimeout实现超时控制
func TimeoutWithContext(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	start := time.Now()

	go func() {
		done <- SlowQuery()
	}()

	select {
	case err := <-done:
		if err == nil {
			fmt.Printf("Execution time: %v\n", time.Since(start))
		}
	case <-ctx.Done():
		fmt.Println("timeout occur")
	}
}

func main() {
	fmt.Println("Using time.After:")
	TimeoutWithTime(500 * time.Millisecond)

	fmt.Println("Using context.WithTimeout:")
	TimeoutWithContext(500 * time.Millisecond)
}
