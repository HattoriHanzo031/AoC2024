package main

import (
	"common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := common.FileIter("input.txt")
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	resP1 := 0
	resP2 := 0
	do := true
	for _, line := range scanner {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			switch match {
			case "do()":
				do = true
			case "don't()":
				do = false
			default:
				mul := calculate(match)
				resP1 += mul
				if do {
					resP2 += mul
				}
			}
		}
	}
	fmt.Println("Result P1:", resP1)
	fmt.Println("Result P2:", resP2)
}

func calculate(s string) int {
	s = strings.TrimPrefix(s, "mul(")
	s = strings.TrimSuffix(s, ")")
	parts := strings.Split(s, ",")
	a := common.Must(strconv.Atoi(parts[0]))
	b := common.Must(strconv.Atoi(parts[1]))
	return a * b
}
