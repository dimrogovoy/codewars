package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	reverseStr := make([]string, 0)
	splitStr := make([]string, 0)
	splitStr = strings.Split(str, " ")
	for _, i := range splitStr {
		var reverse string
		for _, j := range i {
			reverse = string(j) + reverse
		}
		reverseStr = append(reverseStr, reverse)
	}
	return strings.Join(reverseStr, " ")
}

func main() {
	fmt.Println(ReverseWords("The quick brown fox jumps over the lazy dog."))
}
