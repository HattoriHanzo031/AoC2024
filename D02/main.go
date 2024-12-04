package main

import (
	"common"
	"strconv"
	"strings"
)

func main() {
	scanner := common.FileIter("input.txt")

	p1, p2 := 0, 0
	for _, line := range scanner {
		report := make([]int, 0)
		for _, field := range strings.Fields(line) {
			report = append(report, common.Must(strconv.Atoi(field)))
		}
		if reportSafe(report) {
			p1++
		}
		p2 += func() int {
			for i := 0; i < len(report); i++ {
				if reportSafe(common.DeleteClone(report, i)) {
					return 1
				}
			}
			return 0
		}()
	}
	println("Results:", p1, p2)
}

func reportSafe(report []int) bool {
	last := report[0]
	sign := report[0] > report[1]
	for _, n := range report[1:] {
		if sign != (last >= n) || common.Abs(last-n) > 3 || common.Abs(last-n) == 0 {
			return false
		}
		last = n
	}
	return true
}

// func reportSafe2(report []int) bool {
// 	sign := report[1] > report[0]
// 	return slices.IsSortedFunc(report, func(a, b int) int {
// 		if sign != (a >= b) || a == b || common.Abs(a-b) > 3 {
// 			return -1
// 		}
// 		return 1
// 	})
// }
