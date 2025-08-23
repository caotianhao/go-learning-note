package main

import "fmt"

// begin 和 end 之间为参考答案得来
// 为动态规划内容
func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	skillLen, skillNum, skillMap := len(reqSkills), 0, map[string]int{}
	peopleLen := len(people)
	peopleNum := make([]int, peopleLen)
	//begin------------------------------------------------
	dp := make([][]int, 1<<skillLen)
	dp[0] = []int{}
	//end--------------------------------------------------
	for i := 0; i < skillLen; i++ {
		skillNum |= 1 << i
		skillMap[reqSkills[i]] = i
	}
	for i := 0; i < peopleLen; i++ {
		for j := 0; j < len(people[i]); j++ {
			peopleNum[i] |= 1 << (skillLen - skillMap[people[i][j]] - 1)
		}
		//begin--------------------------------------------
		for prev := 0; prev < len(dp); prev++ {
			if dp[prev] == nil {
				continue
			}
			comb := prev | peopleNum[i]
			if dp[comb] == nil || len(dp[prev])+1 < len(dp[comb]) {
				dp[comb] = make([]int, len(dp[prev]))
				copy(dp[comb], dp[prev])
				dp[comb] = append(dp[comb], i)
			}
		}
		//end----------------------------------------------
	}
	//begin------------------------------------------------
	return dp[(1<<skillLen)-1]
	//end--------------------------------------------------
}

func main() {
	fmt.Println(smallestSufficientTeam([]string{"java", "nodejs", "reactjs"},
		[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}})) // 0, 2
}
