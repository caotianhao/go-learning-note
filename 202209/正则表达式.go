package main

import (
	"fmt"
	"regexp"
)

func main() {
	//在字符串中查找与指定正则表达式匹配的子串
	str := "Chocolate is my favorite!"

	//a 是否在 str 中
	a := "Chocolate"
	isHave, _ := regexp.MatchString(a, str)
	fmt.Println(isHave) //true

	a2 := "chocolate"
	isHave2, _ := regexp.MatchString(a2, str)
	//正则表达式是分大小写的
	fmt.Println(isHave2) //false

	//如果需要忽略大小写的话，需要特殊语法,在字符串前面加上 (?i)
	a3 := "(?i)chocolate"
	isHave3, _ := regexp.MatchString(a3, str)
	fmt.Println(isHave3) //true

	//-----------------------------------------------------------------------
	//正则表达式可用于验证程序的输入数据，这是一种分析和理解数据的高效方式
	//要将正则表达式赋给变量，必须先对其进行分析
	//用于分析正则表达式的函数有两个
	//1.Compile：在正则表达式未能通过编译时返回错误
	//2.MustCompile：在正则表达式无法编译时引发 panic
	//该使用哪一个取决于具体情况，但 MustCompile 通常是更佳的选择
	//-----------------------------------------------------------------------
	// . 与除换行符之前的其他任何字符都匹配
	// * 与零个或多个指定的字符匹配
	// ^ 表示行首
	// $ 表示行尾
	// + 匹配一次或多次
	// ? 匹配零或一次
	// [] 与方括号内指定的任何字符都匹配
	// {n} 匹配 n 次
	// {n,} 匹配 n 次或更多次
	// {m,n} 最少匹配 m 次，最多匹配 n 次，两边都是闭区间
	//-----------------------------------------------------------------------
	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}$")
	fmt.Println(re.MatchString("study67890123")) //false, 因为有 13 次匹配
	fmt.Println(re.MatchString("study!"))        //false, 因为有叹号
	fmt.Println(re.MatchString("G"))             //false, 因为有 1 次匹配
	fmt.Println(re.MatchString("caoTianHao"))    //true

	//-----------------------------------------------------------------------
	//有些用户名是无效的，因为其长度不正确或包含非法字符
	//可以对数据进行清洗，以便能够安全地使用它们
	//有关用户名的规则：包含的字符不能超过 12，要清洗太长的用户名
	//可先将其截断为只包含 12 个字符，然后根据正则表达式对其进行评估，看看它是否包含非法字符
	//如果包含非法字符，可将其替换为合法字符
	//-----------------------------------------------------------------------
	usernames := []string{"caoTianHao", "0123456789101112131415", "!User%Name@"}
	an := regexp.MustCompile("[[:^alnum:]]")
	for _, username := range usernames {
		if len(username) > 12 {
			username = username[:12]
			fmt.Println("截断：", username) //截断： 012345678910
		}
		if !re.MatchString(username) {
			//把所有违规字符替换为 x
			username = an.ReplaceAllLiteralString(username, "X")
			fmt.Println("替换：", username) //替换： XUserXNameX
		}
	}
}
