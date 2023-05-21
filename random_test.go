package random_test

import (
	"github.com/azeroth-sha/random"
	"testing"
)

// func TestUint64(t *testing.T) {
// 	// random.DisableRandPool()
// 	random.EnableRandPool()
// 	for i := 0; i < 256; i++ {
// 		var count int64
// 		for j := 0; j < 100000; j++ {
// 			// random.DisableRandPool()
// 			// fmt.Printf("----%d---- %d\r\n", i, random.Uint64())
// 			// random.EnableRandPool()
// 			// fmt.Printf("----%d---- %d\r\n", i, random.Uint64())
// 			if random.Uint8() == uint8(i) {
// 				count++
// 			}
// 		}
// 		log.Printf("num: %d; count: %d\r\n", i, count)
// 	}
// }

func BenchmarkInt32(b *testing.B) {
	random.DisableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Int32()
	}
}

func BenchmarkInt32WithPool(b *testing.B) {
	random.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		random.Int32()
	}
}

func BenchmarkInt32WithParallel(b *testing.B) {
	random.DisableRandPool()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			random.Int32()
		}
	})
}

func BenchmarkInt32WithPoolAndParallel(b *testing.B) {
	random.EnableRandPool()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			random.Int32()
		}
	})
}
