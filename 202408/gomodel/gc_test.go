package main

import "testing"

const (
	gcSize = 1 << 24
	looper = 1<<6 - 1
)

func TestMakeLargeSlice(t *testing.T) {
	for i := 0; i < looper; i++ {
		slice := makeLargeSlice(gcSize)
		if len(slice) != gcSize {
			t.Errorf("Expected slice length to be %d, got %d", gcSize, len(slice))
		}
	}
}
