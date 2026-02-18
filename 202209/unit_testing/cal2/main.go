package cal2

func productCal(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}
