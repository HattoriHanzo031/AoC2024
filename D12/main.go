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

func main() {
	scanner := common.FileIter("test_input_2.txt")
	topo := make(map[coord]rune, 0)
	for y, line := range scanner {
		for x, r := range line {
			topo[coord{x, y}] = r
		}
	}

	totalPrice := 0
	totalPriceP2 := 0
	visited := make(map[coord]bool, 0)
	fenceLenP2 := 0

	for cur := range topo {
		if visited[cur] {
			continue
		}
		fenceLen := 0
		plotSize := 0
		fenceLenP2 = 0

		//fmt.Println("cur:", cur)
		fences := make(map[coord]ofType, 0)
		for c, t := range findFence(topo, cur) {
			switch t {
			case plot:
				visited[c] = true
				plotSize++
			case fenceLeft, fenceRight, fenceUp, fenceDown:
				fenceLen++
				fmt.Print("fence:", c, t.String())
				if !hasConnectedFence(fences, t, c) {
					fenceLenP2++
					fmt.Println(" COUNTED")
				} else {
					fmt.Println()
				}
				fences[c] = t
			}
		}
		testTest := 0
		for c, t := range fences {
			if !hasConnectedFence2(fences, t, c) {
				testTest++
			}
		}
		fmt.Println(string(topo[cur]))
		fmt.Println("plotSize:", plotSize)
		fmt.Println("testTest:", testTest)
		//fmt.Println("fenceLen:", fenceLen)
		fmt.Println("fenceLenP2:", fenceLenP2)
		totalPrice += plotSize * fenceLen
		totalPriceP2 += plotSize * testTest
	}
	fmt.Println("Result:", totalPrice)
	fmt.Println("Result P2:", totalPriceP2)
}

func findFence(topo map[coord]rune, start coord) iter.Seq2[coord, ofType] {
	type visitFrom struct {
		visit coord
		from  ofType
	}

	plotRune := topo[start]
	toVisit := []visitFrom{{start, -1}}
	visited := make(map[coord]bool, 0)
	return func(yield func(coord, ofType) bool) {
		for len(toVisit) > 0 {
			newToVisit := make([]visitFrom, 0)
			for _, current := range toVisit {
				if topo[current.visit] != plotRune {
					if !yield(current.visit, current.from) {
						return
					}
					continue
				}

				if visited[current.visit] {
					continue
				}
				visited[current.visit] = true

				if !yield(current.visit, plot) {
					return
				}

				for _, dir := range []coord{up, down, left, right} {
					newToVisit = append(newToVisit, visitFrom{current.visit.add(dir), dirToFence(dir)})
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

func hasConnectedFence(fences map[coord]ofType, cur ofType, pos coord) bool {
	switch cur {
	case fenceUp:
		return fences[pos.add(left)] == fenceUp || fences[pos.add(right)] == fenceUp
	case fenceDown:
		return fences[pos.add(left)] == fenceDown || fences[pos.add(right)] == fenceDown
	case fenceLeft:
		return fences[pos.add(up)] == fenceLeft || fences[pos.add(down)] == fenceLeft
	case fenceRight:
		return fences[pos.add(up)] == fenceRight || fences[pos.add(down)] == fenceRight
	}
	return false
}

func hasConnectedFence2(fences map[coord]ofType, cur ofType, pos coord) bool {
	switch cur {
	case fenceUp:
		return fences[pos.add(left)] == fenceUp
	case fenceDown:
		return fences[pos.add(left)] == fenceDown
	case fenceLeft:
		return fences[pos.add(up)] == fenceLeft
	case fenceRight:
		return fences[pos.add(up)] == fenceRight
	}
	return false
}
