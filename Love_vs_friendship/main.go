package main

import "fmt"

func WordsToMarks(s string) int {
	var total int32
	for _, i2 := range s {
		total += i2 - 'a' + 1
	}
	return int(total)
}

func main() {
	fmt.Println(WordsToMarks("attitude"))
}
