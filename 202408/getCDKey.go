package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go-cth/202408/model"
)

func getCDK() bool {
	var response model.Response
	resp, err := http.Post(model.Cdk, "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response.Code == 0
}

func getOrder() bool {
	var response model.Response1
	resp, err := http.Post(model.Order, "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response.Code == 0
}

func main() {
	cdk1, order1 := 0, 0
	for i := 0; i < model.Success; i++ {
		if getCDK() {
			cdk1++
		}
		if getOrder() {
			order1++
		}
	}
	fmt.Printf("cdkey success = %.2f%%\norder success = %.2f%%\n",
		float64(cdk1)/float64(model.Success)*100,
		float64(order1)/float64(model.Success)*100)
}
