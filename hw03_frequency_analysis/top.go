package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Word struct {
	Value string
	Count int
}

func Top10(text string) []string {
	if len(text) == 0 {
		return make([]string, 0)
	}

	dict := make(map[string]int)
	fields := strings.Fields(text)

	for _, w := range fields {
		dict[w]++
	}

	words := make([]Word, 0, len(dict))
	for k, v := range dict {
		words = append(words, Word{Value: k, Count: v})
	}

	sort.SliceStable(words, func(i, j int) bool {
		if words[i].Count == words[j].Count {
			return words[i].Value < words[j].Value
		}

		return words[i].Count > words[j].Count
	})

	var wordAmount int
	if len(words) > 10 {
		wordAmount = 10
	} else {
		wordAmount = len(words)
	}

	words = words[:wordAmount]

	topWords := make([]string, wordAmount)
	for i, w := range words {
		topWords[i] = w.Value
	}

	return topWords
}
