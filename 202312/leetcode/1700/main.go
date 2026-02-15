package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	for len(students) != 0 {
		if students[0] == sandwiches[0] {
			students, sandwiches = popStack(students), popStack(sandwiches)
		} else {
			students = reSortQueue(students)
		}
		if len(students) != 0 &&
			isArrOnlyHaveOneKindNum(students) && students[0] != sandwiches[0] {
			return len(students)
		}
	}
	return len(students)
}

func reSortQueue(queue []int) []int {
	qu := make([]int, 0)
	for i := 1; i < len(queue); i++ {
		qu = queue[1:]
	}
	qu = append(qu, queue[0])
	return qu
}

func popStack(stack []int) []int {
	if len(stack) == 1 {
		return []int{}
	} else {
		return stack[1:]
	}
}

func isArrOnlyHaveOneKindNum(arr []int) bool {
	ele := arr[0]
	for _, v := range arr {
		if v != ele {
			return false
		}
	}
	return true
}

func main() {
	//student := []int{1, 0, 1, 0, 1, 0}
	//l := len(student)
	//fmt.Println(l)
	//student = reSortQueue(student)
	//student = popStack(student)
	//fmt.Println(student, len(student))
	//fmt.Println(reSortQueue([]int{1}))
	//fmt.Println(reSortQueue([]int{}))
	//fmt.Println(popStack(popStack([]int{1, 0, 1, 0, 1, 0})))
	//fmt.Println(popStack([]int{1}))
	//fmt.Println(popStack([]int{}))
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
	fmt.Println(countStudents([]int{1}, []int{0}))
	fmt.Println(countStudents([]int{1, 1, 1, 1, 1, 1}, []int{0, 0, 0, 0, 0, 0}))
}
