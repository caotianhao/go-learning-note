package main

import "fmt"

type OrderedStream struct {
	osMap map[int]string
	ptr   int
}

func Constructor1656(n int) OrderedStream {
	myOsMap := make(map[int]string, n)
	return OrderedStream{
		osMap: myOsMap,
		ptr:   1,
	}
}

func (os *OrderedStream) Insert(idKey int, value string) []string {
	os.osMap[idKey] = value
	var ret []string
	flag, tmp := 0, os.ptr
	_, ok := os.osMap[os.ptr]
	for ok {
		flag++
		ret = append(ret, os.osMap[tmp])
		tmp++
		_, ok = os.osMap[tmp]
	}
	if flag != 0 {
		os.ptr += flag
	}
	return ret
}

func main() {
	a := Constructor1656(5)

	fmt.Println("3cc =", a.Insert(3, "cc"))
	fmt.Println("a.Insert(3, \"cc\") =", a)

	fmt.Println("1aa =", a.Insert(1, "aa"))
	fmt.Println("a.Insert(1, \"aa\") =", a)

	fmt.Println("2bb =", a.Insert(2, "bb"))
	fmt.Println("a.Insert(2, \"bb\") =", a)

	fmt.Println("5ee =", a.Insert(5, "ee"))
	fmt.Println("a.Insert(5, \"ee\") =", a)

	fmt.Println("4dd =", a.Insert(4, "dd"))
	fmt.Println("a.Insert(4, \"dd\") =", a)
}
