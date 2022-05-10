package sort

import (
	"reflect"
	"sort"
	"testing"
)

func Test_Strings(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Тест №1",
			args: args{
				s: []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"},
			},
			want: []string{"Alpha", "Bravo", "Delta", "Go", "Gopher", "Grin"},
		},
		{
			name: "Тест №2",
			args: args{
				s: []string{"2", "1", "b", "A", "a", "0"},
			},
			want: []string{"0", "1", "2", "A", "a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.args.s)
			got := tt.args.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got: %v Want: %v", got, tt.want)
			}
		})
	}
}
