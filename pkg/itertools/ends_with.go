package itertools

import (
	"github.com/theMagicalKarp/iter/pkg/iter"
)

// EndsWith checks if the given `items` ends with the specified `suffix`.
// It returns `true` if `items` ends with `suffix`, and `false` otherwise.
// The `items` parameter is an iterable of type `T`, and the `suffix` parameter
// is also an iterable of type `T`. The `T` type must be comparable.
func EndsWith[T comparable](items, suffix iter.Iterable[T]) bool {
	suffixSlice := Slice(suffix)

	return Equals[T](Tail(items, len(suffixSlice)), iter.New(suffixSlice...))
}
