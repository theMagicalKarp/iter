package itertools_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter"
	"github.com/theMagicalKarp/iter/itertools"
)

func TestStringBasic(t *testing.T) {
	t.Parallel()

	res := itertools.String(iter.NewIter("a", "bc", "d", "e", "fff"))

	assert.Equal(t, res, "abcdefff")
}

func TestStringEmpty(t *testing.T) {
	t.Parallel()

	res := itertools.String(iter.NewIter[string]())

	assert.Equal(t, res, "")
}
