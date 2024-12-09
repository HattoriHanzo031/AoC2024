package main

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	tests := []struct {
		file     file
		expected int
	}{
		{file{true, 1, 1, 2}, 3},
		{file{true, 2, 5, 5}, 70},
		{file{true, 3, 3, 7}, 126},
		{file{true, 4, 0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.file.checksum(); got != tt.expected {
				t.Errorf("file.checksum() = %d, want %d", got, tt.expected)
			}
		})
	}
}
