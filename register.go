package stopwords

import (
	"strings"
)

// Register an instance of the register of the stopwords.
type Register struct {
	// stopWordsIndex - vocabulary of stop-words for quick search
	stopWordsIndex map[string]bool
	stopWords      []string
}

// IsStopWord returns true if given string as a stop-word.
func (r *Register) IsStopWord(s string) bool {
	return r.stopWordsIndex[s]
}

// Slice returns list of all registered stop-words.
func (r *Register) Slice() []string {
	return append([]string(nil), r.stopWords...)
}

func create() *Register {
	return &Register{
		stopWordsIndex: map[string]bool{},
		stopWords:      []string{},
	}
}

// Setup configures stpwords.
func Setup(opts ...Option) *Register {
	c := &config{
		words: make([][]string, 0),
	}

	if len(opts) == 0 {
		return registerDefaultStopWords()
	}

	// apply custom configuration
	for _, option := range opts {
		option(c)
	}

	reg := create()

	for _, ws := range c.words {
		registerStopWords(ws, reg)
	}

	return reg
}

// registerStopWords registers given list of stop-words to use as vocabulary.
func registerStopWords(words []string, reg *Register) {
	for _, s := range words {
		w := strings.TrimSpace(s)
		if w == "" || reg.stopWordsIndex[w] {
			continue
		}
		reg.stopWords = append(reg.stopWords, w)
		reg.stopWordsIndex[w] = true
	}
}

// registerDefaultStopWords registers default stop words
// listed in `stopwords.go`.
func registerDefaultStopWords() *Register {
	return Setup(Words(StopWords))
}
