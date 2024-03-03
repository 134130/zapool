```shell
go test -bench=.
```
Tested with MacOS M1 Pro / 32GB

```
goos: darwin
goarch: arm64
pkg: github.com/134130/zapool
BenchmarkPool_Get_BuffSize1024_PoolSize1-10             985093951                1.084 ns/op           0 B/op          0 allocs/op
BenchmarkPool_Get_BuffSize1024_PoolSize100-10           1000000000               1.072 ns/op           0 B/op          0 allocs/op
BenchmarkPool_Get_BuffSize1024_PoolSize10000-10         1000000000               1.072 ns/op           0 B/op          0 allocs/op
BenchmarkPool_Get_BuffSize1024_PoolSize1000000-10       1000000000               1.083 ns/op           0 B/op          0 allocs/op
BenchmarkNoPool_Get_BuffSize1024-10                     4807077                  255.4 ns/op        1048 B/op          2 allocs/op
BenchmarkPool_Get_BuffSize4096_PoolSize1-10             1000000000               1.060 ns/op           0 B/op          0 allocs/op
BenchmarkPool_Get_BuffSize4096_PoolSize100-10           1000000000               1.058 ns/op           0 B/op          0 allocs/op
BenchmarkPool_Get_BuffSize4096_PoolSize10000-10         1000000000               1.471 ns/op           0 B/op          0 allocs/op
BenchmarkPool_Get_BuffSize4096_PoolSize1000000-10       1000000000               1.069 ns/op           0 B/op          0 allocs/op
BenchmarkNoPool_Get_BuffSize4096-10                     1468178                  829.7 ns/op        4120 B/op          2 allocs/op
```
