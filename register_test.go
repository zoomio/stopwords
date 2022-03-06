package stopwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Setup_noOptions(t *testing.T) {
	r := Setup()
	assert.Subset(t, r.stopWords, []string{"the", "an"})
}

func Test_registerStopWords(t *testing.T) {
	r := create()
	registerStopWords([]string{"foo", "bar"}, r)
	assert.Subset(t, r.stopWords, []string{"foo", "bar"})
}

func Test_IsStopWord(t *testing.T) {
	r := create()
	registerStopWords([]string{"foo", "bar"}, r)
	assert.True(t, r.IsStopWord("foo"))
	assert.True(t, r.IsStopWord("bar"))
}

func Test_IsStopWord_Negative(t *testing.T) {
	r := create()
	registerStopWords([]string{"foo", "bar"}, r)
	assert.False(t, r.IsStopWord("bee"))
}

func Test_Slice(t *testing.T) {
	r := create()
	registerStopWords([]string{"foo", "bar"}, r)
	assert.Subset(t, r.Slice(), []string{"foo", "bar"})
}

func Test_Slice_NoDupes(t *testing.T) {
	r := create()
	registerStopWords([]string{"foo", "foo"}, r)
	assert.Subset(t, r.Slice(), []string{"foo"})
}
