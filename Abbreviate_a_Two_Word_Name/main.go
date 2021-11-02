package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	a := strings.Split(name, " ")
	b := strings.Title(string(a[0][0]))
	c := strings.Title(string(a[1][0]))
	return b + "." + c
}

func main() {
	fmt.Println(AbbrevName("sam harris"))
}
