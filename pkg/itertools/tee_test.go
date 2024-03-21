package itertools_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/iter/pkg/iter"
	"github.com/theMagicalKarp/iter/pkg/itertools"
)

func ExampleTee() {
	a, b := itertools.Tee(
		iter.New(1, 2, 3, 4),
	)

	itertools.Print(a)
	itertools.Print(b)
	// Output:
	// 1, 2, 3, 4
	// 1, 2, 3, 4
}

func TestTeeBasic(t *testing.T) {
	t.Parallel()

	a, b := itertools.Tee(
		iter.New(1, 2, 3, 4),
	)

	v, more := a.Next()
	assert.True(t, more)
	assert.Equal(t, v, 1)

	v, more = b.Next()
	assert.True(t, more)
	assert.Equal(t, v, 1)

	v, more = a.Next()
	assert.True(t, more)
	assert.Equal(t, v, 2)

	v, more = b.Next()
	assert.True(t, more)
	assert.Equal(t, v, 2)

	v, more = b.Next()
	assert.True(t, more)
	assert.Equal(t, v, 3)

	v, more = b.Next()
	assert.True(t, more)
	assert.Equal(t, v, 4)

	v, more = a.Next()
	assert.True(t, more)
	assert.Equal(t, v, 3)

	v, more = a.Next()
	assert.True(t, more)
	assert.Equal(t, v, 4)

	v, more = a.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)

	v, more = b.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)

	v, more = a.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)

	v, more = b.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)
}

func TestTeeLopSided(t *testing.T) {
	t.Parallel()

	a, b := itertools.Tee(
		iter.New(1, 2, 3),
	)

	v, more := a.Next()
	assert.True(t, more)
	assert.Equal(t, v, 1)

	v, more = a.Next()
	assert.True(t, more)
	assert.Equal(t, v, 2)

	v, more = a.Next()
	assert.True(t, more)
	assert.Equal(t, v, 3)

	v, more = a.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)

	v, more = a.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)

	v, more = b.Next()
	assert.True(t, more)
	assert.Equal(t, v, 1)

	v, more = b.Next()
	assert.True(t, more)
	assert.Equal(t, v, 2)

	v, more = b.Next()
	assert.True(t, more)
	assert.Equal(t, v, 3)

	v, more = b.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)

	v, more = b.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)
}

func TestTeeEmpty(t *testing.T) {
	t.Parallel()

	a, b := itertools.Tee(iter.New[int]())

	v, more := a.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)

	v, more = a.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)

	v, more = b.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)

	v, more = b.Next()
	assert.False(t, more)
	assert.Equal(t, v, 0)
}

func TestAsyncPull(t *testing.T) {
	t.Parallel()

	r1 := rand.New(rand.NewSource(42))
	r2 := rand.New(rand.NewSource(7))

	n := 100000

	a, b := itertools.Tee(itertools.Range(0, n))
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			time.Sleep(time.Duration(r1.Intn(10)) * time.Microsecond)
			v, more := a.Next()
			assert.True(t, more)
			assert.Equal(t, v, i)
		}
		v, more := a.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)

		v, more = a.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			time.Sleep(time.Duration(r2.Intn(15)) * time.Microsecond)
			v, more := b.Next()
			assert.True(t, more)
			assert.Equal(t, v, i)
		}
		v, more := b.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)

		v, more = b.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)
	}()

	wg.Wait()
}

func TestAsyncPullLopside(t *testing.T) {
	t.Parallel()

	r1 := rand.New(rand.NewSource(42))

	n := 100000

	a, b := itertools.Tee(itertools.Range(0, n))
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			v, more := a.Next()
			assert.True(t, more)
			assert.Equal(t, v, i)
			time.Sleep(time.Duration(r1.Intn(15)) * time.Microsecond)
		}
		v, more := a.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)

		v, more = a.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			v, more := b.Next()
			assert.True(t, more)
			assert.Equal(t, v, i)
		}
		v, more := b.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)

		v, more = b.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)
	}()

	wg.Wait()
}

func TestAsyncPullFast(t *testing.T) {
	t.Parallel()

	n := 100000

	a, b := itertools.Tee(itertools.Range(0, n))
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			v, more := a.Next()
			assert.True(t, more)
			assert.Equal(t, v, i)
		}
		v, more := a.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)

		v, more = a.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			v, more := b.Next()
			assert.True(t, more)
			assert.Equal(t, v, i)
		}
		v, more := b.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)

		v, more = b.Next()
		assert.False(t, more)
		assert.Equal(t, v, 0)
	}()

	wg.Wait()
}
