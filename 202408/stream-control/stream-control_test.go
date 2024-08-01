package main

import (
	"testing"
	"time"
)

func TestTimeoutWithTime(t *testing.T) {
	TimeoutWithTime(500 * time.Millisecond)
}

func TestTimeoutWithContext(t *testing.T) {
	TimeoutWithContext(500 * time.Millisecond)
}
