package main

import (
	"common"
	"fmt"
	"strings"
)

func main() {
	scanner, close := common.FileScaner("input.txt")
	defer close()

	pageOrder := make(map[int]map[int]bool, 0)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			break
		}

		order := common.ToInts(strings.Split(scanner.Text(), "|"))
		if _, ok := pageOrder[order[1]]; !ok {
			pageOrder[order[1]] = make(map[int]bool, 0)
		}
		pageOrder[order[1]][order[0]] = true
	}

	totalP1 := 0
	totalP2 := 0
	for scanner.Scan() {
		pages := common.ToInts(strings.Split(scanner.Text(), ","))

		sorted := func() bool {
			for i, page := range pages {
				for j := i + 1; j < len(pages); j++ {
					if pageOrder[page][pages[j]] {
						return false
					}
				}
			}
			return true
		}

		if sorted() {
			totalP1 += pages[len(pages)/2]
		} else {
			for i, page := range pages {
				numBefore := 0
				for j, otherPage := range pages {
					if i == j {
						continue
					}
					if pageOrder[page][otherPage] {
						numBefore++
					}
				}
				if numBefore == len(pages)/2 {
					totalP2 += page
					break
				}
			}
		}
	}
	fmt.Println("Result P1:", totalP1)
	fmt.Println("Result P2:", totalP2)
}
