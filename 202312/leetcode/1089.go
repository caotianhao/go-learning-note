package main

func duplicateZeros(arr []int) {
	l := len(arr)
	for i := 0; i < l; {
		if arr[i] != 0 {
			i++
		} else {
			for j := l - 2; j >= i+1; j-- {
				arr[j+1] = arr[j]
			}
			if i+1 < l {
				arr[i+1] = 0
			}
			i += 2
		}
	}
}

func main() {
	arr := []int{1, 0, 0, 0, 0, 0, 0, 0}
	duplicateZeros(arr)
}
