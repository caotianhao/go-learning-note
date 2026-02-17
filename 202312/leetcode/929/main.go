package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	l, emailArr := len(emails), make([][]string, 0)
	for i := 0; i < l; i++ {
		atInd, plusInd, local, host := 0, 0, "", ""
		for j := 0; j < len(emails[i]); j++ {
			if emails[i][j] == '+' {
				plusInd = j
				break
			}
		}
		for j := 0; j < len(emails[i]); j++ {
			if emails[i][j] == '@' {
				atInd = j
			}
		}
		if plusInd != 0 {
			local = emails[i][:plusInd]
		} else {
			local = emails[i][:atInd]
		}
		host = emails[i][atInd+1:]
		local = strings.Replace(local, ".", "", -1)
		emailArr = append(emailArr, []string{local, host})
	}
	emailMap := map[string]int{}
	for i := range emailArr {
		emailMap[emailArr[i][0]+"@"+emailArr[i][1]]++
	}
	return len(emailMap)
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	fmt.Println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}))
	fmt.Println(numUniqueEmails([]string{"fg.r.u.uzj+o.pw@kziczvh.com", "r.cyo.g+d.h+b.ja@tgsg.z.com",
		"fg.r.u.uzj+o.f.d@kziczvh.com", "r.cyo.g+ng.r.iq@tgsg.z.com", "fg.r.u.uzj+lp.k@kziczvh.com",
		"r.cyo.g+n.h.e+n.g@tgsg.z.com", "fg.r.u.uzj+k+p.j@kziczvh.com", "fg.r.u.uzj+w.y+b@kziczvh.com",
		"r.cyo.g+x+d.c+f.t@tgsg.z.com", "r.cyo.g+x+t.y.l.i@tgsg.z.com", "r.cyo.g+brxxi@tgsg.z.com",
		"r.cyo.g+z+dr.k.u@tgsg.z.com", "r.cyo.g+d+l.c.n+g@tgsg.z.com", "fg.r.u.uzj+vq.o@kziczvh.com",
		"fg.r.u.uzj+uzq@kziczvh.com", "fg.r.u.uzj+mvz@kziczvh.com", "fg.r.u.uzj+taj@kziczvh.com",
		"fg.r.u.uzj+fek@kziczvh.com"}))
}
