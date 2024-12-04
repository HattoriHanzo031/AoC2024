package main

import (
	"common"
	"fmt"
)

type coord struct {
	x, y int
}

var (
	up        coord = coord{0, 1}
	down      coord = coord{0, -1}
	left      coord = coord{-1, 0}
	right     coord = coord{1, 0}
	upLeft    coord = coord{-1, 1}
	upRight   coord = coord{1, 1}
	downLeft  coord = coord{-1, -1}
	downRight coord = coord{1, -1}
)

func (c coord) add(d coord) coord {
	return coord{c.x + d.x, c.y + d.y}
}

func main() {
	scanner := common.FileIter("input.txt")

	input := make(map[coord]rune, 0)
	for y, line := range scanner {
		for x, c := range line {
			input[coord{x, y}] = c
		}
	}

	resP1 := 0
	resP2 := 0
	for k, v := range input {
		switch v {
		case 'X':
			resP1 += searchP1(input, k)
		case 'A':
			resP2 += searchP2(input, k)
		}
	}
	fmt.Println("Result P1:", resP1)
	fmt.Println("Result P2:", resP2)
}

func searchP1(input map[coord]rune, start coord) int {
	directions := []coord{up, down, left, right, upLeft, upRight, downLeft, downRight}
	res := 0
	for _, dir := range directions {
		res += searchXmas(input, start, dir)
	}
	return res
}

func searchXmas(input map[coord]rune, start coord, dir coord) int {
	for i, r := range "MAS" {
		if input[start.add(coord{dir.x * (i + 1), dir.y * (i + 1)})] != r {
			return 0
		}
	}
	return 1
}

func searchP2(input map[coord]rune, start coord) int {
	opposite := func(x rune) rune {
		switch x {
		case 'M':
			return 'S'
		case 'S':
			return 'M'
		}
		return ' '
	}

	if input[start.add(upLeft)] == opposite(input[start.add(downRight)]) &&
		input[start.add(upRight)] == opposite(input[start.add(downLeft)]) {
		return 1
	}

	return 0
}
