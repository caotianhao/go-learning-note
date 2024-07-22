package main

import "fmt"

type showInfo interface {
	show() string
	showAdd(s string) string
}

type s1 struct {
	name string
}

type s2 struct {
	name string
}

func interfacE(sh showInfo) string {
	return sh.showAdd("++")
}

func (s s1) show() string {
	return s.name
}

func (s s1) showAdd(ss string) string {
	return s.name + ss
}

func (s s2) show() string {
	return s.name
}

func (s s2) showAdd(ss string) string {
	return s.name + ss
}

func main() {
	s := s1{"111"}
	ss := s2{"222"}
	fmt.Println(s.show(), ss.show())
	fmt.Println(s.showAdd("++"), ss.showAdd("++"))
	fmt.Println(interfacE(s), interfacE(ss))
}
