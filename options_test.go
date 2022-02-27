package stopwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WithDomains(t *testing.T) {
	reg := Setup(WithDomains(true))
	assert.True(t, reg.IsStopWord("com"))
	assert.True(t, reg.IsStopWord("net"))
	assert.True(t, reg.IsStopWord("au"))
	assert.True(t, reg.IsStopWord("org"))
	assert.True(t, reg.IsStopWord("ai"))
}
