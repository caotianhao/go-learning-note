package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://studygolang.com/pkgdoc",
		"https://www.baidu.com",
		"https://yiyan.baidu.com",
	}
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("http get error:", err)
			}
			fmt.Println(url, resp.Status, resp.Proto)
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					fmt.Println("Body Close:", err)
				}
			}(resp.Body)
		}(url)
	}
	wg.Wait()
}
