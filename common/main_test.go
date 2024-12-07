package common

import (
	"reflect"
	"slices"
	"testing"
)

func TestPermutations(t *testing.T) {
	type args struct {
		numChars int
		charSet  []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"2,2", args{2, []int{1, 2}}, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}},
		{"2,3", args{2, []int{1, 2, 3}}, [][]int{{1, 1}, {2, 1}, {3, 1}, {1, 2}, {2, 2}, {3, 2}, {1, 3}, {2, 3}, {3, 3}}},
		{"3,2", args{3, []int{0, 1}}, [][]int{{0, 0, 0}, {1, 0, 0}, {0, 1, 0}, {1, 1, 0}, {0, 0, 1}, {1, 0, 1}, {0, 1, 1}, {1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permutations(tt.args.numChars, tt.args.charSet)
			out := make([][]int, 0)
			i := 0
			for s := range got {
				out = append(out, slices.Clone(s))
				if i > 100 {
					break
				}
				i++
			}
			if !reflect.DeepEqual(out, tt.want) {
				t.Errorf("Permutations() = %v, want %v", out, tt.want)
			}
		})
	}
}
