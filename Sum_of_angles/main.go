package main

import "fmt"

func Angle(n int) int {
	return 180 * (n - 2)
}

func main() {
	fmt.Println(Angle(3))
}
