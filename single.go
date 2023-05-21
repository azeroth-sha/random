package random

import (
	"bytes"
	"io"
	"sync"
)

var bufPool = &sync.Pool{New: newBuf}

func newBuf() interface{} {
	return bytes.NewBuffer(make([]byte, 8, 8))
}

func randFromReader() int64 {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)
	if _, err := io.CopyN(buf, reader, 8); err != nil {
		panic(err)
	}
	var num int64
	var b byte
	for i := 0; i < 64; i += 8 {
		b, _ = buf.ReadByte()
		num += int64(b) << i
	}
	return num
}
