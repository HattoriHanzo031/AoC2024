package main

import (
	"common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	scaner := common.FileIter("input.txt")

	tcrP1 := 0
	tcrP2 := 0
	resP2 := 0
	for i, line := range scaner {
		num, nums, _ := strings.Cut(line, ": ")
		testValue := common.Must(strconv.Atoi(num))
		operands := common.ToInts(strings.Fields(nums))

		if testP2(operands, testValue) {
			resP2 += testValue
		}

		fmt.Println("-----\nLine:", i, len(operands), operands, testValue)
		//start := time.Now()
		maxOperations := (1 << (len(operands) - 1)) - 1
		//fmt.Println("operands:", operands)
		for i := 0; i <= maxOperations; i++ {
			combined := combine(i, operands)
			fmt.Println("Combined:", combined)
			//fmt.Printf("%016b\n", i)
			if testOperations(combined, testValue) {
				fmt.Println("Found:", testValue)
				if i == 0 {
					tcrP1 += testValue
				}
				tcrP2 += testValue
				break
			}
		}
		//fmt.Println("Time:", time.Since(start))
	}
	fmt.Println("Results:", tcrP1, tcrP2)
	fmt.Println("Results:", resP2)
}

func testP2(operands []int, testValue int) bool {
	for p := range common.Permutations(len(operands)-1, []func(a, b int) int{add, mul, combineInts}) {
		acc := operands[0]
		for i, v := range operands[1:] {
			acc = p[i](acc, v)
			fmt.Println("acc", acc)
		}
		fmt.Println("------")
		if acc == testValue {
			return true
		}
	}
	return false
}

func testOperations(operands []int, testValue int) bool {
	maxOperations := (1 << (len(operands) - 1)) - 1
	for i := 0; i <= maxOperations; i++ {
		res := calculate(i, operands)
		//fmt.Println(i, operands, res)
		if res == testValue {
			return true
		}
	}
	return false
}

func calculate(operations int, operands []int) int {
	//fmt.Printf("%b\n", operations)
	acculumator := operands[0]
	//fmt.Print(operands[0])
	for i := 1; i < len(operands); i++ {
		if (operations>>(i-1))&1 == 1 {
			acculumator *= operands[i]
			//fmt.Print("*", operands[i])
		} else {
			acculumator += operands[i]
			//fmt.Print("+", operands[i])
		}
	}
	//fmt.Println("=", acculumator)

	return acculumator
}

func combine(operations int, operands []int) []int {
	//fmt.Printf("%b\n", operations)
	out := []int{operands[0]}
	//fmt.Print(operands[0])
	for i := 1; i < len(operands); i++ {
		if (operations>>(i-1))&1 == 1 {
			out[len(out)-1] = combineInts(out[len(out)-1], operands[i])
			//fmt.Print("||", operands[i])
		} else {
			out = append(out, operands[i])
			//fmt.Print(" ", operands[i])
		}
	}
	//fmt.Println("=", out)

	return out
}

func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}

func combineInts(a, b int) int {
	if b == 0 {
		return a * 10
	}

	numDigits := 0
	reversed := 0
	for b > 0 {
		reversed = reversed*10 + b%10
		b /= 10
		numDigits++
	}

	for range numDigits {
		a = a*10 + reversed%10
		reversed /= 10
	}
	return a
}
