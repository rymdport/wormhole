package wordlist

import "testing"

var globalWords string

func BenchmarkChooseWords(b *testing.B) {
	words := ""

	for i := 0; i < b.N; i++ {
		words = ChooseWords(10)
	}

	globalWords = words
}
