package main

import (
	"errors"
	"fmt"
)

// 自定义错误
var testError = errors.New("this is a test error")

func getError() error {
	return testError
}

func getError1() error {
	// 在 Go 1.13 及更高版本中
	// %w 是用于格式化错误字符串时的特殊标记,这是一种错误包装（error wrapping）的方式
	// %w 用于将当前错误包装在新的错误中，目的是在错误链中保留原始错误
	// 使错误信息更加详细，并且在处理错误时能够追溯到根本原因
	return fmt.Errorf("通过 %%w 构造的新错误 %w", testError)
}

func main() {
	// 获取错误
	// getError 返回 testError
	// 而 getError1 返回的错误是通过 fmt.Errorf("e: %w", testError) 构造的一个新错误
	// 其中包含 testError 作为原始错误
	e := getError()
	fmt.Println("e =", e)

	// 编译器的进步告诉我，fmt.Println(e == testError) 这种写法已经是 warning 了
	// 所以还是用 errors.Is 吧             ------2024.07.26
	//fmt.Println(e == testError)          // true

	fmt.Println(errors.Is(e, testError)) // true

	// errors.Is 用于检查错误链中是否包含特定类型的错误
	e1 := getError1()
	fmt.Println("e1 =", e1)
	//fmt.Println(e1 == testError)          // false
	fmt.Println(errors.Is(e1, testError)) // true
}
