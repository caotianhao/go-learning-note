package main

import "fmt"

//给你一个数组 items ，其中items[i] = [type, color, name] ，描述第 i 件物品的类型、颜色以及名称
//另给你一条由两个字符串ruleKey 和 ruleValue 表示的检索规则
//如果第 i 件物品能满足下述条件之一，则认为该物品与给定的检索规则 匹配
//ruleKey == "type" 且 ruleValue == type
//ruleKey == "color" 且 ruleValue == color
//ruleKey == "name" 且 ruleValue == name
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	l, cnt, index := len(items), 0, -1
	if ruleKey == "type" {
		index = 0
	} else if ruleKey == "color" {
		index = 1
	} else {
		index = 2
	}
	for i := 0; i < l; i++ {
		if items[i][index] == ruleValue {
			cnt++
		}
	}
	return cnt
}

func main() {
	items := [][]string{
		{"phone", "blue", "pixel"},
		{"computer", "silver", "lenovo"},
		{"phone", "gold", "iphone"},
	}
	ruleKey := "color"
	ruleValue := "silver"
	fmt.Println(countMatches(items, ruleKey, ruleValue))
}
