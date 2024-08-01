package calculator

type Calculator struct {
}

func (c Calculator) Add(i, j int) int {
	return i + j
}

func (c Calculator) Subtract(i, j int) int {
	return i - j
}
