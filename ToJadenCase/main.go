package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {
	return strings.Title(strings.ToLower(str))
}

func main() {
	fmt.Println(ToJadenCase("most tReeS and bLUe"))

}
