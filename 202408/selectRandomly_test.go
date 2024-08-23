package main

import (
	"math/rand"
	"testing"
	"time"
)

var source = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestSelectRandomly(t *testing.T) {
	cc := source.Intn(1e6) + 9999

	ch := make(chan int)
	selectRandomly(ch)
	var c1, c2, c3 int
	for i := 0; i < cc; i++ {
		switch <-ch {
		case 1:
			c1++
		case 2:
			c2++
		case 3:
			c3++
		}
	}
	if cc != c1+c2+c3 {
		t.Errorf("%d don't match", cc-c1-c2-c3)
	}
}
