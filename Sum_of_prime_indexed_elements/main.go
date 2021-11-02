package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Solve(arr []int) int {
	var result int
	for i := 0; i < len(arr); i++ {
		if isPrime(i) == true {
			result += arr[i]
		}
	}
	return result
}

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println(Solve([]int{1, 2, 3, 4}))
}
