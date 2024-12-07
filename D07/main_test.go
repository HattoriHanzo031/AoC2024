package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_combineInts(t *testing.T) {
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
			if got := combineInts(tt.first, tt.second); got != tt.want {
				t.Errorf("combineInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combine(t *testing.T) {
	type args struct {
		operations int
		operands   []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args{0b0, []int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{args{0b1, []int{1, 2, 3, 4}}, []int{12, 3, 4}},
		{args{0b10, []int{1, 2, 3, 4}}, []int{1, 23, 4}},
		{args{0b11, []int{1, 2, 3, 4}}, []int{123, 4}},
		{args{0b100, []int{1, 2, 3, 4}}, []int{1, 2, 34}},
		{args{0b101, []int{1, 2, 3, 4}}, []int{12, 34}},
		{args{0b110, []int{1, 2, 3, 4}}, []int{1, 234}},
		{args{0b111, []int{1, 2, 3, 4}}, []int{1234}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := combine(tt.args.operations, tt.args.operands); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		operations int
		operands   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0b000", args{0b000, []int{1, 2, 3, 4}}, 1 + 2 + 3 + 4},
		{"0b001", args{0b001, []int{1, 2, 3, 4}}, 1*2 + 3 + 4},
		{"0b010", args{0b010, []int{1, 2, 3, 4}}, (1+2)*3 + 4},
		{"0b011", args{0b011, []int{1, 2, 3, 4}}, 1*2*3 + 4},
		{"0b100", args{0b100, []int{1, 2, 3, 4}}, (1 + 2 + 3) * 4},
		{"0b101", args{0b101, []int{1, 2, 3, 4}}, (1*2 + 3) * 4},
		{"0b110", args{0b110, []int{1, 2, 3, 4}}, (1 + 2) * 3 * 4},
		{"0b111", args{0b111, []int{1, 2, 3, 4}}, 1 * 2 * 3 * 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.operations, tt.args.operands); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
