package predicates_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
	"github.com/theMagicalKarp/iter/pkg/predicates"
)

func ExampleIsError() {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	items := itertools.Filter(iter.New(err1, err2, err3), predicates.IsError(err1))

	itertools.Print(items)
	// Output:
	// error 1
}

func ExampleHasError() {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	in := iter.New(
		tuple.New(1, err1),
		tuple.New(2, err2),
		tuple.New(3, err3),
	)

	items := itertools.Filter(in, predicates.HasError[int](err1))

	values := itertools.Map(items, func(t tuple.Tuple[int, error]) int { return t.First() })
	itertools.Print(values)
	// Output:
	// 1
}

func TestIsError(t *testing.T) {
	t.Parallel()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	isErr1 := predicates.IsError(err1)

	assert.True(t, isErr1(err1))
	assert.False(t, isErr1(err2))
	assert.False(t, isErr1(err3))
}

func TestErrorNested(t *testing.T) {
	t.Parallel()

	err1 := errors.New("error 1")
	err2 := fmt.Errorf(">%w", err1)
	err3 := fmt.Errorf(">%w", err2)

	isErr1 := predicates.IsError(err1)

	assert.True(t, isErr1(err1))
	assert.True(t, isErr1(err2))
	assert.True(t, isErr1(err3))
}

func TestHasError(t *testing.T) {
	t.Parallel()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	hasErr1 := predicates.HasError[int](err1)

	assert.True(t, hasErr1(tuple.New(1, err1)))
	assert.False(t, hasErr1(tuple.New(1, err2)))
	assert.False(t, hasErr1(tuple.New(1, err3)))
}

func TestHasErrorNested(t *testing.T) {
	t.Parallel()

	err1 := errors.New("error 1")
	err2 := fmt.Errorf(">%w", err1)
	err3 := fmt.Errorf(">%w", err2)

	hasErr1 := predicates.HasError[int](err1)

	assert.True(t, hasErr1(tuple.New(1, err1)))
	assert.True(t, hasErr1(tuple.New(1, err2)))
	assert.True(t, hasErr1(tuple.New(1, err3)))
}
