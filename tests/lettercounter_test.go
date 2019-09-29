package tests

import (
	"testing"

	"github.com/Indbrony/lettercounter"
)

func basicNumberOfVocalsAndConsonantsTest(t *testing.T, input string, expectedVocals, expectedConsonants int) {
	vocals, consonants := lettercounter.NumberOfVocalsAndConsonants(input)
	if vocals != expectedVocals || consonants != expectedConsonants {
		t.Errorf("NumberOfVocalsAndConsonants(%v) fail, expexting vocals = %d, consonants = %d, got vocals = %d, consonants = %d", input, expectedVocals, expectedConsonants, vocals, consonants)
	}
}
func basicNumberOfVocalsAndConsonantsConcurrentTest(t *testing.T, input string, expectedVocals, expectedConsonants int) {
	vocals, consonants := lettercounter.NumberOfVocalsAndConsonantsConcurrent(input)
	if vocals != expectedVocals || consonants != expectedConsonants {
		t.Errorf("NumberOfVocalsAndConsonants(%v) fail, expexting vocals = %d, consonants = %d, got vocals = %d, consonants = %d", input, expectedVocals, expectedConsonants, vocals, consonants)
	}
}

func TestNumberOfVocalsAndConsonants(t *testing.T) {
	//Testing with empty input
	basicNumberOfVocalsAndConsonantsTest(t, "", 0, 0)
	//Testing with input "osama"
	basicNumberOfVocalsAndConsonantsTest(t, "osama", 2, 2)
	//Testing with input "omama"
	basicNumberOfVocalsAndConsonantsTest(t, "omama", 2, 1)
	//Testing with input "OmAmAK" for case-sensitivity
	basicNumberOfVocalsAndConsonantsTest(t, "OmAmAK", 2, 2)
	//Testing with input "Fuad Mustamirrul Ishlah" to make sure it can handle real world situation
	basicNumberOfVocalsAndConsonantsTest(t, "Fuad Mustamirrul Ishlah", 3, 8)
}

func TestNumberOfVocalsAndConsonantsConcurrent(t *testing.T) {
	//Testing with empty input
	basicNumberOfVocalsAndConsonantsConcurrentTest(t, "", 0, 0)
	//Testing with input "osama"
	basicNumberOfVocalsAndConsonantsConcurrentTest(t, "osama", 2, 2)
	//Testing with input "omama"
	basicNumberOfVocalsAndConsonantsConcurrentTest(t, "omama", 2, 1)
	//Testing with input "OmAmAK" for case-sensitivity
	basicNumberOfVocalsAndConsonantsConcurrentTest(t, "OmAmAK", 2, 2)
	//Testing with input "Fuad Mustamirrul Ishlah" to make sure it can handle real world situation
	basicNumberOfVocalsAndConsonantsConcurrentTest(t, "Fuad Mustamirrul Ishlah", 3, 8)
}

func BenchmarkNumberOfVocalsAndConsonants(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lettercounter.NumberOfVocalsAndConsonants("NumberOfVocalsAndConsonantsConcurrent is used to count the number of vocals and consonants in a string concurrently")
	}
}

func BenchmarkNumberOfVocalsAndConsonantsConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lettercounter.NumberOfVocalsAndConsonantsConcurrent("NumberOfVocalsAndConsonantsConcurrent is used to count the number of vocals and consonants in a string concurrently")
	}
}
