package main

import (
	"fmt"
	"math"
)

const (
	MAX_STAGE             = 35
	ZB_COUNT              = 6
	NOW_LEVEL             = 202
	NOW_IRON              = 220000 + 250000
	AVG_IRON              = 2707
	HOW_MANY_DAYS_A_LEVEL = 5
)

var cost = [MAX_STAGE + 1]int{
	0,
	3475, 6950, 10425, 13900, 17375,
	20850, 24325, 27800, 31275, 34750,
	41700, 48650, 55600, 62550, 69500,
	83400, 97300, 111200, 125100, 139000,
	152900, 166800, 180700, 194600, 208500,
	222400, 236300, 250200, 264100, 278000,
	291900, 305800, 319700, 333600, 347500,
}

func main() {
	for i := 0; i < len(cost); i++ {
		cost[i] *= ZB_COUNT
	}

	for days := 0; ; days += HOW_MANY_DAYS_A_LEVEL {
		lv := days/HOW_MANY_DAYS_A_LEVEL + NOW_LEVEL
		t := NOW_IRON + days*AVG_IRON

		zb := int(math.Ceil(float64(lv) / 10.0))
		if lv%10 == 0 {
			zb++
		}
		if zb > MAX_STAGE {
			break
		}
		need := cost[zb]

		fmt.Printf("%d days, lv = %d, t = %d, zb = %d, need = %d, want = %d\n",
			days, lv, t, zb, need, need-t)

		if need-t < 0 {
			break
		}
	}
}
