package main

import (
	"common"
	"fmt"
)

type coord struct {
	x, y int
}

var (
	up    coord = coord{0, -1}
	down  coord = coord{0, 1}
	left  coord = coord{-1, 0}
	right coord = coord{1, 0}
)

func (c coord) add(d coord) coord {
	return coord{c.x + d.x, c.y + d.y}
}

func (c coord) turnRight() coord {
	switch c {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	return coord{}
}

func (c coord) inBounds(maxX, maxY int) bool {
	return c.x >= 0 && c.x <= maxX && c.y >= 0 && c.y <= maxY
}

func main() {
	scanner := common.FileIter("input.txt")

	obstacles := make(map[coord]bool, 0)
	start := coord{}
	//path := make([]coord, 0)
	path := make(map[coord]bool, 0)

	maxX, maxY := 0, 0
	direction := up
	for y, line := range scanner {
		for x, r := range line {
			switch r {
			case '#':
				obstacles[coord{x, y}] = true
			case '^':
				start = coord{x, y}
			}
			maxX = max(maxX, x)
			maxY = max(maxY, y)
		}
	}
	fmt.Println("start:", start)
	fmt.Println("obstacles:", obstacles)
	cur := start
	//path = append(path, cur)
	path[cur] = true
	for cur.inBounds(maxX, maxY) {
		next := cur.add(direction)
		if obstacles[next] {
			direction = direction.turnRight()
		} else {
			cur = next
			//path = append(path, cur)
			path[cur] = true
		}
	}

	fmt.Println("Path:", path)
	fmt.Println("Length:", len(path)-1)
}
