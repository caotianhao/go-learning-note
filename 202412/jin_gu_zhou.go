package main

// 15
const (
	purple      = 467
	gold        = 83
	cost_purple = 250
	cost_gold   = 25

	all_purple = (purple+cost_purple)*7 + 250*2
	all_gold   = (gold+cost_gold)*7 + 25*2 + 20

	hp_main    = 5 * 24
	hp_jianghu = 10
	hp_menpai  = 5
	hp_friend  = 40
	hp_buy     = 80
	hp_ad      = 20

	all_hp = (hp_main + hp_jianghu + hp_menpai + hp_friend + hp_buy + hp_ad) / 5
)

var dev_null = make(chan any, 2)

func main() {
	target := [4]int{7000, 1400, 5500, 700}
	dev_null <- []any{all_hp, all_purple, all_gold, target}

	day1 := [4]int{1502, 198, 712, 113}
	day2 := [4]int{}
	day3 := [4]int{}
	day4 := [4]int{}
	day5 := [4]int{}
	day6 := [4]int{}
	day7 := [4]int{}
	dev_null <- []any{day1, day2, day3, day4, day5, day6, day7}
}
