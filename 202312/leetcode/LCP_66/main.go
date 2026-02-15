package main

import "fmt"

// LCP 的题思路都...很清晰，但是代码实现真的要细节...
func minNumBooths(demand []string) int {
	l, exCnt := len(demand), 0
	//存储暂时这一轮字符串需要的字母数
	tempMap := make(map[string]int)
	//存储总的所需要的字母数
	allMap := make(map[string]int)
	//存储每一轮字符串需要的额外字母数
	exMap := make(map[string]int)
	for i := 0; i < l; i++ {
		//因为是暂时的，所以每轮都要初始化
		tempMap = make(map[string]int)
		for j := 0; j < len(demand[i]); j++ {
			//暂时存储
			tempMap[string(demand[i][j])]++
			//全部需要的字母数，主要是统计需要的种类数，相应地，其 value 用处不大
			allMap[string(demand[i][j])]++
			//如果下一轮需要的额外字母大于等于上一轮需要的额外字母，就把本轮的替换上去
			if tempMap[string(demand[i][j])]-1 >= exMap[string(demand[i][j])] {
				exMap[string(demand[i][j])] = tempMap[string(demand[i][j])] - 1
			} //else {
			//如果下一轮需要的额外字母小于上一轮需要的额外字母，则直接就可以用上一轮的额外字母来支持本轮
			//所以这里注释掉的代码其实是不需要的，如果加上就是双倍的
			//	exMap[string(demand[i][j])] += tempMap[string(demand[i][j])] - 1
			//}
		}
		//打印测试
		//fmt.Println("T", tempMap)
		//fmt.Println("A", allMap)
		//打印测试
		//fmt.Println("E", exMap)

		//for _, v1 := range exMap {
		//	exCnt += v1
		//}
		//fmt.Println("EC", exCnt)
		//fmt.Println(exMap)
		//fmt.Println(len(allMap), "---------", demand[i], "-------------", exCnt)
	}
	//统计 exMap 每个字母的额外使用量
	for _, v1 := range exMap {
		//exCnt 就是所有额外使用量的和
		exCnt += v1
	}
	//allMap 的长度就是本来基本所需的字母量
	return len(allMap) + exCnt
}

func main() {
	fmt.Println(minNumBooths([]string{"lxutzb", "lweyedayd", "ducohycnm",
		"h", "pcvcgcykuj", "puqpyzo", "ekg", "cn", "tfrr", "hucfpu"}))
	fmt.Println(minNumBooths([]string{"acd", "bed", "accd", "abcc"}))
	fmt.Println(minNumBooths([]string{"nnwwalkfduiayikfatpvvofudbhz",
		"udgvypgkrxdetwuszjfhhchfdquijvlvllhhopozlsfi",
		"lzhtwacmkhpamgzxrrxbrebonacgibttnmxwed",
		"qgaucsgjfkeyiphnflmxvzanthuag", "skzoslgztfpqumvhk",
		"fxjvcroutluvewubpkckjxxyxgbnjkjmyqhznuljmzxdxrronc",
		"nfuyxmlaxkywhndvdlvnlyoqzftwkbzpbjvovyuxmrzekvu",
		"oirizespwtaly", "mufekyfbjgsfwldjnvbuieudbgckkjmvrfqjd",
		"uudjbvrvwxjnycmnifrxbrqyarcqterliwlaowxsbabnzxtgqs",
		"jzrozdfhinyqogmzsqor", "bgxherthqtnqeaocg", "ovriqvdpgptcghvoxxhjcmoiilmwvit",
		"yzmuoygipuadrh", "sdchvwwbmnfueufoemfx",
		"huohmmroyzsquumzlfiitfdaofxcxskifmvxflxrtx",
		"kwfgwjyrwrkxruurhmwoozocoldzasmzsnacfrzityljnexs",
		"tizyhjntkag", "bffappdycekwrntjdhneelbxlkwxrvsuxpzoy",
		"afb", "uyymihfxowhu", "dzuthgplbbfwerzwhjhicymtv", "weksdsb",
		"pagjaqtpipxvmfzrz", "tbkvgzsqpe", "cyvzguabix", "xptebllpqppj",
		"lzutjwyxwgwevwxfmyujykvpiyrgvovpur", "hzcydgsnvzmoxxkkrodoqkopfripdqixkfysqio",
		"hfpyqstnkogaagqdkjiyjgw", "jmblbeeg", "kihstpwjajljuzusjim", "humvqvhshzbkedqqtqorbc",
		"igddnsxeigqdwiwzqesax", "lfwk", "kflmbpbzctcsepsthxpjlbymlajduqvddc",
		"dasxhxlrpjhippbgbhzmaaspugousstkteybwqtreoievms",
		"huntgxdfjghuulovdkpeozymbdsbqhoqxmflm", "mgatuxmpntgppwcltkjfrksdbfjxjxedslgckxzyuqqqqnfn",
		"x", "lmcqsofrjjmelscohm", "lonxdxjygctuaipnzgysjtcycumerai",
		"gjlphazcjwyntkwqyqimbwgpewtilinswpyyzjhqzpqmckolpa", "bjtvuprqfdrohjgzjggmbwyzfmwfii",
		"fagyepxyazciiignofz", "pkwsuo", "pytarfzlrururqhgbumzgvzmhxwgxclgbnbc",
		"thzkpeadzrjuomqqsjeuawygjnylbla", "xvkioalcnxhpbupjopftygxtaallhxanc",
		"yswglduvavyflqwlwlbcxiikqhxjwicxplplqfbfrjayxeqstn",
	}))
}
