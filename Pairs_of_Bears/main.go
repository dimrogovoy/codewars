package main

import "fmt"

func CheckPairs(nbBear int, bearList string) (bearCouple string, hasEnoughCouple bool) {
	defer func() {
		if !hasEnoughCouple {
			fmt.Printf("-> %d %q\n<- %q %t\n", nbBear, bearList, bearCouple, hasEnoughCouple)
		}
	}()

	var countBears int
	for i := 1; i < len(bearList); i++ {
		if (bearList[i] == 'B' && bearList[i-1] == '8') || (bearList[i] == '8' && bearList[i-1] == 'B') {
			countBears += 2 // <- тут должно быть общее кол-во медведей, не пар
			bearCouple += string(bearList[i-1]) + string(bearList[i])
			i += 1
		}
	}
	if countBears >= nbBear {
		hasEnoughCouple = true
	} else {
		hasEnoughCouple = false
	}
	return bearCouple, hasEnoughCouple
}

func main() {
	fmt.Println(CheckPairs(6, "B8B8f8B8gB"))
}
