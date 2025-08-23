package main

import (
	"fmt"
	"strconv"
)

//func getFolderNames(names []string) (ret []string) {
//	total := map[string]int{}
//	for _, v := range names {
//		_, ok := total[v]
//		if !ok {
//			ret = append(ret, v)
//		} else {
//			i := 1
//			tmp := v
//			//v += "(" + fmt.Sprint(i) + ")"
//			v += "(" + strconv.Itoa(i) + ")"
//			_, ok = total[v]
//			for ok {
//				i++
//				v = tmp
//				v += "(" + strconv.Itoa(i) + ")"
//				_, ok = total[v]
//			}
//			ret = append(ret, v)
//		}
//		total[v]++
//	}
//	return
//}

func getFolderNames(names []string) []string {
	ans := make([]string, len(names))
	index := map[string]int{}
	for p, name := range names {
		i := index[name]
		if i == 0 {
			index[name] = 1
			ans[p] = name
			continue
		}
		for index[name+"("+strconv.Itoa(i)+")"] > 0 {
			i++
		}
		t := name + "(" + strconv.Itoa(i) + ")"
		ans[p] = t
		index[name] = i + 1
		index[t] = 1
	}
	return ans
}

func main() {
	//fmt.Println(getFolderNames([]string{"pes", "fifa", "gta", "pes(2019)"}))
	//fmt.Println(getFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
	//fmt.Println(getFolderNames([]string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}))
	//fmt.Println(getFolderNames([]string{"wano", "wano", "wano", "wano"}))
	//fmt.Println(getFolderNames([]string{"kaido", "kaido(1)", "kaido", "kaido(1)"}))
	//fmt.Println(getFolderNames([]string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)"}))
	//fmt.Println(getFolderNames([]string{"a", "a", "a(2)", "a"}))
	//fmt.Println(getFolderNames([]string{"l", "p(4)(4)", "v", "c", "i", "k(2)", "y",
	//	"a(4)(2)", "t", "a", "e", "e(1)(4)", "p", "p", "s", "z", "s", "q(2)(1)",
	//	"r(1)", "m", "b(3)", "u(2)", "j", "p", "t", "s", "g", "o", "o", "b(3)(3)",
	//	"k", "w(2)(3)", "q", "a", "z", "t(4)(4)", "s(2)", "c", "w", "u", "h", "g",
	//	"b", "r", "c(1)", "v", "n(1)", "r(2)(1)", "u", "t(2)(3)", "p", "m(3)(1)",
	//	"h", "o", "y", "n", "s", "z", "k(2)", "x(4)", "l(1)(4)", "g(2)", "u", "t(2)",
	//	"m", "c", "t", "g", "c", "a(2)", "r", "d", "y", "b", "p", "m", "w(2)(4)", "v",
	//	"t(4)(2)", "x", "e(4)(1)", "h", "f", "z", "e(1)(4)", "t", "b(1)", "x(4)", "m(3)",
	//	"j(2)(4)", "s(3)(2)", "z", "l", "p(2)(2)", "g(3)(2)", "q(1)(4)", "h(1)(1)", "h",
	//	"o(3)", "h", "f", "n(4)", "s(4)", "g", "s", "r", "n(1)(1)", "x(3)", "r", "f(3)(1)",
	//	"e", "j", "s", "g", "d", "l", "g", "o(4)(3)", "x(4)", "u(2)(4)", "x", "t", "f",
	//	"j(1)", "v(2)", "w", "v", "t(2)(2)", "l", "o(1)(1)", "a", "y", "q(4)", "m(1)(2)",
	//	"i", "u", "l", "c(1)", "g", "l(2)", "p(1)(1)", "k", "d", "l", "o", "m", "i(1)",
	//	"j", "i", "f", "y", "b", "k", "z", "n", "t", "c(2)", "y", "q", "b", "t", "m",
	//	"g(1)", "r", "m", "l", "s", "n", "j(4)(4)", "m(3)(3)", "n", "n(2)(3)", "s", "t",
	//	"l", "e", "p", "q(2)", "v", "v", "b", "w", "m", "p", "g", "h(2)", "n(1)", "q",
	//	"x(4)", "q(2)(4)", "s(2)", "w", "k", "f", "v", "n", "q", "w", "y(3)", "a(4)(1)",
	//	"r(1)", "r", "f(4)(2)", "l", "f(2)(3)", "p", "o", "h", "t", "i", "r", "k", "p",
	//	"l", "o(3)(4)", "f(1)", "j(4)", "f(3)(4)", "o(2)(1)", "j", "k(2)", "k", "x(4)",
	//	"s(3)(4)", "p(1)(3)", "y", "r(2)", "i(3)(4)", "j", "p(4)(3)", "e", "j(4)", "g(4)(2)",
	//	"x", "l", "x", "g", "w", "o", "t(2)(4)", "s", "f", "o", "h", "h(2)(4)", "z(4)(2)",
	//	"w", "l", "n", "q", "w", "l", "a", "o", "v", "h", "b", "v", "b", "k", "x(2)",
	//	"r(2)(1)", "g(2)", "c(4)(3)", "w(1)", "g(4)", "z", "i(3)(2)", "r", "e", "p",
	//	"z(4)(2)", "s", "c(1)", "h", "l", "j(3)(4)", "x(1)(2)", "a", "b", "t", "c"}))
	//fmt.Println(getFolderNames([]string{"a", "a(1)(2)", "a", "a(1)(4)", "a(1)", "a(1)", "a(1)"}))
	fmt.Println(getFolderNames([]string{"i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i"}))
}
