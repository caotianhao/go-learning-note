package main

import "fmt"

func seqSearch(arr [20]int, num int) int {
	for i := 0; i < len(arr); i++ {
		if num == arr[i] {
			return i
		}
	}
	return -1
}

func binSearch(arr [20]int, num int) int {
	for low, high := 0, len(arr)-1; low <= high; {
		mid := (low + high) / 2
		if num == arr[low] {
			return low
		} else if num == arr[high] {
			return high
		} else if num == arr[mid] {
			return mid
		} else if num > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func myBubbleSort(arr *[20]int) [20]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return *arr
}

func main() {
	arr := [20]int{84, 66, 50, 49, 94, 86, 55, 90, 43, 3, 57, 25, 93, 78, 44, 34, 14, 5, 15, 96}
	fmt.Println(seqSearch(arr, 94))
	fmt.Println(myBubbleSort(&arr))
	fmt.Println(binSearch(arr, 58))
}
