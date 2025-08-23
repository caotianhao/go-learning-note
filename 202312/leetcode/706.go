package main

import "fmt"

type MyHashMap struct {
	key   []int
	value []int
}

func Constructor706() MyHashMap {
	return MyHashMap{
		key:   []int{},
		value: []int{},
	}
}

func (m *MyHashMap) Put(key int, value int) {
	for i := 0; i < len(m.key); i++ {
		if m.key[i] == key {
			m.value[i] = value
			return
		}
	}
	m.key = append(m.key, key)
	m.value = append(m.value, value)
}

func (m *MyHashMap) Get(key int) int {
	for i := 0; i < len(m.key); i++ {
		if m.key[i] == key {
			return m.value[i]
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	ind := -1
	for i := 0; i < len(m.key); i++ {
		if m.key[i] == key {
			ind = i
		}
	}
	if ind == -1 {
		return
	}
	m.key = delete706(m.key, ind)
	m.value = delete706(m.value, ind)
}

func delete706(arr []int, index int) []int {
	slice706 := arr[:index]
	for i := index + 1; i < len(arr); i++ {
		slice706 = append(slice706, arr[i])
	}
	return slice706
}

func main() {
	con := Constructor706()
	fmt.Println("Constructor()", con.key, con.value)
	con.Put(1, 1)
	fmt.Println("Put(1, 1)    ", con.key, con.value)
	con.Put(2, 2)
	fmt.Println("Put(2, 2)    ", con.key, con.value)
	a := con.Get(3)
	fmt.Println("Get(3)       ", a)
	con.Put(2, 1)
	fmt.Println("Put(2, 1)    ", con.key, con.value)
	b := con.Get(2)
	fmt.Println("Get(2)       ", b)
	con.Remove(2)
	fmt.Println("Remove(2)    ", con.key, con.value)
	c := con.Get(2)
	fmt.Println("Get(2)       ", c)
}
