package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := 0
	_, _ = fmt.Scanf("%d", &lines)
	hashMap := make(map[string]string)
	for i := 0; i < lines; i++ {

		//3
		//Write 0x100 8 00001122AABBCCDD
		//Read 0x100 12
		//Clear

		//var opt string
		//_, _ = fmt.Scan(&opt)
		//input := bufio.NewScanner(os.Stdin)
		//var o1 []string
		//for input.Scan() {
		//	o1 = strings.Split(input.Text(), " ")
		//}
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		t := s.Text()
		o := strings.Fields(t)
		//fmt.Println(o)
		f := o[0]
		if f == "Clear" {
			hashMap = make(map[string]string)
		} else if f == "Write" {
			addr := o[1]
			trAddr, _ := strconv.ParseInt(addr[2:], 16, 32)
			if trAddr > 1<<35 {
				continue
			}
			length, _ := strconv.Atoi(o[2])
			trLen, _ := strconv.ParseInt(o[2], 10, 64)
			if trLen == math.MaxInt64 {
				continue
			}
			data := o[3]
			// length = 3,data = aqwccxc
			if 2*length < len(data) {
				data = data[:2*length]
			} else if 2*length > len(data) {
				for index := 0; index < 2*length-len(data); index++ {
					data += "0"
				}
			}
			hashMap[addr] = data
		} else {
			addr := o[1]
			trAddr, _ := strconv.ParseInt(addr[2:], 16, 32)
			if trAddr > 1<<35 {
				continue
			}
			length, _ := strconv.Atoi(o[2])
			trLen, _ := strconv.ParseInt(o[2], 10, 64)
			if trLen == math.MaxInt64 {
				continue
			}
			// 12 16   16 + 8*0
			if v, ok := hashMap[addr]; ok {
				if len(v) < 2*length {
					for index := 0; index < 2*length-len(v); index++ {
						v += "0"
					}
					fmt.Println(v)
				} else if len(v) > 2*length {
					// lenv = 10 , len = 3
					fmt.Println(v[:2*length])
				} else {
					fmt.Println(v)
				}
			} else {
				tmp := ""
				for index := 0; index < 2*length; index++ {
					tmp += "0"
				}
				fmt.Println(tmp)
			}
		}
	}
}
