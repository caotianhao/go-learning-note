package main

import "fmt"

type MyLinkedList struct {
	head *trueList
	size int
}

type trueList struct {
	val  int
	next *trueList
}

func Constructor707() MyLinkedList {
	return MyLinkedList{&trueList{}, 0}
}

func (list *MyLinkedList) Get(index int) int {
	if index < 0 || index >= list.size {
		return -1
	}
	cur := list.head
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (list *MyLinkedList) AddAtHead(val int) {
	list.AddAtIndex(0, val)
}

func (list *MyLinkedList) AddAtTail(val int) {
	list.AddAtIndex(list.size, val)
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index >= 0 && index <= list.size {
		list.size++
		cur := list.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		tmp := &trueList{val: val, next: cur.next}
		cur.next = tmp
	}
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index >= 0 && index < list.size {
		cur := list.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		cur.next = cur.next.next
		list.size--
	}
}

func main() {
	myLinkedList := Constructor707()
	a := myLinkedList.head
	for a != nil {
		fmt.Print(a.val, " ")
		a = a.next
	}
	fmt.Println()

	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	a = myLinkedList.head
	for a != nil {
		fmt.Print(a.val, " ")
		a = a.next
	}
	fmt.Println()

	myLinkedList.AddAtIndex(1, 2) // 链表变为 1->2->3
	a = myLinkedList.head
	for a != nil {
		fmt.Print(a.val, " ")
		a = a.next
	}
	fmt.Println()

	fmt.Println("get1", myLinkedList.Get(1)) // 返回 2

	myLinkedList.DeleteAtIndex(1) // 现在，链表变为 1->3
	a = myLinkedList.head
	for a != nil {
		fmt.Print(a.val, " ")
		a = a.next
	}
	fmt.Println()

	fmt.Println("get1", myLinkedList.Get(1)) // 返回 3
}
