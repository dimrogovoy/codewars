package main

import (
	"fmt"
	"sort"
)

func FindOdd(seq []int) int {
	sort.Ints(seq)
	count := 1
	for i := 1; i < len(seq); i++ {
		if seq[i-1] == seq[i] {
			count++
		}
		if seq[i-1] != seq[i] {
			if count % 2 != 0 {
				return seq[i-1]
			}
			count = 1
		}
	}
	return seq[len(seq)-1]
}

func main() {
	seq := []int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}

	fmt.Println(FindOdd(seq))
}
