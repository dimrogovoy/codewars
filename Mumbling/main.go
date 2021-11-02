package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {

	var result []string
	symbol := strings.Split(strings.ToLower(s), "")
	for i := 0; i < len(symbol); i++ {
		result = append(result, strings.Repeat(symbol[i], i+1))
	}
	return strings.Title(strings.Join(result, "-"))
}
func main() {
	fmt.Println(Accum("ZpglnRxqenU"))
}
