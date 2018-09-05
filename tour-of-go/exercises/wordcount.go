package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(ignoreCase bool) func(string) map[string]int {
	return func(s string) map[string]int {
		counts := make(map[string]int)
		words := strings.Fields(s)
		for i := 0; i < len(words); i++ {
			word := words[i]

			if ignoreCase {
				word = strings.ToLower(word)

			}

			v, exists := counts[word]
			if !exists {
				counts[word] = 1
			} else {
				counts[word] = 1 + v
			}
		}
		return counts
	}
}

func main() {
	counts := wordCount(true)("Hello hello World world test")
	fmt.Println(counts)
	wc.Test(wordCount(false))
}
