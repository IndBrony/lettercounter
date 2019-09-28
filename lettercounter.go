package lettercounter

import (
	"regexp"
	"strings"
)

//NumberOfVocalsAndConsonants is used to count the number of vocals and consonants in a string
func NumberOfVocalsAndConsonants(text string) (int, int) {
	text = strings.ToLower(text)
	return numberOfLetters(text, `[aiueo]`), numberOfLetters(text, `[bcdfghjklmnpqrstvwxyz]`)
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
