package main

import "fmt"

func kyV8(atk, matk, def, mdef, speed, hp int) int {
	fmt.Println(atk, matk, def, mdef, speed, hp, atk+matk+def+mdef+speed+hp)
	return 0
}

func main() {
	v801 := kyV8(47, 0, 27, 27, 41, 82)
	v802 := kyV8(0, 47, 27, 27, 41, 82)
	v803 := kyV8(62, 0, 27, 27, 30, 79)
	v804 := kyV8(0, 62, 27, 27, 30, 79)
	v805 := kyV8(42, 0, 31, 31, 26, 115)
	v806 := kyV8(0, 42, 31, 31, 26, 115)
	v807 := kyV8(25, 0, 45, 45, 27, 90)
	v808 := kyV8(0, 25, 45, 45, 27, 90)
	_ = []int{v801, v802, v803, v804, v805, v806, v807, v808}
}
