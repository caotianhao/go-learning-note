package main

import "fmt"

func countCharacters(words []string, chars string) int {
	charMap, wordsMap := map[int32]int{}, map[int32]int{}
	length := 0
	for _, v := range chars {
		charMap[v]++
	}
loop:
	for i := 0; i < len(words); i++ {
		wordsMap = map[int32]int{}
		for _, v := range words[i] {
			wordsMap[v]++
		}
		for iw, vw := range wordsMap {
			if vc, ok := charMap[iw]; !ok || vc < vw {
				continue loop
			}
		}
		length += len(words[i])
	}
	return length
}

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
	fmt.Println(countCharacters([]string{"dyiclysmffuhibgfvapygkorkqllqlvokosagyelotobicwcmebnp" +
		"znjbirzrzsrtzjxhsfpiwyfhzyonmuabtlwin", "ndqeyhhcquplmznwslewjzuyfgklssvkqxmqjpwhrshycmvrb",
		"ulrrbpspyudncdlbkxkrqpivfftrggemkpyjl", "boygirdlggnh", "xmqohbyqwagkjzpyawsydmdaattthmuvjbzwpyopyafphx",
		"nulvimegcsiwvhwuiyednoxpugfeimnnyeoczuzxgxbqjvegcxeqnjbwnbvowastqhojepisusvsidhqmszbrnynkyop",
		"hiefuovybkpgzygprmndrkyspoiyapdwkxebgsmodhzpx", "juldqdzeskpffaoqcyyxiqqowsalqumddcufhouhrskozhlmobi" +
			"wzxnhdkidr", "lnnvsdcrvzfmrvurucrzlfyigcycffpiuoo", "oxgaskztzroxuntiwlfyufddl", "tfspedteabxatkaypitjf" +
			"khkkigdwdkctqbczcugripkgcyfezpuklfqfcsccboarbfbjfrkxp", "qnagrpfzlyrouolqquytwnwnsqnmuzphne", "eeilfd" +
			"aookieawrrbvtnqfzcricvhpiv", "sisvsjzyrbdsjcwwygdnxcjhzhsxhpceqz", "yhouqhjevqxtecomahbwoptzlkyvjexhzcb" +
			"ccusbjjdgcfzlkoqwiwue", "hwxxighzvceaplsycajkhynkhzkwkouszwaiuzqcleyflqrxgjsvlegvupzqijbornbfwpefhxekgpu" +
			"vgiyeudhncv", "cpwcjwgbcquirnsazumgjjcltitmeyfaudbnbqhflvecjsupjmgwfbjo", "teyygdmmyadppuopvqdodaczob",
		"qaeowuwqsqffvibrtxnjnzvzuuonrkwpysyxvkijemmpdmtnqxwekbpfzs", "qqxpxpmemkldghbmbyxpkwgkaykaerhmwwjonr" +
			"hcsubchs"}, "usdruypficfbpfbivlrhutcgvyjenlxzeovdyjtgvvfdjzcmikjraspdfp"))
}
