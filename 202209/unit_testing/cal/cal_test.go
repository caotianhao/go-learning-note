package cal

import (
	"fmt"
	"testing"
)

//直接右键 run 的话不行，必须用 go test
//其中	go test 正确的时候不输出日志，仅当错误的时候输出日志
//		go test -v 都输出日志
//测试函数文件名必须是 _test.go 结尾，且函数名必须是 TestXxx 的形式
//(t *testing.T) 是固定写法
//测试单个方法：go test -v -test.run 方法名，如：go test -v -test.run TestHello2
//测试单个文件：go test -v -test.run 方法名，如：go test -v cal.go cal_test.go,一定要带上原文件
func TestAddCal(t *testing.T) {
	res := addCal(10)
	if res != 55 {
		t.Fatalf("err!err!err!")
	}
	t.Logf("yes")
}

func TestHello(t *testing.T) {
	fmt.Println("    TestHello 被调用")
}

func TestHello2(t *testing.T) {
	fmt.Println("    TestHello2 被调用")
}

func TestHello3(t *testing.T) {
	fmt.Println("    TestHello3 被调用")
}

func TestHello4(t *testing.T) {
	fmt.Println("    TestHello4 被调用")
}
