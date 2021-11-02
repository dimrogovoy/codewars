package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FreqSeq(str string, sep string) string {
	b := make(map[int32]int)
	for _, j := range str {
		b[j]++
	}
	var a string
	for _, i := range str {
		a += strconv.Itoa(b[i])
	}
	return strings.Join(strings.Split(a, ""), sep)
}

func main() {
	fmt.Println(FreqSeq("hello world", "-"))
}
