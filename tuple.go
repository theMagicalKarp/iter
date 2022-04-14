package iter

// Tuple is a structure which contains two distinct values. These values are
// typically used for representing key/values pairs of maps. Other useful uses
// for tuples are for representing element/error paris.
type Tuple[T any, V any] struct {
	First  T
	Second V
}
