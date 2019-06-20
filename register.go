package stopwords

import (
	"strings"
)

var (
	// StopWordsIndex - vocabulary of stop-words for quick search on contains
	stopWordsIndex map[string]bool
	stopWords      []string
)

func init() {
	registerDefaultStopWords()
}

// registerStopWords registers given list of stop-words to use as vocabulary.
func registerStopWords(words []string) {
	stopWords = make([]string, 0)
	stopWords = append(stopWords, words...)

	stopWordsIndex = make(map[string]bool)
	for _, s := range words {
		stopWordsIndex[s] = true
	}
}

// registerDefaultStopWords registers default stop words listed in `stopwords.go`.
func registerDefaultStopWords() {
	registerStopWords(strings.Split(StopWords, "\n"))
}

// IsStopWord returns true if given string as a stop-word.
func IsStopWord(s string) bool {
	return stopWordsIndex[s]
}

// Slice returns list of all registered stop-words.
func Slice() []string {
	return append([]string(nil), stopWords...)
}
