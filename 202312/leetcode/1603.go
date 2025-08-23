package main

import "fmt"

type ParkingSystem struct {
	bigCar    int
	mediumCar int
	smallCar  int
}

func Constructor1603(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		bigCar:    big,
		mediumCar: medium,
		smallCar:  small,
	}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if ps.bigCar > 0 {
			ps.bigCar--
			return true
		} else {
			return false
		}
	case 2:
		if ps.mediumCar > 0 {
			ps.mediumCar--
			return true
		} else {
			return false
		}
	case 3:
		if ps.smallCar > 0 {
			ps.smallCar--
			return true
		} else {
			return false
		}
	}
	return true
}

func main() {
	ps := Constructor1603(1, 1, 0)
	fmt.Println(ps.AddCar(1))
	fmt.Println(ps.AddCar(2))
	fmt.Println(ps.AddCar(3))
	fmt.Println(ps.AddCar(1))
}
