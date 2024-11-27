package main

import (
	"fmt"
	"math"
)

const (
	exp_three_free_youli = 10000 * 3                // 3次广告免费游历经验
	exp_five_double_ad   = 10000 * 5 * 2            // 5次双倍广告经验
	exp_jianghu          = 4000*4 + 5000*3 + 6000*3 // 江湖任务获取经验
	exp_j4               = 3650                     // 深渊4每局经验
	exp_main22           = 7750                     // 精英22关每局经验

	hp_everyday  = 5 * 24 // 每日自然恢复体力
	hp_ad_double = -25    // 双倍广告消耗体力
	hp_jianghu   = 10     // 江湖在线时长送体力
	hp_friend    = 40     // 好友送体力
	hp_ad        = 20     // 广告送体力
	hp_activity  = 5      // 各种活动平均每天多一次体力

	all_hp = (hp_everyday +
		hp_ad_double +
		hp_jianghu +
		hp_friend +
		hp_ad +
		hp_activity) / 5
	all_exp = exp_three_free_youli +
		exp_five_double_ad +
		exp_jianghu +
		exp_j4*(all_hp-5) +
		exp_main22*5

	tie_min       = 2300
	tie_cui_22    = -19200*5 - 21600*6 - 24000*6 //达到 22 阶 6 极彩还需要付出的铁
	tie_fenjie_22 = 620928                       // 分解 22 阶 6 极彩装备获得的铁
)

// exp[i] 代表从 i 级到 i+1 级需要的经验
// 当前版本等级上限为 379
var exp = [380]int{}

// cui[i] 代表一件 i 阶装备从普通金色到极彩需要的铁
// 等级上限 379，装备上限为 38
var cui = [39]int{}

func fillExp() {
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

func fillCui() {
	cui[20] = 120800
	cui[21] = 132880
	cui[22] = 144960
	cui[23] = 157040
	cui[24] = 169120
	cui[25] = 181200
	cui[26] = 193280
	cui[27] = 205360
	cui[28] = 217440
	cui[29] = 229520
	cui[30] = 241600
	cui[31] = 253680
	cui[32] = 265760
	cui[33] = 277840
	cui[34] = 289920
	cui[35] = 302000
	cui[36] = cui[35] + 12080   // 推测
	cui[37] = cui[35] + 12080*2 // 推测
	cui[38] = cui[35] + 12080*3 // 推测
}

func main() {
	fillExp()
	fillCui()

	for grade := 27; grade <= 38; grade++ {
		fmt.Printf("换 %d 阶装备，", grade)

		totalNeedExp := 0                  // 总共所需经验
		needGrade := (grade - 1) * 10      // 需要达到的等级
		tie_max := grade * 100             // 到达等级前每日通过分解获得的铁
		tie_avg := (tie_min + tie_max) / 2 // 简单计算平均每日获得的铁

		for lv := 210; lv <= needGrade; lv++ {
			totalNeedExp += exp[lv]
		}

		needDays := totalNeedExp / all_exp
		fmt.Printf("需要 %d 天等级够，", needDays)

		all_tie := needDays*tie_avg + tie_fenjie_22 + tie_cui_22

		fmt.Printf("此时有 %d 铁，", all_tie)
		fmt.Printf("离需要的 %d 铁还差 %d\n", cui[grade]*6, cui[grade]*6-all_tie)
	}
}
