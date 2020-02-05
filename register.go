package stopwords

import (
	"strings"
)

var (
	// StopWordsIndex - vocabulary of stop-words for quick search
	stopWordsIndex map[string]bool
	stopWords      []string
)

func init() {
	stopWords = make([]string, 0)
	stopWordsIndex = make(map[string]bool)
	registerDefaultStopWords()
}

// registerStopWords registers given list of stop-words to use as vocabulary.
func registerStopWords(words []string) {
	for _, s := range words {
		w := strings.TrimSpace(s)
		if w == "" || stopWordsIndex[w] {
			continue
		}
		stopWords = append(stopWords, w)
		stopWordsIndex[w] = true
	}
}

// registerDefaultStopWords registers default stop words
// listed in `stopwords.go` and `stopwords_ru.go`.
func registerDefaultStopWords() {
	registerStopWords(strings.Split(StopWords, "\n"))
	registerStopWords(strings.Split(StopWordsRu, "\n"))
}

// IsStopWord returns true if given string as a stop-word.
func IsStopWord(s string) bool {
	return stopWordsIndex[s]
}

// Slice returns list of all registered stop-words.
func Slice() []string {
	return append([]string(nil), stopWords...)
}
