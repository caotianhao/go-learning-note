package main

import (
	cal "gogogo/202408/gotraining/internal/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	c := cal.Calculator{}
	res := c.Add(1, 2)
	if res != 3 {
		t.Fatalf("add err!err!err!")
	}
	t.Logf("yes")
}

func TestSubtract(t *testing.T) {
	c := cal.Calculator{}
	res := c.Subtract(3, 2)
	if res != 1 {
		t.Fatalf("sub err!err!err!")
	}
	t.Logf("yes")
}
