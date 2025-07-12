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

	return (ab + ac) / 2
}

func GetSeerElementRestraint2v1(atk1, atk2, def int) float64 {
	ac := GetSeerElementRestraint1v1(atk1, def)
	bc := GetSeerElementRestraint1v1(atk2, def)

	var div float64 = 2
	if ac == model.SeerElementRestraintResultRestraint && bc == model.SeerElementRestraintResultRestraint {
		div = 4
	}
	if ac == model.SeerElementRestraintResultZero || bc == model.SeerElementRestraintResultZero {
		div = 4
	}

	return (ac + bc) / div
}

func GetSeerElementRestraint2v2(atk1, atk2, def1, def2 int) float64 {
	abc := GetSeerElementRestraint2v1(atk1, atk2, def1)
	abd := GetSeerElementRestraint2v1(atk1, atk2, def2)
	return (abc + abd) / 2
}

func main() {
	result1v1 := GetSeerElementRestraint1v1(model.SeerElementTypeFire, model.SeerElementTypeGrass)
	result1v2 := GetSeerElementRestraint1v2(model.SeerElementTypeGod, model.SeerElementTypeDark, model.SeerElementTypeEvil)
	result2v1 := GetSeerElementRestraint2v1(model.SeerElementTypeFly, model.SeerElementTypeSuper, model.SeerElementTypeWar)
	result2v2 := GetSeerElementRestraint2v2(model.SeerElementTypeMachine, model.SeerElementTypeSuper, model.SeerElementTypeElectricity, model.SeerElementTypeIce)

	fmt.Println(result1v1, result1v2, result2v1, result2v2)
}
