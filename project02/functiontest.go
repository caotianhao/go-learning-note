package project02

type User struct {
	Name string
	Age  int
}

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func Plus(a int, b int) int {
	return a * b
}

func MyBubbleSort(arr *[20]int) [20]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return *arr
}
