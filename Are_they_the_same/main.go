package main

import "fmt"

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil || len(array1) != len(array2) {
		return false
	}

	var a, b int
	for len(array1) > 0 {
		a = array1[0]
		for j := 0; j < len(array2); {
			b = array2[j]
			if a*a == b {
				array1 = array1[1:]
				array2 = append(array2[:j], array2[j+1:]...)
				break
			}
			j++
		}
		if a*a != b {
			return false
		}
	}

	return true
}

func main() {
	var a1 = []int{}
	var a2 = []int{}
	fmt.Println(Comp(a1, a2))
}
