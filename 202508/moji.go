package main

import (
	"fmt"
	"time"
)

const (
	expSucceed      = 15702
	expFail         = 15408
	expUpgradePerLv = 1300000

	lvMax = 449
	lvNow = 344

	hpAdNormalDouble = -25
	hpAdHardDouble   = -25
	hpFriend         = 40
	hpLogin60min     = 10
	hpLeader         = 5
	hpAd             = 40

	ironBreak = 190000
	ironNow   = 77852
)

const (
	exp0to9    = 10 * expSucceed
	exp9to24   = (15 * 60 / 45) * expSucceed
	expJianghu = 4000*4 + 5000*3 + 6000*3
	expDouble  = (hpAdHardDouble + hpAdNormalDouble) / -5 * expFail
	expOther   = (hpLeader + hpAd + hpLogin60min + hpFriend +
		hpAdHardDouble + hpAdNormalDouble) / 5 * expSucceed

	expPerDay = exp0to9 + exp9to24 + expJianghu + expDouble + expOther
)

func main() {
	cuiFromGold2MaxSingle := make([]int, 51)
	cuiFromGold2MaxSingle[35] = 54000 + 113000 + 135000
	for i := 36; i < 51; i++ {
		cuiFromGold2MaxSingle[i] += (i-35)*12080 + cuiFromGold2MaxSingle[35]
	}

	startDate := time.Now()

	for days := 1; ; days++ {
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
