package main

import (
	"sync"
	"sync/atomic"
)

const loop = 10000

func sync1() int {
	a := 1
	mu := &sync.Mutex{}
	wg := sync.WaitGroup{}
	for range loop {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			a++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return a
}

func sync2() int {
	a := 1
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range loop {
			a++
		}
	}()
	wg.Wait()
	return a
}

func sync3() int {
	var a int32 = 1
	wg := sync.WaitGroup{}
	for range loop {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&a, 1)
		}()
	}
	wg.Wait()
	return int(a)
}

// loop 太大电脑就卡死了，死死的，慎用！！！
//func sync4() int {
//	a := 1
//	mu := &sync.Mutex{}
//	done := make(chan struct{})
//	for i := 0; i < loop; i++ {
//		go func() {
//			mu.Lock()
//			a++
//			mu.Unlock()
//			done <- struct{}{}
//		}()
//	}
//	for i := 0; i < loop; i++ {
//		<-done
//	}
//	return a
//}

func sync5() int {
	a := 1
	mu := &sync.RWMutex{}
	done := make(chan struct{})
	for range loop {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			a++
			done <- struct{}{}
		}()
	}
	for range loop {
		<-done
	}
	return a
}

func sync6() int {
	a := 1
	for range loop {
		a++
	}
	return a
}

func main() {}
