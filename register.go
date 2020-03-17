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

// Setup configures stpwords.
func Setup(opts ...Option) {
	c := &config{
		words: make([][]string, 0),
	}

	// apply custom configuration
	for _, option := range opts {
		option(c)
	}

	for _, ws := range c.words {
		registerStopWords(ws)
	}
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
// listed in `stopwords.go`.
func registerDefaultStopWords() {
	Setup(Words(StopWords))
}

// IsStopWord returns true if given string as a stop-word.
func IsStopWord(s string) bool {
	return stopWordsIndex[s]
}

// Slice returns list of all registered stop-words.
func Slice() []string {
	return append([]string(nil), stopWords...)
}
