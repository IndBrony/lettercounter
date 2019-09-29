package lettercounter

import (
	"regexp"
	"strings"
)

//NumberOfVocalsAndConsonants is used to count the number of vocals and consonants in a string
//recomended using the concurrent version for more effectivity
func NumberOfVocalsAndConsonants(text string) (int, int) {
	text = strings.ToLower(text)
	return numberOfLetters(text, `[aiueo]`), numberOfLetters(text, `[bcdfghjklmnpqrstvwxyz]`)
}

//NumberOfVocalsAndConsonantsConcurrent is used to count the number of vocals and consonants in a string concurrently
func NumberOfVocalsAndConsonantsConcurrent(text string) (int, int) {
	text = strings.ToLower(text)
	c1, c2 := make(chan int), make(chan int)
	go func() {
		c1 <- numberOfLetters(text, `[aiueo]`)
	}()
	go func() {
		c2 <- numberOfLetters(text, `[bcdfghjklmnpqrstvwxyz]`)
	}()

	return <-c1, <-c2
}

func numberOfLetters(text string, regex string) int {
	encountered := make(map[string]bool)
	compiledRegex := regexp.MustCompile(regex)
	found := compiledRegex.FindAll([]byte(text), -1)
	for _, letter := range found {
		encountered[string(letter)] = true
	}
	return len(encountered)
}
