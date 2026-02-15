package main

import (
	"errors"
	"fmt"
	"sync"
)

type Promise[T any] struct {
	value T
	err   error
	done  bool
	mu    sync.Mutex
	cond  *sync.Cond
}

// NewPromise 创建一个新的 Promise
func NewPromise[T any]() *Promise[T] {
	p := &Promise[T]{}
	p.cond = sync.NewCond(&p.mu)
	return p
}

// Get 获取值和错误，阻塞直到 promise 完成
func (p *Promise[T]) Get() (T, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for !p.done {
		p.cond.Wait()
	}
	return p.value, p.err
}

// Value 获取值，阻塞直到 promise 完成
func (p *Promise[T]) Value() T {
	value, _ := p.Get()
	return value
}

// Error 获取错误，阻塞直到 promise 完成
func (p *Promise[T]) Error() error {
	_, err := p.Get()
	return err
}

// Resolve 已解决（设置值）
func (p *Promise[T]) Resolve(value T) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.done {
		return errors.New("promise already completed")
	}
	p.value = value
	p.done = true
	p.cond.Broadcast()
	return nil
}

// Reject 已拒绝（设置错误）
func (p *Promise[T]) Reject(err error) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.done {
		return errors.New("promise already completed")
	}
	p.err = err
	p.done = true
	p.cond.Broadcast()
	return nil
}

// Then 当 promise 完成时执行的函数
func (p *Promise[T]) Then(f func(value T, err error)) {
	go func() {
		value, err := p.Get()
		f(value, err)
	}()
}

func main() {
	// 示例使用
	p := NewPromise[int]()

	// 启动一个 goroutine 来解决 promise
	go func() {
		// 模拟一些工作
		// time.Sleep(2 * time.Second)
		_ = p.Resolve(42)
	}()

	// 使用 Then 方法设置回调函数
	p.Then(func(value int, err error) {
		if err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			fmt.Println("Value:", value)
		}
	})

	// 等待 promise 完成
	_, _ = p.Get()
	fmt.Println(p.Get())
}
