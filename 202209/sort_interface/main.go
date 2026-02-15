package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type hero struct {
	name string
	age  int
}

// 这个要给个别名，不给会出错
type heroSlice []hero

func (h heroSlice) Len() int {
	return len(h)
}

func (h heroSlice) Less(i, j int) bool {
	//按年龄升序排序，如果是 > 就代表降序排序
	return h[i].age < h[j].age
}

func (h heroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())
	var myHeroSlice heroSlice
	for i := 0; i < 10; i++ {
		myHero := hero{
			//随机生成英雄的名字和年龄
			name: "hero" + fmt.Sprintf("%d", i+1),
			age:  rand.Intn(179) + 20,
		}
		myHeroSlice = append(myHeroSlice, myHero)
	}
	fmt.Println("before sort:", myHeroSlice)
	// Len,Less,Swap是sort.Sort接口的三个方法，要实现才可排序
	sort.Sort(myHeroSlice)
	fmt.Println("after sort: ", myHeroSlice)
}
