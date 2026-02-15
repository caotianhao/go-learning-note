package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var c1 = make(chan struct{}, 1)
var c2 = make(chan struct{}, 1)
var c3 = make(chan struct{}, 1)

func print1() {
	for i := 0; i < 10; i++ {
		<-c1
		fmt.Println("1111111111111111111111111111111")
		time.Sleep(time.Second)
		c2 <- struct{}{}
	}
}

func print2() {
	for i := 0; i < 10; i++ {
		<-c2
		fmt.Println("22222222222222222222222")
		time.Sleep(time.Second)
		c3 <- struct{}{}
	}
}

func print3() {
	for i := 0; i < 10; i++ {
		<-c3
		fmt.Println("333333333333")
		time.Sleep(time.Second)
		c1 <- struct{}{}
	}
}

func myPrint() {
	wg.Add(3)

	go func() {
		defer wg.Done()
		print1()
	}()

	go func() {
		defer wg.Done()
		print2()
	}()

	go func() {
		defer wg.Done()
		print3()
	}()

	c1 <- struct{}{}
	wg.Wait()

	fmt.Println()
	fmt.Println("ok")
}

func main() {
	myPrint()
}
