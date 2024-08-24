package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var it *Single

type Single struct {
	Info string
}

func getSingle() *Single {
	once.Do(func() {
		it = &Single{Info: "a"}
	})
	return it
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			i := getSingle()
			fmt.Println(i.Info)
		}()
	}
	wg.Wait()
}
