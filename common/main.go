package common

import (
	"bufio"
	"iter"
	"os"
	"slices"

	"golang.org/x/exp/constraints"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func FileScaner(path string) (*bufio.Scanner, func()) {
	file := Must(os.Open(path))
	scanner := bufio.NewScanner(file)
	return scanner, func() {
		Must(struct{}{}, scanner.Err())
		file.Close()
	}
}

func FileIter(path string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		file := Must(os.Open(path))
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for i := 0; scanner.Scan(); i++ {
			if !yield(i, scanner.Text()) {
				break
			}
		}
		Must(struct{}{}, scanner.Err())
	}
}

func Abs[T constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func DeleteClone(slice []int, i int) []int {
	return slices.Delete(slices.Clone(slice), i, i+1)
}
