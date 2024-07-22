package main

import (
	"context"
	"fmt"
	"time"
)

func c1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "k1", "v1")
	return child
}

func c2() {
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("ctx.err =", ctx.Err())
	}
}

func c3() {
	ctx, cancel := context.WithCancel(context.TODO())
	t0 := time.Now()
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	select {
	case <-ctx.Done():
		t1 := time.Now()
		fmt.Println(t1.Sub(t0).Milliseconds())
	}
}

func main() {
	grandpa := context.Background()
	father := c1(grandpa)
	fmt.Println("father", father)
	fmt.Println("father.v", father.Value("k"))
	c2()
	c3()
}
