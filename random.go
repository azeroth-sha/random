package random

func readNum() int64 {
	if !StatsRandPool() {
		return randFromReader()
	}
	return randFromPool()
}

func Int8() int8 {
	return int8(readNum())
}

func Uint8() uint8 {
	return uint8(readNum())
}

func Int16() int16 {
	return int16(readNum())
}

func Uint16() uint16 {
	return uint16(readNum())
}

func Int32() int32 {
	return int32(readNum())
}

func Uint32() uint32 {
	return uint32(readNum())
}

func Int64() int64 {
	return readNum()
}

func Uint64() uint64 {
	return uint64(readNum())
}

func String(size int) string {
	return StringWithDict(size, defaultDict)
}

func StringWithDict(size int, dict []byte) string {
	res := make([]byte, size, size)
	dictLen := uint8(len(dict))
	for i := 0; i < size; i++ {
		res[i] = dict[Uint8()%dictLen]
	}
	return string(res)
}
