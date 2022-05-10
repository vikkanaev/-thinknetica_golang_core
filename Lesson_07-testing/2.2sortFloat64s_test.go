package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func sampleData64() []float64 {
	rand.Seed(time.Now().UnixNano())
	var data []float64
	for i := 0; i < 1_000_000; i++ {
		n := -100 + rand.Float64()*200
		data = append(data, n)
	}
	return data
}

func BenchmarkFloat64s(b *testing.B) {
	data := sampleData64()
	for i := 0; i < b.N; i++ {
		sort.Float64s(data)
	}
}

// goos: darwin
// goarch: amd64
// pkg: thinknetica_golang_core/Lesson_07-testing
// cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
// BenchmarkFloat64s-8   	      12	  97037032 ns/op	 3473206 B/op	       4 allocs/op
// PASS
// ok  	thinknetica_golang_core/Lesson_07-testing	3.749s
