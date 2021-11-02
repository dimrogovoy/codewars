package main

import "fmt"

func BlackOrWhiteKey(keyPressCount int) string {
	pianoKey1 := [8]string{
		"white", "black", "white", "white", "black", "white", "black", "white",
	}
	pianoKey2 := [12]string{
		"white", "black", "white", "black", "white", "black", "white", "white", "black", "white", "black", "white",
	}
	pianoKey3 := [8]string{
		"white", "black", "white", "black", "white", "black", "white", "white",
	}

	lastKey := (keyPressCount - 1) % 88
	if lastKey < 8 {
		return pianoKey1[lastKey]
	}
	if lastKey < 80 {
		lastKey = (lastKey - 8) % 12
		return pianoKey2[lastKey]
	}
	return pianoKey3[lastKey-80]
}

func main() {
	fmt.Println(BlackOrWhiteKey(88))
}
