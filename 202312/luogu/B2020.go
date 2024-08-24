package main

import "fmt"

func main() {
	var a, b, c, d, e, eat int
	_, _ = fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)
	arr := []int{a, b, c, d, e}
	eat += arr[0] % 3
	arr[4], arr[1] = arr[4]+arr[0]/3, arr[1]+arr[0]/3
	arr[0] /= 3
	eat += arr[1] % 3
	arr[0], arr[2] = arr[0]+arr[1]/3, arr[2]+arr[1]/3
	arr[1] /= 3
	eat += arr[2] % 3
	arr[3], arr[1] = arr[3]+arr[2]/3, arr[1]+arr[2]/3
	arr[2] /= 3
	eat += arr[3] % 3
	arr[4], arr[2] = arr[4]+arr[3]/3, arr[2]+arr[3]/3
	arr[3] /= 3
	eat += arr[4] % 3
	arr[3], arr[0] = arr[3]+arr[4]/3, arr[0]+arr[4]/3
	arr[4] /= 3
	fmt.Printf("%d %d %d %d %d\n%d", arr[0], arr[1], arr[2], arr[3], arr[4], eat)
}
