package pprof

import "testing"

func BenchmarkCacheWater_CaoTianhao(b *testing.B) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	for i := 0; i < b.N; i++ {
		CacheWater_CaoTianhao(height)
	}
}

//func BenchmarkCacheWater_CaoTianhao(b *testing.B) {
//	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
//
//	// Create a sync.Pool for height slices.
//	var heightPool = sync.Pool{
//		New: func() interface{} {
//			return make([]int, len(height))
//		},
//	}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		// Get a height slice from the pool.
//		h := heightPool.Get().([]int)
//		copy(h, height) // Copy the original height slice to the pooled slice.
//
//		// Use the pooled slice for the function.
//		CacheWater_CaoTianhao(h)
//
//		// Put the height slice back to the pool.
//		heightPool.Put(h)
//	}
//}

//func BenchmarkCacheWater_CaoTianhao(b *testing.B) {
//	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
//
//	// Create a sync.Pool for height slices.
//	var heightPool = sync.Pool{
//		New: func() interface{} {
//			return make([]int, len(height))
//		},
//	}
//
//	b.ResetTimer()
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			// Get a height slice from the pool.
//			h := heightPool.Get().([]int)
//			copy(h, height) // Copy the original height slice to the pooled slice.
//
//			// Use the pooled slice for the function.
//			CacheWater_CaoTianhao(h)
//
//			// Put the height slice back to the pool.
//			heightPool.Put(h)
//		}
//	})
//}
