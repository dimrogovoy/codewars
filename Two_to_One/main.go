package main

import (
	"fmt"
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {

	chars := append(strings.Split(s1, ""), strings.Split(s2, "")...)
	uniqueChars := make(map[string]struct{})

	var result []string
	for _, s := range chars {
		uniqueChars[s] = struct{}{}
	}
	for key := range uniqueChars {
		result = append(result, key)
	}

	sort.Strings(result)
	return strings.Join(result, "")

}

func main() {
	fmt.Println(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
}
