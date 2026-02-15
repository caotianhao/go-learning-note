package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func funTimeWithBooks(books []int) (maxFun int) {
	n := len(books)
	sort.Ints(books)
	partFun := 0
	for i := n - 1; i >= 0; i-- {
		if partFun+books[i] <= 0 {
			break
		}
		partFun += books[i]
		maxFun += partFun
	}
	return maxFun
}

func main() {
	var books string
	bookArray := make([]int, 0)
	_, _ = fmt.Scanf("%s", &books)
	book := strings.Split(books, ",")
	for _, b := range book {
		tmp, _ := strconv.Atoi(b)
		bookArray = append(bookArray, tmp)
	}
	fmt.Println(funTimeWithBooks(bookArray))
}
