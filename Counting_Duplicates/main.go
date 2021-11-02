package main

import (
	"fmt"
	"sort"
	"strings"
)

func duplicate_count(s1 string) int {
	s1 = strings.ToLower(s1)
	symbols := strings.Split(s1, "")
	sort.Strings(symbols)
	var count int
	var letter string
	for i := 1; i < len(symbols); i++ {
		if letter == symbols[i-1] {
			continue
		}
		if symbols[i] == symbols[i-1] {
			count++
			letter = symbols[i-1]
		}
	}
	return count
}

func main() {
	fmt.Println(duplicate_count("aabbcde"))
}
