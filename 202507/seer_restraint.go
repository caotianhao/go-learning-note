package main

import (
	"go-cth/model"
)

func GetSeerElementRestraint1v1(atk, def model.SeerElem) model.SeerElemKezhiRes {
	if atk == model.SeerElemOrdinary || def == model.SeerElemOrdinary {
		return model.SeerElemKezhiResNormal
	}

	for _, v := range model.SeerElemKe[atk] {
		if v == def {
			return model.SeerElemKezhiResKe
		}
	}

	for _, v := range model.SeerElemKed[atk] {
		if v == def {
			return model.SeerElemKezhiResKed
		}
	}

	if zero, ok := model.SeerElemInvalid[atk]; ok && zero == def {
		return model.SeerElemKezhiResZero
	}

	return model.SeerElemKezhiResNormal
}

func GetSeerElementRestraint1v2(atk, def1, def2 model.SeerElem) model.SeerElemKezhiRes {
	ab := GetSeerElementRestraint1v1(atk, def1)
	ac := GetSeerElementRestraint1v1(atk, def2)

	return singleCalculate(ab, ac)
}

func GetSeerElementRestraint2v1(atk1, atk2, def model.SeerElem) model.SeerElemKezhiRes {
	ac := GetSeerElementRestraint1v1(atk1, def)
	bc := GetSeerElementRestraint1v1(atk2, def)

	return singleCalculate(ac, bc)
}

func GetSeerElementRestraint2v2(atk1, atk2, def1, def2 model.SeerElem) model.SeerElemKezhiRes {
	abc := GetSeerElementRestraint2v1(atk1, atk2, def1)
	abd := GetSeerElementRestraint2v1(atk1, atk2, def2)

	res := (abc + abd) / 2
	if res > model.SeerElemKezhiResMax {
		return model.SeerElemKezhiResMax
	}
	return res
}

func singleCalculate(res1, res2 model.SeerElemKezhiRes) model.SeerElemKezhiRes {
	div := model.SeerElemKezhiResKe
	if res1 == model.SeerElemKezhiResKe && res2 == model.SeerElemKezhiResKe {
		div = model.SeerElemKezhiResNormal
	}
	if res1 == model.SeerElemKezhiResZero || res2 == model.SeerElemKezhiResZero {
		div = model.SeerElemKezhiResMax
	}

	return (res1 + res2) / div
}

func main() {

}
