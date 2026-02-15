package main

import "fmt"

// 题本身不难，但是思考了很久
// map和slice应用还不熟练?
// 逻辑问题，很绕！！
func uniqueMorseRepresentations(words []string) int {
	letterAndMorseMap, letter := make(map[string]string), make([]string, 26)
	wordSlice, trueWordSliceMap := make([]string, 0), make(map[string]int)
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---",
		"-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-",
		".--", "-..-", "-.--", "--.."}
	for i := 0; i < 26; i++ {
		letter[i] = fmt.Sprintf("%d", i+97)
		letterAndMorseMap[letter[i]] = morse[i]
	}
	for i := 0; i < len(words); i++ {
		var temp string
		tempString := words[i]
		for i := 0; i < len(tempString); i++ {
			temp += letterAndMorseMap[fmt.Sprintf("%d", tempString[i])]
		}
		wordSlice = append(wordSlice, temp)
	}
	for i := 0; i < len(wordSlice); i++ {
		trueWordSliceMap[wordSlice[i]]++
	}
	return len(trueWordSliceMap)
}

func main() {
	arr := []string{"gin", "zen", "gig", "msg", "cth", "jyz"}
	fmt.Println(uniqueMorseRepresentations(arr))
}
