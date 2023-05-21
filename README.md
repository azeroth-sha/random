# random
this is a random libary

## 基准测试
```text
goos: windows
goarch: amd64
pkg: github.com/azeroth-sha/random
cpu: Intel(R) Xeon(R) CPU E3-1231 v3 @ 3.40GHz
BenchmarkInt32-8                         8632359               410.5 ns/op            48 B/op          2 allocs/op
BenchmarkInt32-8                         8375444               494.4 ns/op            48 B/op          2 allocs/op
BenchmarkInt32-8                         7797870               480.3 ns/op            48 B/op          2 allocs/op
BenchmarkInt32-8                         8743336               399.9 ns/op            48 B/op          2 allocs/op
BenchmarkInt32-8                         8873595               400.4 ns/op            48 B/op          2 allocs/op
BenchmarkInt32WithPool-8                47969299                71.42 ns/op            0 B/op          0 allocs/op
BenchmarkInt32WithPool-8                48939175                71.64 ns/op            0 B/op          0 allocs/op
BenchmarkInt32WithPool-8                50669826                69.38 ns/op            0 B/op          0 allocs/op
BenchmarkInt32WithPool-8                53325590                68.87 ns/op            0 B/op          0 allocs/op
BenchmarkInt32WithPool-8                51553034                70.14 ns/op            0 B/op          0 allocs/op
BenchmarkInt32WithParallel-8            25159042               133.6 ns/op            48 B/op          2 allocs/op
BenchmarkInt32WithParallel-8            25129222               134.8 ns/op            48 B/op          2 allocs/op
BenchmarkInt32WithParallel-8            28208655               135.4 ns/op            48 B/op          2 allocs/op
BenchmarkInt32WithParallel-8            26260273               136.6 ns/op            48 B/op          2 allocs/op
BenchmarkInt32WithParallel-8            24430795               140.4 ns/op            48 B/op          2 allocs/op
BenchmarkInt32WithPoolAndParallel-8     147377596               26.79 ns/op            0 B/op          0 allocs/op
BenchmarkInt32WithPoolAndParallel-8     141114391               27.97 ns/op            0 B/op          0 allocs/op
BenchmarkInt32WithPoolAndParallel-8     164581246               23.79 ns/op            0 B/op          0 allocs/op
BenchmarkInt32WithPoolAndParallel-8     158806753               24.32 ns/op            0 B/op          0 allocs/op
BenchmarkInt32WithPoolAndParallel-8     165639486               25.62 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/azeroth-sha/random   88.789s
```
