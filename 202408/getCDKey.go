package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response 定义响应结构体
type Response struct {
	Code int    `json:"code"`
	Data Data   `json:"data"`
	Msg  string `json:"msg"`
}

type Data struct {
	CDKey string `json:"cd_key"`
}

func main() {
	// 定义要访问的URL
	url := "http://localhost:8200/heishenhua/cd_key/create"

	// 发送HTTP GET请求
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// 读取响应体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印原始的响应内容
	fmt.Println("Response Body:")
	fmt.Println(string(body))

	// 定义一个Response类型的变量来存储解码后的数据
	var response Response

	// 解析JSON响应
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 提取cd_key字段并打印
	cdKey := response.Data.CDKey
	fmt.Println("cd_key:", cdKey)
}
