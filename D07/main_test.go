package main

import (
	"strconv"
	"testing"
)

func Test_combine(t *testing.T) {
	tests := []struct {
		first  int
		second int
		want   int
	}{
		{1, 2, 12},
		{12345, 876, 12345876},
		{14365, 1422352, 143651422352},
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 10},
		{81, 40, 8140},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.first)+" "+strconv.Itoa(tt.second), func(t *testing.T) {
			if got := combine(tt.first, tt.second); got != tt.want {
				t.Errorf("combineInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func combineInts(a, b int) int {
// 	if b == 0 {
// 		return a * 10
// 	}

// 	numDigits := 0
// 	reversed := 0
// 	for b > 0 {
// 		reversed = reversed*10 + b%10
// 		b /= 10
// 		numDigits++
// 	}

// 	for range numDigits {
// 		a = a*10 + reversed%10
// 		reversed /= 10
// 	}
// 	return a
// }

// func combineLog(a, b int) int {
// 	numDigits := int(math.Floor(math.Log10(float64(b))) + 1)
// 	return a*int(math.Pow10(numDigits)) + b
// }
