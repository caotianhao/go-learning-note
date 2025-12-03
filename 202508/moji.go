package main

import (
	"fmt"
	"time"
)

const (
	expSucceed      = 16875
	expFailNormal   = 16600
	expFailHard     = 16600
	expHell6        = 6450
	expUpgradePerLv = 1500000

	lvMax = 460
	lvNow = 434

	hpAdNormalDouble = -25
	hpAdHardDouble   = -25
	hpFriend         = 40
	hpLogin60min     = 10
	hpLeader         = 5
	hpAd             = 40
	hpHell6          = -40

	ironBreakWeapon = 83328
	ironBreakShoe   = 111552
	ironBreak27     = 143752
	ironNow         = 568302

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

func calIron() {
	startDate := time.Now()

	for days := 1; ; days++ {
		lv := lvNow + (days*expPerDay)/expUpgradePerLv
		if lv > lvMax {
			break
		}
		tmp := lv/10 + 1
		if tmp > equipLv-1 {
			break
		}

		ironPerDay := (tmp + 1) * 100

		ironW := ironNow + ironBreakWeapon + days*ironPerDay
		needIronW := cuiFromGold2MaxSingle[tmp]
		ironS := ironNow + ironBreakShoe + days*ironPerDay
		needIronS := cuiFromGold2MaxSingle[tmp]
		iron2 := ironNow + ironBreak + days*ironPerDay
		needIron2 := 2 * cuiFromGold2MaxSingle[tmp]
		ironAll := ironNow + ironBreakAll + days*ironPerDay
		needIronAll := 6 * cuiFromGold2MaxSingle[tmp]

		date := startDate.AddDate(0, 0, days).Format("06-01-02")

		fmt.Printf("%s, lv.%d, W %d, S %d, 2 %d, 6 %d\n", date, lv, needIronW-ironW,
			needIronS-ironS, needIron2-iron2, needIronAll-ironAll)
	}
}

func main() {
	initCuiSlice()
	calIron()
}
