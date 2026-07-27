// Package tuple provides a small generic container holding two values of
// independent types, useful for representing key/value or value/error pairs
// as they flow through an iterable.
package tuple

// Tuple is a structure which contains two distinct values. These values are
// typically used for representing key/values pairs of maps. Other useful uses
// for tuples are for representing element/error paris.
type Tuple[T any, V any] struct {
	first  T
	second V
}

// New creates a new Tuple with the given values.
// It takes two parameters, `t` and `v`, of any type `T` and `V` respectively,
// and returns a Tuple[T, V] with the `t` value assigned to the `first` field
// and the `v` value assigned to the `second` field.
func New[T any, V any](t T, v V) Tuple[T, V] {
	return Tuple[T, V]{first: t, second: v}
}

// First returns the first element of the tuple.
func First[T any, V any](t Tuple[T, V]) T {
	return t.first
}

// First returns the first element of the tuple.
func (t Tuple[T, V]) First() T {
	return t.first
}

// Second returns the second value of the tuple.
func Second[T any, V any](t Tuple[T, V]) V {
	return t.second
}

// Second returns the second value of the tuple.
func (t Tuple[T, V]) Second() V {
	return t.second
}

// Unpack returns the values of the Tuple.
func (t Tuple[T, V]) Unpack() (T, V) {
	return t.first, t.second
}
