package main

import (
	"fmt"
	"math"
)

// 体力
const (
	hp_everyday  = 5 * 24 // 每日自然恢复体力
	hp_ad_double = -25    // 双倍广告消耗体力
	hp_jianghu   = 10     // 江湖在线时长送体力
	hp_friend    = 40     // 好友送体力
	hp_ad        = 20     // 广告送体力
	hp_activity  = 5      // 各种活动平均每天多一次体力
	hp_leader    = 5      // 门派俸禄

	all_hp = (hp_everyday + hp_ad_double + hp_jianghu +
		hp_friend + hp_ad + hp_activity + hp_leader) / 5
)

// 经验
const (
	exp_three_free  = 11275 * 3                // 3 次广告免费游历经验
	exp_five_double = 10600 * 5 * 2            // 5 次双倍广告经验
	exp_jianghu     = 4000*4 + 5000*3 + 6000*3 // 江湖任务获取经验
	exp_j4          = 3650                     // 深渊 4 每局经验
	exp_j5          = 4950                     // 深渊 5 每局经验
	exp_j6          = 6450                     // 深渊 6 每局经验
	exp_j7          = 7950                     // 深渊 7 每局经验
	exp_m19         = 6025                     // 精英 19 每局经验
	exp_m20         = 6550                     // 精英 20 每局经验
	exp_m22         = 7550                     // 精英 22 每局经验
	exp_m25         = 9050

	exp_base = exp_three_free + exp_five_double + exp_jianghu
	all_exp  = exp_base + exp_three_free*(all_hp-6) + exp_m25*6
)

// 铁
const (
	tie_activity  = 150
	tie_min       = 2600 + tie_activity
	tie_now       = 130000
	tie_fenjie_22 = (66192+56112)*2 + 83328 + 94752 + 16800

	forever27 = 205360 * 4
	diff      = 120800 - 108720
)

var (
	exp  = [380]int{} // exp[i] 代表从 i 级到 i+1 级需要的经验
	iron = [39]int{}  // iron[i] 代表一件 i 阶装备从普通金色到极彩需要的铁
	hole = make(chan any, 1)
)

func fill() {
	// solve unused variable hole
	hole <- []any{all_exp, exp_j5, exp_j6, exp_m20, exp_j7, exp_m19, exp_m22, exp_j4}

	// fill iron
	for i := 27; i < 39; i++ {
		iron[i] = (205360+diff*(i-27))*2 + forever27
	}

	// fill exp
	exp[137] = 695000
	for i := 138; i <= 200; i++ {
		exp[i] = exp[i-1] + 5000
	}
	for i := 201; i <= 220; i++ {
		exp[i] = exp[i-1] + 6000
	}
	for i := 221; i <= 238; i++ {
		exp[i] = exp[i-1] + 7000
	}
	exp[239] = 1263000 // 推测
	exp[240] = 1270000 // 推测
	exp[241] = 1300000 // 推测
	exp[242] = 1300000 // 推测
	exp[243] = 1300000 // 推测
	for i := 244; i <= 354; i++ {
		exp[i] = 1300000
	}
	for i := 355; i <= 378; i++ {
		exp[i] = 1300000 // 推测
	}
	exp[379] = math.MaxInt64
}

func main() {
	fill()

	for grade := 27; grade <= 38; grade++ {
		fmt.Printf("换 %d 阶武器和鞋，", grade)

		totalNeedExp := 0                   // 总共所需经验
		needGrade := (grade - 1) * 10       // 需要达到的等级
		tie_max := grade*100 + tie_activity // 到达等级前每日通过分解获得的铁
		tie_avg := (tie_min + tie_max) / 2  // 简单计算平均每日获得的铁

		for lv := 240; lv <= needGrade; lv++ {
			totalNeedExp += exp[lv]
		}

		needDays := totalNeedExp / all_exp
		fmt.Printf("需要 %d 天等级够，", needDays)

		all_tie := needDays*tie_avg + tie_fenjie_22 + tie_now

		fmt.Printf("此时有 %d 铁，", all_tie)
		fmt.Printf("离需要的 %d 铁还差 %d\n", iron[grade], iron[grade]-all_tie)
	}
}
