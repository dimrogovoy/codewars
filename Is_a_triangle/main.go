package main

import (
	"fmt"
	"sort"
)

func IsTriangle(a, b, c int) bool {
	arr := []int{a, b, c}
	sort.Ints(arr)
	return arr[0]+arr[1] > arr[2]
}

func main() {
	fmt.Println(IsTriangle(-1, -2, 0))
}
