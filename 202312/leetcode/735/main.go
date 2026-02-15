package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	// 冷静，思考，分析
	// 返回值，用栈模拟
	stack := make([]int, 0)
	for _, asteroid := range asteroids {
		// 如果小于 0
		if asteroid < 0 {
			// 栈中有元素 && 栈顶大于 0 && 栈顶大小小于该行星
			for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < -asteroid {
				stack = stack[:len(stack)-1]
			}
			// 栈中若还有，则新行星必没
			if len(stack) > 0 && stack[len(stack)-1] > 0 {
				// 相等的话栈顶也没
				if stack[len(stack)-1] == -asteroid {
					stack = stack[:len(stack)-1]
				}
				continue
			}
		}
		stack = append(stack, asteroid)
	}
	return stack
}

func main() {
	fmt.Println(asteroidCollision([]int{10, 2, -6, 10, 2, -6}))
}
