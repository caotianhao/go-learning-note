package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func readFile(file *os.File) {
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
	fmt.Println()
}

func main() {
	f, err := os.Open("no_exist.txt")
	if err != nil {
		log.Println(err)
	} else {
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Println("failed to close f")
			}
		}(f)
		readFile(f)
	}

	f2, err := os.Open("exist.txt")
	if err != nil {
		log.Println(err)
	} else {
		defer func(f2 *os.File) {
			err := f2.Close()
			if err != nil {
				log.Println("failed to close f2")
			}
		}(f2)
		readFile(f2)
	}
}
