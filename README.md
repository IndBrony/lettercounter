# Letter Counter
go package to count vocals and consonants in a string

## Getting Started 
1. run ``go get github.com/IndBrony/lettercounter`` to pull this package
2. then install the package using ``go build github.com/IndBrony/lettercounter``
3. you're good to go

## Documentation
- NumberOfVocalsAndConsonants(text string) (int, int)
  - this function will return the number of vocals and consonant in english alphabet from the input
  - note: its recommended to use the concurren function as it is faster with longer input

- NumberOfVocalsAndConsonantsConcurrent(text string) (int, int)
  - essentialy the same as NumberOfVocalsAndConsonants but using concurrency in its algorithm for faster output on longer input
