package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func DigPow(n, p int) int {
	digits := strings.Split(strconv.Itoa(n), "")
	var total float64
	for i := 0; i < len(digits); i++ {
		s, _ := strconv.Atoi(digits[i])
		total += math.Pow(float64(s), float64(p))
		p++
	}
	if int(total)%n == 0 {
		return int(total) / n
	} else {
		return -1
	}
}

func main() {
	fmt.Println(DigPow(46288, 3))
}
