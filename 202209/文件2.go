package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//不用打开也不用关闭
	//适合小文件
	file, err := os.ReadFile("../project02/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", file)
}
