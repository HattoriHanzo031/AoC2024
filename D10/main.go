package main

import (
	"common"
	"fmt"
)

type coord struct {
	x, y int
}

var (
	up    coord = coord{0, 1}
	down  coord = coord{0, -1}
	left  coord = coord{-1, 0}
	right coord = coord{1, 0}
)

func (c coord) add(d coord) coord {
	return coord{c.x + d.x, c.y + d.y}
}

func (c coord) findNext(topo map[coord]int) []coord {
	res := make([]coord, 0)
	cur := topo[c]
	for _, dir := range []coord{up, down, left, right} {
		next := c.add(dir)
		if topo[next] == cur+1 {
			res = append(res, next)
		}
	}
	return res
}

func main() {
	scanner := common.FileIter("input.txt")
	maxX, maxY := 0, 0
	trailHeads := make([]coord, 0)
	topo := make(map[coord]int, 0)
	for y, line := range scanner {
		for x, r := range line {
			if r == '0' {
				trailHeads = append(trailHeads, coord{x, y})
			}
			topo[coord{x, y}] = int(r - '0')
			maxX = max(maxX, x)
			maxY = max(maxY, y)
		}
	}

	scoreP1 := 0
	scoreP2 := 0
	for _, head := range trailHeads {
		trailEnds := make(map[coord]int, 0)
		trail(topo, trailEnds, head)

		scoreP1 += len(trailEnds)
		for _, v := range trailEnds {
			scoreP2 += v
		}
	}

	fmt.Println("Score P1:", scoreP1)
	fmt.Println("Score P2:", scoreP2)
}

func trail(topo map[coord]int, trailEnds map[coord]int, head coord) {
	if topo[head] == 9 {
		trailEnds[head]++
		return
	}
	for _, next := range head.findNext(topo) {
		trail(topo, trailEnds, next)
	}
}
