package random

import (
	"bytes"
	"io"
	"sync"
)

var (
	randPool = &sync.Pool{New: newReader}
)

type pool struct {
	pos  int
	mu   *sync.Mutex
	pool []byte
}

func (p *pool) Int64() int64 {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.pos+8 > poolSize {
		if _, err := io.ReadFull(reader, p.pool[:]); err != nil {
			panic(err)
		}
		p.pos = 0
	}
	buffer := bytes.NewBuffer(p.pool[p.pos:])
	p.pos += 8
	var num int64
	var b byte
	for i := 0; i < 64; i += 8 {
		b, _ = buffer.ReadByte()
		num += int64(b) << i
	}
	return num
}

func newReader() interface{} {
	return &pool{
		pos:  poolSize,
		mu:   new(sync.Mutex),
		pool: make([]byte, poolSize, poolSize),
	}
}

func randFromPool() int64 {
	rand := randPool.Get().(*pool)
	defer randPool.Put(rand)
	return rand.Int64()
}
