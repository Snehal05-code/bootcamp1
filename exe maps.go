package main

import (
	"strings"

	"golang.org/x/exercises/wordcount/wc"
)

func WordCount(s string) map[string]int {
	// Create empty map
	counts := make(map[string]int)

	// Split string into words
	words := strings.Fields(s)

	// Count each word
	for _, w := range words {
		counts[w]++
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
