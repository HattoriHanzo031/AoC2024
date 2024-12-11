package main

import (
	"common"
	"fmt"
	"math"
	"strings"
)

func main() {
	scaner, close := common.FileScaner("input.txt")
	defer close()
	scaner.Scan()
	stones := common.ToInts(strings.Fields(scaner.Text()))

	total := 0
	for _, stone := range stones {
		total += blink(stone, 75)
	}
	fmt.Println("Results:", total)
}

type key struct {
	stone, blinks int
}

var cache = make(map[key]int)

func blink(stone, blinks int) int {
	if blinks == 0 {
		return 1
	}

	if v, ok := cache[key{stone, blinks}]; ok {
		return v
	}

	numStones := 0
	if stone == 0 {
		numStones = blink(1, blinks-1)
	} else if digits := numDigits(stone); digits%2 == 0 {
		mask := int(math.Pow(10, float64(digits/2)))
		numStones = blink(stone/mask, blinks-1) + blink(stone%mask, blinks-1)
	} else {
		numStones = blink(stone*2024, blinks-1)
	}

	cache[key{stone, blinks}] = numStones
	return numStones
}

func numDigits(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(n))) + 1
}
