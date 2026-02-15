package main

import "fmt"

func streamControl(minute, maxRequest, n int, request []int) int {
	window := 0
	flow := 0
	for i := 0; i < n; i++ {
		if i >= minute {
			window -= request[i-minute]
		}
		window += request[i]
		if window > maxRequest {
			diff := window - maxRequest
			if diff > flow {
				flow = diff
			}
			window = maxRequest
		}
	}
	return flow
}

func main() {
	minute, maxRequest := 0, 0
	_, _ = fmt.Scanf("%d %d", &minute, &maxRequest)
	n, r := 0, 0
	_, _ = fmt.Scanf("%d", &n)
	request := make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &r)
		request = append(request, r)
	}
	fmt.Println(streamControl(minute, maxRequest, n, request))
}
