package main

import (
	"fmt"
	"strings"
)

func countValidWords(sentence string) (cnt int) {
	s1 := strings.Fields(sentence)
	for i := 0; i < len(s1); i++ {
		if isHaveNum(s1[i]) && isHavePunctuation(s1[i]) && isHaveHyphen(s1[i]) {
			cnt++
			//fmt.Printf("s[%d] = %v\n", i, s1[i])
		}
	}
	return
}

func isHaveNum(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return false
		}
	}
	return true
}

func isHavePunctuation(s string) bool {
	l := len(s)
	for i := 0; i < l-1; i++ {
		if s[i] == '!' || s[i] == '.' || s[i] == ',' {
			return false
		}
	}
	if (s[l-1] == '!' || s[l-1] == '.' || s[l-1] == ',') && l != 1 {
		if s[l-2] == '-' {
			return false
		}
	}
	return true
}

func isHaveHyphen(s string) bool {
	l, cnt := len(s), 0
	if s[0] == '-' || s[l-1] == '-' {
		return false
	}
	for i := 1; i < l-1; i++ {
		if s[i] == '-' {
			cnt++
		}
	}
	if cnt != 0 && cnt != 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(countValidWords("    cat and    dog"))                      //3
	fmt.Println(countValidWords("!this  1-s b8d!"))                         //0
	fmt.Println(countValidWords("alice and  bob are playing stone-game10")) //5
	fmt.Println(countValidWords("c cat 1c c1a c1"))                         //2
	fmt.Println(countValidWords("!"))                                       //1
	fmt.Println(countValidWords(" 62   nvtk0wr4f  8 qt3r! w1ph 1l ,e0d 0n 2v 7c.  " +
		"n06huu2n9 s9   ui4 nsr!d7olr  q-, vqdo!btpmtmui.bb83lf g .!v9-lg 2fyoykex u" +
		"y5a 8v whvu8 .y sc5 -0n4 zo pfgju 5u 4 3x,3!wl  fv4   s  aig cf j1 a i  8m5o1 " +
		" !u n!.1tz87d3 .9    n a3  .xb1p9f  b1i a j8s2 cugf l494cx1! hisceovf3 8d93 sg" +
		" 4r.f1z9w   4- cb r97jo hln3s h2 o .  8dx08as7l!mcmc isa49afk i1 fk,s e !1 ln" +
		" rt2vhu 4ks4zq c w  o- 6  5!.n8ten0 6mk 2k2y3e335,yj  h p3 5 -0  5g1c  tr49," +
		" ,qp9 -v p  7p4v110926wwr h x wklq u zo 16. !8  u63n0c l3 yckifu 1cgz t.i   l" +
		"h w xa l,jt   hpi ng-gvtk8 9 j u9qfcd!2  kyu42v dmv.cst6i5fo rxhw4wvp2 1 okc8" +
		"!  z aribcam0  cp-zp,!e x  agj-gb3 !om3934 k vnuo056h g7 t-6j! 8w8fncebuj-lq " +
		"   inzqhw v39,  f e 9. 50 , ru3r  mbuab  6  wz dw79.av2xp . gbmy gc s6pi pra4f" +
		"o9fwq k   j-ppy -3vpf   o k4hy3 -!..5s ,2 k5 j p38dtd   !i   b!fgj,nx qgif ")) //49
}
