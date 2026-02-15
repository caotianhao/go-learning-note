package main

import "fmt"

var flag = false

func stackCatalan(queue []byte, l, r int) {
	out, mid, in := 0, 0, 0
	if l == r {
		flag = true
		for out = 0; out < r-1; out++ {
			for mid = out + 1; mid < r; mid++ {
				for in = mid + 1; in <= r; in++ {
					if queue[mid] < queue[in] && queue[in] < queue[out] {
						flag = false
					}
				}
			}
		}
		if flag {
			fmt.Println(string(queue))
		}
	}
	for i := l; i <= r; i++ {
		tmp := queue[i]
		queue[i] = queue[l]
		queue[l] = tmp
		stackCatalan(queue, l+1, r)
		queue[l] = queue[i]
		queue[i] = tmp
	}
}

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)
	stackCatalan([]byte(s), 0, len(s)-1)
}
