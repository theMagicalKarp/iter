package predicates_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

func ExampleContains() {
	items := iter.New("hello", "world", "foo", "bar")
	items = itertools.Filter(items, predicates.Contains("o"))

	itertools.Print(items)
	// Output: hello, world, foo
}

func ExampleHasPrefix() {
	items := iter.New("hello", "world", "foo", "bar", "foot")
	items = itertools.Filter(items, predicates.HasPrefix("foo"))

	itertools.Print(items)
	// Output: foo, foot
}

func ExampleHasSuffix() {
	items := iter.New("helloo", "world", "foo", "bar", "foot")
	items = itertools.Filter(items, predicates.HasSuffix("oo"))

	itertools.Print(items)
	// Output: helloo, foo
}

func TestContains(t *testing.T) {
	t.Parallel()

	substr := "world"
	predicate := predicates.Contains(substr)

	assert.True(t, predicate("hello world"))
	assert.False(t, predicate("hello"))
}

func TestHasPrefix(t *testing.T) {
	t.Parallel()

	prefix := "hello"
	predicate := predicates.HasPrefix(prefix)

	assert.True(t, predicate("hello world"))
	assert.False(t, predicate("world"))
}

func TestHasPrefix_CustomPrefix(t *testing.T) {
	t.Parallel()

	prefix := "foo"
	predicate := predicates.HasPrefix(prefix)

	assert.True(t, predicate("foobar"))
	assert.False(t, predicate("bar"))
}

func TestHasPrefix_EmptyString(t *testing.T) {
	t.Parallel()

	prefix := ""
	predicate := predicates.HasPrefix(prefix)

	assert.True(t, predicate("hello world"))
	assert.True(t, predicate("world"))
}

func TestHasSuffix(t *testing.T) {
	t.Parallel()

	suffix := "world"
	predicate := predicates.HasSuffix(suffix)

	assert.True(t, predicate("hello world"))
	assert.False(t, predicate("hello"))
}
