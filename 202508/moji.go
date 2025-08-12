package main

import (
	"fmt"
	"time"
)

const (
	expSucceed      = 15408
	expFail         = 14338
	expUpgradePerLv = 1300000

	lvMax = 449
	lvNow = 344

	hp_ad_normal_double = -25
	hp_ad_hard_double   = -25
	hp_friend           = 40
	hp_login_60min      = 10
	hp_leader           = 5
	hp_ad               = 40

	ironBreak = 190000
	ironNow   = 77852
)

const (
	lvTarget  = lvMax / 10 * 10
	expTarget = expUpgradePerLv * (lvTarget - lvNow)

	exp0to9    = 10 * expSucceed
	exp9to24   = (15 * 60 / 45) * expSucceed
	expJianghu = 4000*4 + 5000*3 + 6000*3
	expDouble  = (hp_ad_hard_double + hp_ad_normal_double) / -5 * expFail
	expOther   = (hp_leader + hp_ad + hp_login_60min + hp_friend +
		hp_ad_hard_double + hp_ad_normal_double) / 5 * expSucceed

	expPerDay = exp0to9 + exp9to24 + expJianghu + expDouble + expOther
)

func main() {
	cuiFromGold2MaxSingle := make([]int, 51)
	cuiFromGold2MaxSingle[35] = 54000 + 113000 + 135000
	for i := 36; i < 51; i++ {
		cuiFromGold2MaxSingle[i] += (i-35)*12080 + cuiFromGold2MaxSingle[35]
	}

	startDate := time.Now()

	for days := 1; days < 365; days++ {
		curLv := lvNow + (days*expPerDay)/expUpgradePerLv
		if curLv > lvMax {
			break
		}
		tmp := curLv/10 + 1
		if tmp > 50 {
			break
		}

		ironPerDay := (tmp + 1) * 100
		curIron := ironNow + ironBreak + days*ironPerDay
		needIron := 2 * cuiFromGold2MaxSingle[tmp]
		dateStr := startDate.AddDate(0, 0, days).Format("2006-01-02")

		fmt.Printf("%s, lv.%d, need %d iron\n", dateStr, curLv, needIron-curIron)
	}
}
