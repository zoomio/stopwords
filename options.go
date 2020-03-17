package stopwords

import (
	"strings"
)

// Option allows to customise stopwords configuration.
type Option func(*config)

type config struct {
	words [][]string
}

var (
	// Words registers given word set.
	Words = func(words string) Option {
		return func(c *config) {
			c.words = append(c.words, strings.Split(words, "\n"))
		}
	}
)
