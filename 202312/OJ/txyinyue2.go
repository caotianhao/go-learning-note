package main

import "fmt"

type TreeNodeTxyy struct {
	Val   int
	Left  *TreeNodeTxyy
	Right *TreeNodeTxyy
}

func cntOfMethods(trees []*TreeNodeTxyy) (q int) {
	const mod = int(1e9 + 7)
	ans := make([]int, 0)
	for _, t := range trees {
		ans = append(ans, 2*countSon(t))
	}
	n := len(ans)
	if n == 1 {
		return 1
	}
	fc := facMod(n-1) % mod
	for i := 0; i < n; i++ {
		result := fc
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			result *= ans[j]
			result %= mod
		}
		q += result
	}
	fmt.Println(ans)
	return q
}

func facMod(n int) int {
	mod := int(1e9 + 7)
	if n == 1 {
		return 1
	}
	return n * facMod(n-1) % mod
}

func countSon(root *TreeNodeTxyy) int {
	var dfs func(*TreeNodeTxyy) int
	cnt := 0
	dfs = func(node *TreeNodeTxyy) int {
		if node.Left == nil && node.Right == nil {
			cnt++
		} else {
			if node.Left != nil {
				dfs(node.Left)
			}
			if node.Right != nil {
				dfs(node.Right)
			}
		}
		return cnt
	}
	return dfs(root)
}

func main() {
	a := &TreeNodeTxyy{2, nil, nil}
	t1 := &TreeNodeTxyy{1, nil, a}

	c := &TreeNodeTxyy{4, nil, nil}
	c1 := &TreeNodeTxyy{5, nil, nil}
	t2 := &TreeNodeTxyy{3, c, c1}

	d := &TreeNodeTxyy{4, nil, nil}
	d1 := &TreeNodeTxyy{5, nil, nil}
	t3 := &TreeNodeTxyy{3, d, d1}
	fmt.Println(cntOfMethods([]*TreeNodeTxyy{t1, t2, t3}))
	//fmt.Println(facMod(4))
}
