package main

import (
	"fmt"
	"strconv"
)

// 4540181102 -> 45.40.181.102
func getAllRuleIP(originIP string) (res []string) {
	var ans string
	l := len(originIP)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				for m := 1; m <= 3; m++ {
					if i+j+k+m == l {
						ii, _ := strconv.Atoi(originIP[:i])
						jj, _ := strconv.Atoi(originIP[i : i+j])
						kk, _ := strconv.Atoi(originIP[i+j : i+j+k])
						mm, _ := strconv.Atoi(originIP[i+j+k : i+j+k+m])
						if ii <= 255 && jj <= 255 && kk <= 255 && mm <= 255 {
							ans = fmt.Sprintf("%d.%d.%d.%d", ii, jj, kk, mm)
							if len(ans) == l+3 {
								res = append(res, ans)
							}
						}
					}
				}
			}
		}
	}
	return
}

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)
	ip := getAllRuleIP(s)
	for _, ipv := range ip {
		fmt.Println(ipv)
	}
}
