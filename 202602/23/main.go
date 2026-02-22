package main

import (
	"fmt"
	"time"
)

type dateStr time.Time

func (d dateStr) String() string {
	return time.Time(d).Format("January 2, 2006")
}

type placeStr string

func (p placeStr) String() string {
	return string(p)
}

func DiaryEntry[D fmt.Stringer, L fmt.Stringer](date D, location L) string {
	return fmt.Sprintf("%s, in %s", date.String(), location.String())
}

func main() {
	kst := time.FixedZone("KST", 9*60*60)
	today := dateStr(time.Date(2026, 2, 23, 0, 0, 0, 0, kst))
	loc := placeStr("Suwon, South Korea")
	fmt.Println(DiaryEntry(today, loc))
	fmt.Println()
	fmt.Println("singing")
	fmt.Println("왜이래 나도 바빠 바빠 바빠")
	fmt.Println("왜 그래 너만 바빠 바빠 바빠......")
}
