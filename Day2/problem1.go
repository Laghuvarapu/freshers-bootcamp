package main

import (
	"fmt"
	"strings"
)

func freqofLetters(setofStrings []string, char chan string) {
	for str := 0; str < len(setofStrings); str++ {
		sli := strings.Split(setofStrings[str], "")
		for ch := 0; ch < len(sli); ch++ {
			char <- sli[ch]
		}
	}
	close(char)
}

func main() {
	m := make(map[string]int)

	setofNames := []string{"balu", "tarun"}

	char := make(chan string)

	go freqofLetters(setofNames, char)

	for chars := range char {
		m[chars] += 1
	}

	asciiNum := 97
	fmt.Println("{")

	for i := 0; i < 26; i++ {
		character := string(asciiNum + i)
		fmt.Printf("  %q: ", character)
		fmt.Println(m[character])
	}

	fmt.Println("}")

}
