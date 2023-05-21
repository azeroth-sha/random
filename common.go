package random

import (
	"crypto/rand"
	"sync/atomic"
)

var (
	reader      = rand.Reader
	poolSize    = 256
	poolEnable  = int32(0)
	defaultDict = []uint8{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm',
		'n', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'M',
		'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'Y', 'Z',
	}
)

// EnableRandPool 开启随机缓存池
func EnableRandPool() {
	atomic.StoreInt32(&poolEnable, 1)
}

// DisableRandPool 关闭随机缓存池
func DisableRandPool() {
	atomic.StoreInt32(&poolEnable, 0)
}

// StatsRandPool 随机缓存池状态
func StatsRandPool() bool {
	return atomic.LoadInt32(&poolEnable) == 1
}
