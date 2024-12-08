package main

import (
	"common"
	"fmt"
)

type coord struct {
	x, y int
}

func (c coord) add(d coord) coord {
	return coord{c.x + d.x, c.y + d.y}
}

func (c coord) sub(d coord) coord {
	return coord{c.x - d.x, c.y - d.y}
}

func (c coord) inBounds(maxX, maxY int) bool {
	return c.x >= 0 && c.x <= maxX && c.y >= 0 && c.y <= maxY
}

func (c coord) antinodes(b coord) (coord, coord) {
	diff := c.sub(b)
	return c.add(diff), b.sub(diff)
}

func main() {

	scanner := common.FileIter("test_input.txt")

	antinode := make(map[coord]bool)
	allStations := make(map[rune][]coord, 0)
	maxX, maxY := 0, 0
	for y, line := range scanner {
		for x, r := range line {
			if r != '.' {
				allStations[r] = append(allStations[r], coord{x, y})
			}
			maxX = max(maxX, x)
			maxY = max(maxY, y)
		}
	}

	for _, stations := range allStations {
		for i, station := range stations {
			for j := i + 1; j < len(stations); j++ {
				first, second := station.antinodes(stations[j])
				if first.inBounds(maxX, maxY) {
					antinode[first] = true
				}
				if second.inBounds(maxX, maxY) {
					antinode[second] = true
				}
			}
		}
	}
	fmt.Println("Results:", len(antinode))
}
