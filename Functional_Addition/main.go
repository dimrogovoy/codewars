package main

import "fmt"

func Add(a int) func(int) int {
	return func(i int) int {
		return a + i
	}
}

func main() {
	f := Add(5)
	fmt.Println(f(3))
	fmt.Println(f(2))
}
