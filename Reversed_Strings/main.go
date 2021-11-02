package main

import "fmt"

func Solution(word string) string {
	var res string
	for _, i := range word {
		res = string(i) + res
	}
	return res
}

func main() {
	fmt.Println(Solution("world"))
}
