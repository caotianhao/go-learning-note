// Mersenne Twister（MT）是一种高质量的伪随机数生成器
// 它具有长周期（2^19937 - 1），统计特性较好，并且适用于许多应用领域
package main

import (
	"fmt"
	"time"
)

const (
	w = 32
	n = 624
	m = 397
	//r = 31
	a = 0x9908B0DF
	u = 11
	d = 0xFFFFFFFF
	s = 7
	b = 0x9D2C5680
	t = 15
	c = 0xEFC60000
	l = 18
	f = 1812433253
)

type MersenneTwister struct {
	mt    [n]uint32
	index int
}

func NewMersenneTwister(seed uint32) *MersenneTwister {
	mt := MersenneTwister{index: n}
	mt.mt[0] = seed
	for i := 1; i < n; i++ {
		mt.mt[i] = f*(mt.mt[i-1]^(mt.mt[i-1]>>(w-2))) + uint32(i)
	}
	return &mt
}

func (mt *MersenneTwister) Twist() {
	for i := 0; i < n; i++ {
		x := (mt.mt[i] & upperMask) + (mt.mt[(i+1)%n] & lowerMask)
		xa := x >> 1
		if x%2 != 0 {
			xa = xa ^ a
		}
		mt.mt[i] = mt.mt[(i+m)%n] ^ xa
	}
	mt.index = 0
}

func (mt *MersenneTwister) generateNumbers() {
	for i := 0; i < n-m; i++ {
		y := (mt.mt[i] & upperMask) | (mt.mt[i+1] & lowerMask)
		mt.mt[i] = mt.mt[i+m] ^ (y >> 1) ^ mag01[y&1]
	}
	for i := n - m; i < n-1; i++ {
		y := (mt.mt[i] & upperMask) | (mt.mt[i+1] & lowerMask)
		mt.mt[i] = mt.mt[i+(m-n)] ^ (y >> 1) ^ mag01[y&1]
	}
	y := (mt.mt[n-1] & upperMask) | (mt.mt[0] & lowerMask)
	mt.mt[n-1] = mt.mt[m-1] ^ (y >> 1) ^ mag01[y&1]
	mt.index = 0
}

func (mt *MersenneTwister) Next() uint32 {
	if mt.index >= n {
		mt.generateNumbers()
	}

	y := mt.mt[mt.index]
	y ^= (y >> u) & d
	y ^= (y << s) & b
	y ^= (y << t) & c
	y ^= y >> l

	mt.index++
	return y
}

const (
	upperMask = 0x80000000
	lowerMask = 0x7FFFFFFF
)

var mag01 = [2]uint32{0, a}

func main() {
	// 使用时间作为种子初始化
	mt := NewMersenneTwister(uint32(time.Now().UnixNano()))

	// 生成一些随机数示例
	for i := 0; i < 10; i++ {
		fmt.Println(mt.Next())
	}
}
