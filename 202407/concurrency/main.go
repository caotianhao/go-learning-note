package main

import (
	"fmt"
)

type foo struct {
	v1, v2 int64
}

func main() {
	var f foo
	other := 0
	f12, f34, f00 := 0, 0, 0
	f14, f32 := 0, 0

	for i := 0; i < 1<<20; i++ {
		go func() {
			a := foo{1, 2}
			f = a
		}()

		go func() {
			b := foo{3, 4}
			f = b
		}()

		//time.Sleep(50 * time.Millisecond)

		if f.v1 == 1 && f.v2 == 2 {
			f12++
		} else if f.v1 == 3 && f.v2 == 4 {
			f34++
		} else if f.v1 == 0 && f.v2 == 0 {
			f00++
		} else if f.v1 == 1 && f.v2 == 4 {
			f14++
		} else if f.v1 == 3 && f.v2 == 2 {
			f32++
		} else {
			fmt.Println(f)
			other++
		}
	}

	//time.Sleep(time.Millisecond)

	fmt.Printf("f00 = %d\n", f00)
	fmt.Printf("f12 = %d\n", f12)
	fmt.Printf("f34 = %d\n", f34)
	fmt.Printf("f14 = %d\n", f14)
	fmt.Printf("f32 = %d\n", f32)
	fmt.Printf("other = %d\n", other)
}
