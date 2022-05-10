package sort

import (
	"reflect"
	"sort"
	"testing"
)

func Test_Ints(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	got := s
	want := []int{1, 2, 3, 4, 5, 6} // sorted
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %d Want: %d", got, want)
	}
}
