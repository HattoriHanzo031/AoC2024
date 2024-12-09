package main

import (
	"slices"
)

type file struct {
	free            bool
	id, start, size int
}

func (f file) end() int {
	return f.start + f.size
}

func (f file) checksum() int {
	return f.size * (f.start + f.end() - 1) / 2 * f.id
}

func solveP2(files []int) int {
	sections := make([]file, 0, len(files))
	start := 0
	for i, b := range files {
		sections = append(sections, file{
			free:  i%2 == 1,
			id:    i / 2,
			start: start,
			size:  b,
		})
		start += b
	}

	printSections(sections)
	checksum := 0
	for i := range slices.Backward(sections) {
		sectionToMove := &sections[i]
		if sectionToMove.free {
			continue
		}
		if i == 0 {
			checksum += sectionToMove.checksum()
			break
		}
		found := false
		for j := 0; j < i; j++ {
			var freeSpace *file
			if freeSpace = &sections[j]; !sections[j].free || freeSpace.size < sectionToMove.size {
				continue
			}

			checksum += file{
				free:  false,
				id:    sectionToMove.id,
				start: freeSpace.start,
				size:  sectionToMove.size,
			}.checksum()

			freeSpace.size -= sectionToMove.size
			freeSpace.start += sectionToMove.size

			if sections[i-1].free {
				sections[i-1].size += sectionToMove.size
				sectionToMove.size = 0
			} else {
				sectionToMove.free = true
			}
			found = true
			break
		}
		if !found {
			checksum += sectionToMove.checksum()
		}
	}
	return checksum
}

func printSections(sections []file) {
	for _, s := range sections {
		for i := 0; i < s.size; i++ {
			if s.free {
				print(".")
			} else {
				print(s.id)
			}
		}
	}
	println()
}
