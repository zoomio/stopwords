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
	// WordsSlice registers given word set.
	WordsSlice = func(words []string) Option {
		return func(c *config) {
			c.words = append(c.words, words)
		}
	}
	// Text splits provided text based on the sep.
	Text = func(text, sep string) Option {
		return func(c *config) {
			c.words = append(c.words, strings.Split(text, sep))
		}
	}
	// Words registers given word set.
	Words = func(words string) Option {
		return func(c *config) {
			Text(words, "\n")(c)
		}
	}
	WithDomains = func(v bool) Option {
		return func(c *config) {
			if v {
				Text(Domains, "\n")(c)
			}
		}
	}
)
