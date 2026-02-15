package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// SimpleIDGenerator 是一个简单的ID生成器实现
type SimpleIDGenerator struct {
	counter uint64
}

// GetID 生成并返回一个唯一的ID
// 注意：这个实现简单使用了时间戳和原子计数器来确保ID的唯一性，但跨程序实例时不保证全局唯一
func (idg *SimpleIDGenerator) GetID() string {
	// 获取当前时间戳（纳秒）并转换为毫秒
	ts := time.Now().UnixNano()

	// 使用原子操作来增加计数器
	count := atomic.AddUint64(&idg.counter, 1)

	// 拼接时间戳和计数器为ID
	return fmt.Sprintf("%d%02d", ts, count)
}

func main() {
	gen := SimpleIDGenerator{}

	// 模拟并发调用
	for i := 0; i < 10; i++ {
		go func() {
			id := gen.GetID()
			fmt.Println(id)
		}()
	}

	// 等待足够的时间以确保所有goroutine都执行完毕
	time.Sleep(1 * time.Second)
}
