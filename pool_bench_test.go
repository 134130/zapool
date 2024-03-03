package zapool

import (
	"testing"
)

func newByteBuffConstructor(size int) func() any {
	return func() any {
		return make([]byte, size)
	}
}

func benchmarkPool_Get(size int, constructor func() any, b *testing.B) {
	pool := New[any](constructor)
	for i := 0; i < size; i++ {
		pool.Put(constructor())
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			obj := pool.Get()
			pool.Put(obj)
		}
	})
}

func BenchmarkPool_Get_BuffSize1024_PoolSize1(b *testing.B) {
	benchmarkPool_Get(1, newByteBuffConstructor(1024), b)
}
func BenchmarkPool_Get_BuffSize1024_PoolSize100(b *testing.B) {
	benchmarkPool_Get(100, newByteBuffConstructor(1024), b)
}
func BenchmarkPool_Get_BuffSize1024_PoolSize10000(b *testing.B) {
	benchmarkPool_Get(10000, newByteBuffConstructor(1024), b)
}
func BenchmarkPool_Get_BuffSize1024_PoolSize1000000(b *testing.B) {
	benchmarkPool_Get(1000000, newByteBuffConstructor(1024), b)
}

func BenchmarkNoPool_Get_BuffSize1024(b *testing.B) {
	constructor := newByteBuffConstructor(1024)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			obj := constructor()
			_ = obj
		}
	})
}

func BenchmarkPool_Get_BuffSize4096_PoolSize1(b *testing.B) {
	benchmarkPool_Get(1, newByteBuffConstructor(4096), b)
}
func BenchmarkPool_Get_BuffSize4096_PoolSize100(b *testing.B) {
	benchmarkPool_Get(100, newByteBuffConstructor(4096), b)
}
func BenchmarkPool_Get_BuffSize4096_PoolSize10000(b *testing.B) {
	benchmarkPool_Get(10000, newByteBuffConstructor(4096), b)
}
func BenchmarkPool_Get_BuffSize4096_PoolSize1000000(b *testing.B) {
	benchmarkPool_Get(1000000, newByteBuffConstructor(4096), b)
}

func BenchmarkNoPool_Get_BuffSize4096(b *testing.B) {
	constructor := newByteBuffConstructor(4096)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			obj := constructor()
			_ = obj
		}
	})
}
