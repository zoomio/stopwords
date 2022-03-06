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
	Words = func(words []string) Option {
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
	WithDomains = func(v bool) Option {
		return func(c *config) {
			if v {
				c.words = append(c.words, strings.Split(Domains, "\n"))
			}
		}
	}
)
