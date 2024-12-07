package common

import (
	"bufio"
	"iter"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

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

func DeleteClone[T any](slice []T, i int) []T {
	return slices.Delete(slices.Clone(slice), i, i+1)
}

func ToInts[T ~string | ~[]byte](ss []T) []int {
	ints := make([]int, 0, len(ss))
	for _, s := range ss {
		ints = append(ints, Must(strconv.Atoi(strings.TrimSpace(string(s)))))
	}
	return ints
}

func Permutations[T any](numChars int, charSet []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		perm := make([]T, numChars)

		maxCombinations := math.Pow(float64(len(charSet)), float64(numChars))
		for i := 0; i < int(maxCombinations); i++ {
			cur := i
			for j := 0; j < numChars; j++ {
				perm[j] = charSet[cur%len(charSet)]
				cur /= len(charSet)
			}
			if !yield(perm) {
				break
			}
		}
	}
}
