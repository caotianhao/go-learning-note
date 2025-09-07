package main

import (
	"fmt"
	"time"
)

const (
	expSucceed      = 16274
	expFailNormal   = 15119
	expFailHard     = 15119
	expHell6        = 6450
	expUpgradePerLv = 1500000

	lvMax = 459
	lvNow = 366

	hpAdNormalDouble = -25
	hpAdHardDouble   = -25
	hpFriend         = 40
	hpLogin60min     = 10
	hpLeader         = 5
	hpAd             = 40
	hpHell6          = -35

	ironBreakWeapon = 83328
	ironBreakShoe   = 111552
	ironBreak27     = 143752
	ironNow         = 173502

	equipLv   = 100
	equipBase = 35

	cui35   = 302000
	cuiDiff = 12080
)

const (
	exp0to9         = 10 * expSucceed
	exp9to24        = (15 * 60 / 45) * expSucceed
	expJianghu      = 4000*4 + 5000*3 + 6000*3
	expDoubleNormal = -hpAdNormalDouble / 5 * expFailNormal * 2
	expDoubleHard   = -hpAdHardDouble / 5 * expFailHard * 2
	expFillHell6    = -hpHell6 / 5 * expHell6
	expOther        = (hpLeader + hpAd + hpLogin60min + hpFriend +
		hpAdHardDouble + hpAdNormalDouble + hpHell6) / 5 * expSucceed

	expNatural = exp0to9 + exp9to24
	expDouble  = expDoubleNormal + expDoubleHard
	expEx      = expJianghu + expOther
	expPerDay  = expNatural + expDouble + expEx + expFillHell6

	ironBreak    = ironBreakShoe + ironBreakWeapon
	ironBreakAll = ironBreak + ironBreak27*4
)

var cuiFromGold2MaxSingle = make([]int, equipLv)

func initCuiSlice() {
	cuiFromGold2MaxSingle[equipBase] = cui35
	for i := equipBase + 1; i < equipLv; i++ {
		cuiFromGold2MaxSingle[i] += (i-equipBase)*cuiDiff + cuiFromGold2MaxSingle[equipBase]
	}
}

func wsSame() {
	startDate := time.Now()

	for days := 1; ; days++ {
		curLv := lvNow + (days*expPerDay)/expUpgradePerLv
		if curLv > lvMax {
			break
		}
		tmp := curLv/10 + 1
		if tmp > equipLv-1 {
			break
		}

		ironPerDay := (tmp + 1) * 100
		curIron := ironNow + ironBreak + days*ironPerDay
		needIron := 2 * cuiFromGold2MaxSingle[tmp]
		dateStr := startDate.AddDate(0, 0, days).Format("06-01-02")

		fmt.Printf("%s, lv.%d, need %d iron\n", dateStr, curLv, needIron-curIron)
	}
}

func wsUnSame() {
	startDate := time.Now()

	for days := 1; ; days++ {
		curLv := lvNow + (days*expPerDay)/expUpgradePerLv
		if curLv > lvMax {
			break
		}
		tmp := curLv/10 + 1
		if tmp > equipLv-1 {
			break
		}

		ironPerDay := (tmp + 1) * 100
		curIronW := ironNow + ironBreakWeapon + days*ironPerDay
		curIronS := ironNow + ironBreakShoe + days*ironPerDay
		needIron := cuiFromGold2MaxSingle[tmp]
		dateStr := startDate.AddDate(0, 0, days).Format("06-01-02")

		fmt.Printf("%s, lv.%d, W %d iron, S %d iron\n",
			dateStr, curLv, needIron-curIronW, needIron-curIronS)
	}
}

func changeAll() {
	startDate := time.Now()

	for days := 1; ; days++ {
		curLv := lvNow + (days*expPerDay)/expUpgradePerLv
		if curLv > lvMax {
			break
		}
		tmp := curLv/10 + 1
		if tmp > equipLv-1 {
			break
		}

		ironPerDay := (tmp + 1) * 100
		curIron := ironNow + ironBreakAll + days*ironPerDay
		needIron := 6 * cuiFromGold2MaxSingle[tmp]
		dateStr := startDate.AddDate(0, 0, days).Format("06-01-02")

		fmt.Printf("%s, lv.%d, change all need %d iron\n", dateStr, curLv, needIron-curIron)
	}
}

func main() {
	initCuiSlice()
	wsSame()
	fmt.Println("********************************************************")
	wsUnSame()
	fmt.Println("********************************************************")
	changeAll()
}
