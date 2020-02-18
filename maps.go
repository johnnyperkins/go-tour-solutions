package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordsCnt := make(map[string]int)

	for _, word := range words {
		cnt, ok := wordsCnt[word]

		if ok {
			wordsCnt[word] = cnt + 1
		} else {
			wordsCnt[word] = 1
		}
	}

	return wordsCnt
}

func main() {
	wc.Test(WordCount)
}
