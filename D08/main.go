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

func (c coord) antinodes(b coord) []coord {
	diff := c.sub(b)
	return []coord{c.add(diff), b.sub(diff)}
}

func (c coord) harmonicsAntinodes(b coord, maxX, maxY int) []coord {
	diff := c.sub(b)
	antinodes := []coord{c, b}
	for antinode := c.add(diff); antinode.inBounds(maxX, maxY); antinode = antinode.add(diff) {
		antinodes = append(antinodes, antinode)
	}
	for antinode := b.sub(diff); antinode.inBounds(maxX, maxY); antinode = antinode.sub(diff) {
		antinodes = append(antinodes, antinode)
	}
	return antinodes
}

func main() {

	scanner := common.FileIter("input.txt")

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

	antinodes := make(map[coord]bool)
	harmonicsAntinodes := make(map[coord]bool)
	for _, stations := range allStations {
		for i, station := range stations {
			for j := i + 1; j < len(stations); j++ {
				for _, antinode := range station.antinodes(stations[j]) {
					if antinode.inBounds(maxX, maxY) {
						antinodes[antinode] = true
					}
				}
				for _, antinode := range station.harmonicsAntinodes(stations[j], maxX, maxY) {
					harmonicsAntinodes[antinode] = true
				}
			}
		}
	}
	fmt.Println("Results P1:", len(antinodes))
	fmt.Println("Results P2:", len(harmonicsAntinodes))
}
