package main

import (
	"common"
	"fmt"
	"iter"
)

type ofType int

const (
	plot ofType = iota
	fenceLeft
	fenceRight
	fenceUp
	fenceDown
)

func (ot ofType) String() string {
	switch ot {
	case plot:
		return "plot"
	case fenceLeft:
		return "fenceLeft"
	case fenceRight:
		return "fenceRight"
	case fenceUp:
		return "fenceUp"
	case fenceDown:
		return "fenceDown"
	}
	return "unknown"
}

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

type ct struct {
	c coord
	t ofType
}

func main() {
	scanner := common.FileIter("input.txt")
	plotMap := make(map[coord]rune, 0)
	for y, line := range scanner {
		for x, r := range line {
			plotMap[coord{x, y}] = r
		}
	}

	totalPriceP1 := 0
	totalPriceP2 := 0
	visited := make(map[coord]bool, 0)
	for cur := range plotMap {
		if visited[cur] {
			continue
		}

		fenceLenP1 := 0
		plotSize := 0
		fences := make(map[ct]bool, 0)
		for c, t := range findFence(plotMap, cur) {
			switch t {
			case plot:
				visited[c] = true
				plotSize++
			case fenceLeft, fenceRight, fenceUp, fenceDown:
				fenceLenP1++
				fences[ct{c, t}] = true
			}
		}

		fenceLenP2 := 0
		for cct := range fences {
			if !hasConnectedFence(fences, cct.t, cct.c) {
				fenceLenP2++
			}
		}

		//fmt.Println(string(plotMap[cur]), ": plot:", plotSize, "fenceLenP1:", fenceLenP1, "fenceLenP2:", fenceLenP2)
		totalPriceP1 += plotSize * fenceLenP1
		totalPriceP2 += plotSize * fenceLenP2
	}
	fmt.Println("Price P1:", totalPriceP1)
	fmt.Println("Price P2:", totalPriceP2)
}

func findFence(topo map[coord]rune, start coord) iter.Seq2[coord, ofType] {
	plotRune := topo[start]
	toVisit := []coord{start}
	visited := make(map[coord]bool, 0)
	return func(yield func(coord, ofType) bool) {
		for len(toVisit) > 0 {
			newToVisit := make([]coord, 0)
			for _, current := range toVisit {
				if visited[current] {
					continue
				}
				visited[current] = true

				if !yield(current, plot) {
					return
				}
				for _, dir := range []coord{up, down, left, right} {
					next := current.add(dir)
					if topo[next] == plotRune {
						newToVisit = append(newToVisit, next)
					} else {
						if !yield(current, dirToFence(dir)) {
							return
						}
					}
				}
			}
			toVisit = newToVisit
		}
	}
}

func dirToFence(dir coord) ofType {
	switch dir {
	case up:
		return fenceUp
	case down:
		return fenceDown
	case left:
		return fenceLeft
	case right:
		return fenceRight
	}
	return -1
}

func hasConnectedFence(fences map[ct]bool, cur ofType, pos coord) bool {
	switch cur {
	case fenceUp:
		return fences[ct{pos.add(left), fenceUp}]
	case fenceDown:
		return fences[ct{pos.add(left), fenceDown}]
	case fenceLeft:
		return fences[ct{pos.add(up), fenceLeft}]
	case fenceRight:
		return fences[ct{pos.add(up), fenceRight}]
	}
	return false
}
