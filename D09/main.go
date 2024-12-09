package main

import (
	"common"
	"fmt"
	"iter"
	"slices"
)

func main() {
	scaner, close := common.FileScaner("input.txt")
	defer close()
	scaner.Scan()
	line := scaner.Bytes()

	files := make([]int, 0, len(line))
	totalBlocks := 0
	for i, b := range line {
		files = append(files, int(b-'0'))
		if i%2 != 1 {
			totalBlocks += int(b - '0')
		}
	}

	forward := filesBlocks(files)
	back, stop := iter.Pull(skipEmptySpace(filesBlocksBackwards(files)))
	defer stop()

	checksumP1 := 0
	position := 0
	checksumFn := func(id int) {
		checksumP1 += position * id
		position++
	}
	for id, isFile := range forward {
		if position == totalBlocks {
			break
		}
		if !isFile {
			moved, valid := back()
			if !valid {
				break
			}
			checksumFn(moved)
		} else {
			checksumFn(id)
		}
	}
	fmt.Println("Checksum P1:", checksumP1)
	fmt.Println("Checksum P2:", solveP2(files))
}

func filesBlocks(files []int) iter.Seq2[int, bool] {
	return func(yield func(int, bool) bool) {
		for i, num := range files {
			for j := 0; j < num; j++ {
				if !yield(i/2, i%2 != 1) {
					return
				}
			}
		}
	}
}

func filesBlocksBackwards(files []int) iter.Seq2[int, bool] {
	backwards := slices.Backward(files)
	return func(yield func(int, bool) bool) {
		for i, num := range backwards {
			for j := 0; j < num; j++ {
				if !yield(i/2, i%2 != 1) {
					return
				}
			}
		}
	}
}

func skipEmptySpace[K any](ith iter.Seq2[K, bool]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, isFile := range ith {
			if !isFile {
				continue
			}
			if !yield(k) {
				break
			}
		}
	}
}
