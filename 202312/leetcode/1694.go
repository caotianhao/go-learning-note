package main

import "fmt"

func reformatNumber(number string) string {
	mySlice, a := make([]rune, 0), ""
	for _, v := range number {
		if v >= '0' && v <= '9' {
			mySlice = append(mySlice, v)
		}
	}
	if len(mySlice) <= 3 {
		return string(mySlice)
	} else if len(mySlice) == 4 {
		return string(mySlice[0]) + string(mySlice[1]) + "-" + string(mySlice[2]) + string(mySlice[3])
	} else if len(mySlice)%3 == 0 {
		for i := 0; i < len(mySlice)/3; i++ {
			a += string(mySlice[3*i])
			a += string(mySlice[3*i+1])
			a += string(mySlice[3*i+2])
			if i != len(mySlice)/3-1 {
				a += "-"
			}
		}
	} else if len(mySlice)%3 == 1 {
		for i := 0; i < len(mySlice)/3-1; i++ {
			a += string(mySlice[3*i])
			a += string(mySlice[3*i+1])
			a += string(mySlice[3*i+2])
			if i != len(mySlice)/3-2 {
				a += "-"
			} else {
				a += "-"
				a += string(mySlice[len(mySlice)-4])
				a += string(mySlice[len(mySlice)-3])
				a += "-"
				a += string(mySlice[len(mySlice)-2])
				a += string(mySlice[len(mySlice)-1])
			}
		}
	} else if len(mySlice)%3 == 2 {
		for i := 0; i < len(mySlice)/3; i++ {
			a += string(mySlice[3*i])
			a += string(mySlice[3*i+1])
			a += string(mySlice[3*i+2])
			if i != len(mySlice)/3-1 {
				a += "-"
			} else {
				a += "-"
				a += string(mySlice[len(mySlice)-2])
				a += string(mySlice[len(mySlice)-1])
			}
		}
	}
	return a
}

func main() {
	fmt.Println(reformatNumber("1  2  3  4  5-6"))
}
