package main

import "fmt"

func GetCount(str string) (count int) {
	for _, vowel := range str {
		switch vowel {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}

	}

	return count

}

func main() {
	fmt.Println(GetCount("abracadabra"))
}
