package main

import (
	"fmt"
	"strings"

	"sync"
)



func freqofLetters(s string, m map[string]int,  wg *sync.WaitGroup) {

	sli := strings.Split(s, "")
	for ch := 0; ch < len(sli); ch++ {
		m[sli[ch]]++
	}
	wg.Done()


}

func printCharacters(m map[string]int){
    asciiNum := 97

    fmt.Println("{")

    for i := 0; i < 26; i++ {
    character := string(asciiNum + i)
    fmt.Printf("  %q: ", character)
    fmt.Println(m[character])
    }

    fmt.Println("}")
}

 /*func f(m map[string]int,setofNames string,wg *sync.WaitGroup){
	for _,ch:= range setofNames{
		m[string(ch)]++
	}
	wg.Done()
}*/

func main() {
	var wg sync.WaitGroup
	m := make(map[string]int)

	setofNames := []string{"balu", "tarun"}



	for i := 0; i < len(setofNames); i++ {
		wg.Add(1)


		go freqofLetters(setofNames[i], m, &wg)
	}
	wg.Wait()

	printCharacters(m)


}
