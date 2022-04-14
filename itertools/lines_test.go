package itertools_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestLinesBasic(t *testing.T) {
	t.Parallel()

	items := itertools.Lines(strings.NewReader("foo\nbar\nhello\nworld\n\nsup\n"))

	assert.True(t, items.HasNext())
	assert.Equal(t, "foo", items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, "bar", items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, "hello", items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, "world", items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, "", items.Next())
	assert.True(t, items.HasNext())
	assert.Equal(t, "sup", items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, "", items.Next())
}

func TestLinesNoNewLines(t *testing.T) {
	t.Parallel()

	items := itertools.Lines(strings.NewReader("hello"))

	assert.True(t, items.HasNext())
	assert.Equal(t, "hello", items.Next())
	assert.False(t, items.HasNext())
	assert.Equal(t, "", items.Next())
}

func TestLinesEmpty(t *testing.T) {
	t.Parallel()

	items := itertools.Lines(strings.NewReader(""))

	assert.False(t, items.HasNext())
	assert.Equal(t, "", items.Next())
}
