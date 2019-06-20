package stopwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_registerStopWords(t *testing.T) {
	registerStopWords([]string{"foo", "bar"})
	assert.Subset(t, stopWords, []string{"foo", "bar"})
}

func Test_IsStopWord(t *testing.T) {
	registerStopWords([]string{"foo", "bar"})
	assert.True(t, IsStopWord("foo"))
	assert.True(t, IsStopWord("bar"))
}

func Test_IsStopWord_Negative(t *testing.T) {
	registerStopWords([]string{"foo", "bar"})
	assert.False(t, IsStopWord("bee"))
}

func Test_Slice(t *testing.T) {
	registerStopWords([]string{"foo", "bar"})
	assert.Subset(t, Slice(), []string{"foo", "bar"})
}
