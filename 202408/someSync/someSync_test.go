package someSync

import (
	"testing"
	"time"
)

func TestSync1(t *testing.T) {
	begin := time.Now()
	s1 := sync1()
	end := time.Since(begin)
	if s1 != loop+1 {
		t.Errorf("sync1() returned %d, want %d", s1, loop+1)
	} else {
		t.Logf("sync1() = %d, time = %v", s1, end)
	}
}

func TestSync2(t *testing.T) {
	begin := time.Now()
	s2 := sync2()
	end := time.Since(begin)
	if s2 != loop+1 {
		t.Errorf("sync2() returned %d, want %d", s2, loop+1)
	} else {
		t.Logf("sync2() = %d, time = %v", s2, end)
	}
}

func TestSync3(t *testing.T) {
	begin := time.Now()
	s3 := sync3()
	end := time.Since(begin)
	if s3 != loop+1 {
		t.Errorf("sync3() returned %d, want %d", s3, loop+1)
	} else {
		t.Logf("sync3() = %d, time = %v", s3, end)
	}
}

func TestSync5(t *testing.T) {
	begin := time.Now()
	s5 := sync5()
	end := time.Since(begin)
	if s5 != loop+1 {
		t.Errorf("sync5() returned %d, want %d", s5, loop+1)
	} else {
		t.Logf("sync5() = %d, time = %v", s5, end)
	}
}

func TestSync6(t *testing.T) {
	begin := time.Now()
	s6 := sync6()
	end := time.Since(begin)
	if s6 != loop+1 {
		t.Errorf("sync6() returned %d, want %d", s6, loop+1)
	} else {
		t.Logf("sync6() = %d, time = %v", s6, end)
	}
}
