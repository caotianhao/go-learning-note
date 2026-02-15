package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"go-cth/202408/model"
)

func getCDKGo() bool {
	var response model.Response
	resp, err := http.Post(model.Cdk, "application/json", nil)
	if err != nil {
		fmt.Println("Error posting CDK request:", err)
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return false
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return false
	}
	return response.Code == 0
}

func getOrderGo() bool {
	var response model.Response1
	resp, err := http.Post(model.Order, "application/json", nil)
	if err != nil {
		fmt.Println("Error posting Order request:", err)
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return false
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return false
	}
	return response.Code == 0
}

func main() {
	cdk1, order1 := 0, 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	// 创建一个限制并发数的 chan
	concurrencyLimit := make(chan struct{}, 100) // 允许的最大并发 goroutine 数量

	for i := 0; i < model.Success; i++ {
		wg.Add(2)

		// 获取 chan 来控制并发数
		concurrencyLimit <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-concurrencyLimit }() // 完成时释放 chan
			if getCDKGo() {
				mu.Lock()
				cdk1++
				mu.Unlock()
			}
		}()

		concurrencyLimit <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-concurrencyLimit }()
			if getOrderGo() {
				mu.Lock()
				order1++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("cdkey success = %.2f%%\norder success = %.2f%%\n",
		float64(cdk1)/float64(model.Success)*100,
		float64(order1)/float64(model.Success)*100)
}
