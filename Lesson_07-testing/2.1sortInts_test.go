package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func sampleData() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}
	return data
}

func BenchmarkInts(b *testing.B) {
	data := sampleData()
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

// goos: darwin
// goarch: amd64
// pkg: thinknetica_golang_core/Lesson_07-testing
// cpu: Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz
// BenchmarkInts-8   	      24	  42959036 ns/op	 1736614 B/op	       2 allocs/op
// PASS
// ok  	thinknetica_golang_core/Lesson_07-testing	2.951s
