package main

const (
	purpleNow = 1394
	goldNow   = 246

	purpleAll = 6500
	goldAll   = 1200

	purpleNeed = purpleAll - purpleNow
	goldNeed   = goldAll - goldNow

	purplePinTuan = 4500 + 3000
	goldPinTuan   = 450 + 300
)

var hole = make(chan any, 1<<4)

func main() {
	hole <- purpleNeed
	hole <- goldNeed
	hole <- purplePinTuan
	hole <- goldPinTuan
}
