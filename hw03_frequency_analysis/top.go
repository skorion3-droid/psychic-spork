package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordFr struct {
	Word string
	Fr   int
}

func Top10(text string) []string {
	if strings.TrimSpace(text) == "" {
		return []string{}
	}
	// Place your code here.
	words := strings.Fields(text)
	mapa := make(map[string]int)

	for _, word := range words {
		value, exists := mapa[word]

		if exists {
			mapa[word] = value + 1
		} else {
			mapa[word] = 1
		}
	}

	preResult := make([]WordFr, 0)

	for word, count := range mapa {
		preResult = append(preResult, WordFr{word, count})
	}

	sort.Slice(preResult, func(i, j int) bool {
		if preResult[i].Fr == preResult[j].Fr {
			return preResult[i].Word < preResult[j].Word
		}
		return preResult[i].Fr > preResult[j].Fr
	})

	result := make([]string, 0)

	for i, wordFrequency := range preResult {
		if i == 10 {
			break
		}
		result = append(result, wordFrequency.Word)
	}

	return result
}
