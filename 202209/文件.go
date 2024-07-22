package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	myFile, err := os.Open("../project02/test.txt")
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	defer func(myFile *os.File) {
		err2 := myFile.Close()
		if err2 != nil {
			fmt.Println("err2 =", err2)
		}
	}(myFile)

	//读
	reader := bufio.NewReader(myFile)
	for {
		str, err3 := reader.ReadString('\n')
		fmt.Print(str)
		if err3 == io.EOF {
			break
		}
	}
}
