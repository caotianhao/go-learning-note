package main

import (
	"container/list"
	"fmt"
)

func makeGood(s string) (res string) {
	var strStack list.List
	for i := range s {
		// 这就是一个栈的应用的最简单的题...思考了很久
		// 栈不为空就入栈
		// 或者不满足相消的条件也入栈
		// 就这么简单不应该这么久才做出来的
		if strStack.Len() == 0 || s[i]-strStack.Back().Value.(byte) != 32 &&
			strStack.Back().Value.(byte)-s[i] != 32 {
			strStack.PushBack(s[i])
		} else {
			strStack.Remove(strStack.Back())
		}
	}
	for i := strStack.Front(); i != nil; i = i.Next() {
		res += string(i.Value.(uint8))
	}
	return
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
}
