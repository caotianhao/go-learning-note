package main

import (
	"fmt"
	"gogogo/model"
)

func GetSeerElementRestraint1v1(atk, def int) float64 {
	if atk == model.SeerElementTypeOrdinary || def == model.SeerElementTypeOrdinary {
		return model.SeerElementRestraintResultNormal
	}

	for _, v := range model.SeerElementTypeRestraint[atk] {
		if v == def {
			return model.SeerElementRestraintResultRestraint
		}
	}

	for _, v := range model.SeerElementTypeRestrainted[atk] {
		if v == def {
			return model.SeerElementRestraintResultRestrainted
		}
	}

	if zero, ok := model.SeerElementTypeZero[atk]; ok && zero == def {
		return model.SeerElementRestraintResultZero
	}

	return model.SeerElementRestraintResultNormal
}

func GetSeerElementRestraint1v2(atk, def1, def2 int) float64 {
	ab := GetSeerElementRestraint1v1(atk, def1)
	ac := GetSeerElementRestraint1v1(atk, def2)

	return singleCalculate(ab, ac)
}

func GetSeerElementRestraint2v1(atk1, atk2, def int) float64 {
	ac := GetSeerElementRestraint1v1(atk1, def)
	bc := GetSeerElementRestraint1v1(atk2, def)

	return singleCalculate(ac, bc)
}

func GetSeerElementRestraint2v2(atk1, atk2, def1, def2 int) float64 {
	abc := GetSeerElementRestraint2v1(atk1, atk2, def1)
	abd := GetSeerElementRestraint2v1(atk1, atk2, def2)

	res := (abc + abd) / 2
	if res > 4 {
		return 4
	}
	return res
}

func singleCalculate(res1, res2 float64) float64 {
	var div float64 = 2
	if res1 == model.SeerElementRestraintResultRestraint && res2 == model.SeerElementRestraintResultRestraint {
		div = 1
	}
	if res1 == model.SeerElementRestraintResultZero || res2 == model.SeerElementRestraintResultZero {
		div = 4
	}

	return (res1 + res2) / div
}

func main() {
	fmt.Println(GetSeerElementRestraint2v2(
		model.SeerElementTypeHoly, model.SeerElementTypeNature,
		model.SeerElementTypeElectricity, model.SeerElementTypeFire,
	) == 4)
	fmt.Println(GetSeerElementRestraint2v2(
		model.SeerElementTypeElectricity, model.SeerElementTypeFire,
		model.SeerElementTypeHoly, model.SeerElementTypeNature,
	) == 0.5)
	fmt.Println(GetSeerElementRestraint2v2(
		model.SeerElementTypeChaos, model.SeerElementTypeDragon,
		model.SeerElementTypeIce, model.SeerElementTypeSuper,
	) == 2.5)
	fmt.Println(GetSeerElementRestraint2v2(
		model.SeerElementTypeIce, model.SeerElementTypeSuper,
		model.SeerElementTypeChaos, model.SeerElementTypeDragon,
	) == 0.875)
	fmt.Println(GetSeerElementRestraint1v2(
		model.SeerElementTypeDimension,
		model.SeerElementTypeFly, model.SeerElementTypeSuper,
	) == 4)
	fmt.Println(GetSeerElementRestraint2v1(
		model.SeerElementTypeFly, model.SeerElementTypeSuper,
		model.SeerElementTypeDimension,
	) == 0.75)
	fmt.Println(GetSeerElementRestraint2v2(
		model.SeerElementTypeHoly, model.SeerElementTypeFly,
		model.SeerElementTypeGround, model.SeerElementTypeMystery,
	) == 0.875)
	fmt.Println(GetSeerElementRestraint2v2(
		model.SeerElementTypeGround, model.SeerElementTypeMystery,
		model.SeerElementTypeHoly, model.SeerElementTypeFly,
	) == 0.75)
	fmt.Println(GetSeerElementRestraint1v2(
		model.SeerElementTypeGrass,
		model.SeerElementTypeLight, model.SeerElementTypeMystery,
	) == 1.5)
	fmt.Println(GetSeerElementRestraint2v1(
		model.SeerElementTypeLight, model.SeerElementTypeMystery,
		model.SeerElementTypeGrass,
	) == 0.25)
}
