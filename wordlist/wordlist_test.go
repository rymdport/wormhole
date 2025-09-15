package wordlist

import "testing"

var globalWords string

func BenchmarkChooseWords(b *testing.B) {
	words := ""

	for b.Loop() {
		words = ChooseWords(10)
	}

	globalWords = words
}
