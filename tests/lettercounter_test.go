package tests

import (
	"testing"

	"github.com/Indbrony/lettercounter"
)

func TestNumberOfVocalsAndConsonants(t *testing.T) {
	//Testing with empty input
	vocals, consonants := lettercounter.NumberOfVocalsAndConsonants("")
	if vocals != 0 || consonants != 0 {
		t.Errorf("NumberOfVocalsAndConsonants(\"\") fail, expexting vocals = %d, consonants = %d, got vocals = %d, consonants = %d", 2, 2, vocals, consonants)
	}

	//Testing with input "osama"
	vocals, consonants = lettercounter.NumberOfVocalsAndConsonants("osama")
	if vocals != 2 || consonants != 2 {
		t.Errorf("NumberOfVocalsAndConsonants(\"osama\") fail, expexting vocals = %d, consonants = %d, got vocals = %d, consonants = %d", 2, 2, vocals, consonants)
	}

	//Testing with input "omama"
	vocals, consonants = lettercounter.NumberOfVocalsAndConsonants("omama")
	if vocals != 2 || consonants != 1 {
		t.Errorf("NumberOfVocalsAndConsonants(\"osama\") fail, expexting vocals = %d, consonants = %d, got vocals = %d, consonants = %d", 2, 2, vocals, consonants)
	}

	//Testing with input "OmAmAK" for case-sensitivity
	vocals, consonants = lettercounter.NumberOfVocalsAndConsonants("OmAmAK")
	if vocals != 2 || consonants != 2 {
		t.Errorf("NumberOfVocalsAndConsonants(\"osama\") fail, expexting vocals = %d, consonants = %d, got vocals = %d, consonants = %d", 2, 2, vocals, consonants)
	}
}
