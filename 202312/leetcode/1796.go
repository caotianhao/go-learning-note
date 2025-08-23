package main

import (
	"fmt"

	"github.com/Arafatk/DataViz/trees/redblacktree"
)

func secondHighest(s string) int {
	t := redblacktree.NewWithIntComparator()
	for _, v := range s {
		if '0' <= v && v <= '9' {
			t.Put(int(v-48), nil)
			if t.Size() > 2 {
				t.Remove(t.Left().Key)
			}
		}
	}
	if t.Size() == 2 {
		return t.Left().Key.(int)
	}
	return -1
}

func main() {
	fmt.Println(secondHighest("0aff31fh3d4fg5d41")) //4
}
