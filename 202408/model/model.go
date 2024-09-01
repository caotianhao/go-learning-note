package model

type Response struct {
	Code int    `json:"code"`
	Data Data   `json:"data"`
	Msg  string `json:"msg"`
}

type Response1 struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type Data struct {
	CDKey string `json:"cd_key"`
}

const (
	Cdk     = "http://localhost:8200/heishenhua/cd_key/create"
	Order   = "http://localhost:8200/shunfeng/order/create"
	Success = 1800
)
