package main

import (
	"fmt"
)

func minPartitions(n string) (max int) {
	for _, v := range n {
		tmp := int(v - 48)
		if tmp == 9 {
			return 9
		}
		if tmp > max {
			max = tmp
		}
	}
	return
}

func main() {
	fmt.Println(minPartitions("54685413564"))  //8
	fmt.Println(minPartitions("546854135649")) //9
	fmt.Println(minPartitions("11"))           //1
}
