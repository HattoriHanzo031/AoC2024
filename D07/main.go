package main

import (
	"common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	scaner := common.FileIter("input.txt")
	totalP1, totalP2 := 0, 0
	for _, line := range scaner {
		num, nums, _ := strings.Cut(line, ": ")
		testValue := common.Must(strconv.Atoi(num))
		operands := common.ToInts(strings.Fields(nums))

		if testPermutations(operands, testValue, []operation{add, mul}) {
			totalP1 += testValue
			totalP2 += testValue
		} else if testPermutations(operands, testValue, []operation{add, mul, combine}) {
			totalP2 += testValue
		}
	}
	fmt.Println("Results:", totalP1, totalP2)
}

func testPermutations(operands []int, testValue int, operations []operation) bool {
	permutations := common.Permutations(len(operands)-1, operations)
	for p := range permutations {
		acc := operands[0]
		for i, v := range operands[1:] {
			acc = p[i](acc, v)
			if acc > testValue {
				break
			}
		}
		if acc == testValue {
			return true
		}
	}
	return false
}

type operation func(a, b int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func combine(a, b int) int {
	return common.Must(strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b)))
}
