package pprof

func CacheWater_CaoTianhao(height []int) (rain int) {
	l, r, prefixMax, suffixMax := 0, len(height)-1, 0, 0
	for l < r {
		prefixMax = max(prefixMax, height[l])
		suffixMax = max(suffixMax, height[r])
		if prefixMax < suffixMax {
			rain += prefixMax - height[l]
			l++
		} else {
			rain += suffixMax - height[r]
			r--
		}
	}
	return
}
