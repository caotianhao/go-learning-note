package main

import (
	"fmt"
	"strconv"
)

// 这就是力扣通过率 30 的简单会员题吗...

type StringIterator struct {
	strMap     map[string][][2]int
	ptr        int
	trueLength int
}

func Constructor604(compressedString string) StringIterator {
	// id 记录 compressedString 中每个是字母的 id
	id := make([]int, 0)
	for i, v := range compressedString {
		if v >= 'A' && v <= 'Z' || v <= 'z' && v >= 'a' {
			id = append(id, i)
		}
	}

	// 例如将形如 L156e2t1C1o1d1e7 的字符串
	// 拆为 [[L 156] [e 2] [t 1] [C 1] [o 1] [d 1] [e 7]]
	help := make([][]string, 0)
	for i := 0; i < len(id)-1; i++ {
		t := make([]string, 0)
		t = append(t, string(compressedString[id[i]]))
		t = append(t, compressedString[id[i]+1:id[i+1]])
		help = append(help, t)
	}
	t := make([]string, 0)
	t = append(t, string(compressedString[id[len(id)-1]]))
	t = append(t, compressedString[id[len(id)-1]+1:])
	help = append(help, t)

	// 求实际未压缩字符串的长度
	var length int
	for _, v := range help {
		k, _ := strconv.Atoi(v[1])
		length += k
	}

	// sum 各字母出现次数前缀和数组
	// 例如 [[L 156] [e 2] [t 1] [C 1] [o 1] [d 1] [e 7]]
	// 变为 [156 2 1 1 1 1 7]
	// 再变为 [156 158 159 160 161 162 169]
	hou := make([]int, 0)
	for _, v := range help {
		k, _ := strconv.Atoi(v[1])
		hou = append(hou, k)
	}
	sum := make([]int, 0)
	for i := 0; i < len(hou); i++ {
		s := 0
		for j := 0; j <= i; j++ {
			s += hou[j]
		}
		sum = append(sum, s)
	}

	// 把    [[L 156]  [e 2]   [t 1]   [C 1]   [o 1]   [d 1]   [e 7]]
	// 变为  [[L 156]  [e 158] [t 159] [C 160] [o 161] [d 162] [e 169]]
	for i := 0; i < len(help); i++ {
		help[i][1] = strconv.Itoa(sum[i])
	}

	// m 存储每个字母出现的范围，闭区间
	m := map[string][][2]int{}
	k, _ := strconv.Atoi(help[0][1])
	m[help[0][0]] = append(m[help[0][0]], [2]int{0, k - 1})
	for i := 1; i < len(help); i++ {
		k1, _ := strconv.Atoi(help[i-1][1])
		k2, _ := strconv.Atoi(help[i][1])
		m[help[i][0]] = append(m[help[i][0]], [2]int{k1, k2 - 1})
	}

	return StringIterator{
		strMap:     m,
		ptr:        0,
		trueLength: length,
	}
}

func (si *StringIterator) Next() byte {
	if si.ptr < si.trueLength {
		for i, v := range si.strMap {
			for j := 0; j < len(v); j++ {
				if isIn604(si.ptr, v[j]) {
					res := []byte(i)
					si.ptr++
					return res[0]
				}
			}
		}
	}
	return ' '
}

func (si *StringIterator) HasNext() bool {
	return si.ptr < si.trueLength
}

func isIn604(index int, arr [2]int) bool {
	return index >= arr[0] && index <= arr[1]
}

func main() {
	c := Constructor604("L1e2t1C1o1d1e1")
	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.HasNext())
	fmt.Println(c.Next())
	fmt.Println(c.HasNext())
}
