package main

import "testing"

func TestNumDigits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Zero", 0, 1},
		{"Single digit", 5, 1},
		{"Two digits", 42, 2},
		{"Three digits", 123, 3},
		{"Four digits", 2024, 4},
		{"Five digits", 10000, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDigits(tt.n); got != tt.want {
				t.Errorf("numDigits(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
