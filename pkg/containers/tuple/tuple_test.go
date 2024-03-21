package tuple_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/containers/tuple"
)

func TestNew(t *testing.T) {
	t.Parallel()

	value := tuple.New("foo", 42)

	assert.Equal(t, "foo", value.First())
	assert.Equal(t, "foo", tuple.First(value))
	assert.Equal(t, 42, value.Second())
	assert.Equal(t, 42, tuple.Second(value))

	a, b := value.Unpack()

	assert.Equal(t, "foo", a)
	assert.Equal(t, 42, b)
}
