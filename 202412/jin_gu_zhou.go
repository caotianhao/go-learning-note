package main

import "fmt"

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

func main() {
	// solve no use
	fmt.Println(all_hp, all_purple, all_gold)
}
