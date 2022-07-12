package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	words := strings.Fields(text)

	if len(words) == 0 {
		return nil
	}

	sortedWords := sortWords(getCountedWordMap(words))
	if len(sortedWords) > 10 {
		return sortedWords[:10]
	}

	return sortedWords
}

func getCountedWordMap(words []string) map[int][]string {
	cache := make(map[string]int)
	for _, word := range words {
		cache[word]++
	}

	countedWordMap := make(map[int][]string)
	for word, count := range cache {
		countedWordMap[count] = append(countedWordMap[count], word)
	}

	return countedWordMap
}

func sortWords(wordMap map[int][]string) []string {
	sortedKeys := make([]int, 0, len(wordMap))
	for k := range wordMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return sortedKeys[i] > sortedKeys[j]
	})

	var sortedSlice []string
	for _, k := range sortedKeys {
		sort.Strings(wordMap[k])
		sortedSlice = append(sortedSlice, wordMap[k]...)
	}

	return sortedSlice
}
