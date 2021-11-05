package main

import (
	"fmt"
	"strings"
)

func PosAverage(s string) float64 {
	splitStr := strings.Split(s, ", ")
	var count, total int
	for i := 0; i < len(splitStr)-1; i++ {
		for j := i + 1; j < len(splitStr); j++ {
			for l := 0; l < len(splitStr[i]); l++ {
				total++
				if splitStr[i][l] == splitStr[j][l] {
					count++
				}
			}
		}
	}
	return float64(count) * 100 / float64(total)

}

func main() {
	fmt.Printf("%.10f", PosAverage("466960, 069060, 494940, 060069, 060090, 640009, 496464, 606900, 004000, 944096"))
}
