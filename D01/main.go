package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"common"
)

func main() {
	var l, r []int
	rSimilarity := make(map[int]int, 0)

	scanner, close := common.FileScaner("input.txt")
	defer close()
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		l = append(l, common.Must(strconv.Atoi(fields[0])))
		rn := common.Must(strconv.Atoi(fields[1]))
		r = append(r, rn)
		rSimilarity[rn]++
	}
	slices.Sort(l)
	slices.Sort(r)

	result := 0
	for i := 0; i < len(l); i++ {
		result += int(math.Abs(float64(l[i] - r[i])))
	}
	fmt.Println("Result P1:", result)

	result = 0
	for i := 0; i < len(l); i++ {
		result += l[i] * rSimilarity[l[i]]
	}
	fmt.Println("Result P2:", result)
}
