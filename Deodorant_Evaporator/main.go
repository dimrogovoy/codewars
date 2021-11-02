package main

import "fmt"

func Evaporator(content float64, evapPerDay int, threshold int) int {
	days := 1
	min := content * float64(threshold) / 100.
	for content*(1.-float64(evapPerDay)/100.) >= min {
		content = content * (1. - float64(evapPerDay)/100.)
		days++
	}
	return days
}

func main() {
	fmt.Println(Evaporator(10, 10, 100))
}
