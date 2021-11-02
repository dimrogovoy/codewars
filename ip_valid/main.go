package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	a := strings.Split(ip, ".")
	if len(a) != 4 {
		return false
	}
	for _, s := range a {
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		b, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if b < 0 || b > 255 {

			return false
		}

	}
	return true
}

func main() {
	fmt.Println(Is_valid_ip("123.045.067.089"))
}
