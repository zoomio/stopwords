package stopwords

import (
	"strings"
)

var (
	// StopWordsIndex - vocabulary of stop-words for quick search on contains
	stopWordsIndex = make(map[string]bool)
	stopWords      = make([]string, 0)
)

func init() {
	registerStopWords(strings.Split(StopWords, "\n"))
}

// registerStopWords registers given list of stop-words to use as vocabulary.
func registerStopWords(words []string) {
	stopWords = append(stopWords, words...)
	for _, s := range words {
		stopWordsIndex[s] = true
	}
}

// IsStopWord returns true if given string as a stop-word.
func IsStopWord(s string) bool {
	return stopWordsIndex[s]
}

// Slice returns list of all registered stop-words.
func Slice() []string {
	return append([]string(nil), stopWords...)
}
