package main

import "fmt"

func countSeniors(details []string) (cnt int) {
	for _, v := range details {
		if (v[11]-48)*10+(v[12]-48) > 60 {
			//fmt.Println((v[11]-48)*10 + (v[12] - 48))
			cnt++
		}
	}
	return
}

func main() {
	//fmt.Println(countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"})) //2
	//fmt.Println(countSeniors([]string{"1313579440F2036", "2921522980M5644"}))                    //0
	fmt.Println(countSeniors([]string{"9751302862F0693", "3888560693F7262", "5485983835F06" +
		"49", "2580974299F6042", "9976672161M6561", "0234451011F8013", "4294552179O6482"})) //4
}
