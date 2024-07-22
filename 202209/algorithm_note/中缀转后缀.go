package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	top        int
	sliceStack []string
}

func (st *stack) push(n string) {
	st.top++
	st.sliceStack = append(st.sliceStack, n)
}

func (st *stack) pop() string {
	if st.empty() {
		panic("empty Stack")
	}
	tmp := st.sliceStack[len(st.sliceStack)-1]
	st.sliceStack = st.sliceStack[:len(st.sliceStack)-1]
	st.top--
	return tmp
}

func (st *stack) empty() bool {
	return st.top == -1
}

func (st *stack) size() int {
	return st.top + 1
}

func (st *stack) getTop() string {
	if st.empty() {
		panic("empty stack")
	}
	return st.sliceStack[st.top]
}

func stringToSlice(str string) (sliceInfix []string) {
	s := ""
	for _, v := range str {
		if v > '9' || v < '0' {
			if s != "" {
				sliceInfix = append(sliceInfix, s)
			}
			sliceInfix = append(sliceInfix, string(v))
			s = ""
			continue
		}
		s += string(v)
	}
	if s != "" {
		sliceInfix = append(sliceInfix, s)
	}
	return
}

func infixToPostfix(infix []string) (postfix []string) {
	opMap := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}
	numStack, opStack := stack{-1, make([]string, 0)}, stack{-1, make([]string, 0)}
	for i := 0; i < len(infix); i++ {
		if infix[i] == "+" || infix[i] == "-" || infix[i] == "*" || infix[i] == "/" {
			if opStack.empty() || opMap[infix[i]] > opMap[opStack.getTop()] {
				opStack.push(infix[i])
			} else {
				for !opStack.empty() && opMap[infix[i]] <= opMap[opStack.getTop()] {
					numStack.push(opStack.pop())
				}
				opStack.push(infix[i])
			}
		} else {
			numStack.push(infix[i])
		}
	}
	for !opStack.empty() {
		numStack.push(opStack.pop())
	}
	return numStack.sliceStack
}

func calPostfix(postfix []string) float64 {
	numStack := stack{-1, make([]string, 0)}
	for _, v := range postfix {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			num2, num1 := numStack.pop(), numStack.pop()
			n2, _ := strconv.ParseFloat(num2, 64)
			n1, _ := strconv.ParseFloat(num1, 64)
			switch v {
			case "+":
				numStack.push(fmt.Sprintf("%f", n1+n2))
			case "-":
				numStack.push(fmt.Sprintf("%f", n1-n2))
			case "*":
				numStack.push(fmt.Sprintf("%f", n1*n2))
			case "/":
				numStack.push(fmt.Sprintf("%f", n1/n2))
			}
		} else {
			numStack.push(v)
		}
	}
	ret, _ := strconv.ParseFloat(numStack.getTop(), 64)
	return ret
}

func main() {
	var str string
	_, _ = fmt.Scanln(&str)

	infix := stringToSlice(str)
	fmt.Println("Infix:", infix)

	postfix := infixToPostfix(infix)
	fmt.Println("Postfix:", postfix)

	ret := calPostfix(postfix)
	fmt.Println(ret)
}
