package iter

// Array is an orderd collection of indexed elements.
type Array[T any] struct {
	items []T
}

// Creates a new instance of Array.
// This constructor accepts size and capacity which get used to when
// making the slice to store the elements. Any additional iterators
// passsed are stored as elements in the array.
func NewArray[T any](size, capacity int, inputs ...Iter[T]) *Array[T] {
	array := &Array[T]{
		items: make([]T, size, capacity),
	}

	for _, input := range inputs {
		for input.HasNext() {
			array.items = append(array.items, input.Next())
		}
	}

	return array
}

// Sets element at index "i". Can panic if out of bounds of size.
func (a *Array[T]) Set(i int, item T) {
	a.items[i] = item
}

// Returns element at index "i".
func (a *Array[T]) At(i int) T {
	return a.items[i]
}

// Appends element to end of array.
func (a *Array[T]) Append(item T) {
	a.items = append(a.items, item)
}

// Removes and returns element from end of array.
func (a *Array[T]) Pop() T {
	ret := a.items[len(a.items)-1]
	a.items = a.items[:len(a.items)-1]

	return ret
}

// Fetches the number of elements in the array.
func (a *Array[T]) Len() int {
	return len(a.items)
}

type arrayIter[T any] struct {
	array *Array[T]
	index int
	end   int
}

func (a *arrayIter[T]) Next() T {
	value := a.array.At(a.index)
	a.index++

	return value
}

func (a *arrayIter[T]) HasNext() bool {
	return a.index < a.end
}

// Creates an iterator which iterates over all the elements in order.
func (a *Array[T]) Iter() Iter[T] {
	return &arrayIter[T]{array: a, index: 0, end: a.Len()}
}

// Creates an iterator which iterates over all the elements after provided index in order.
func (a *Array[T]) Tail(start int) Iter[T] {
	return &arrayIter[T]{array: a, index: start, end: a.Len()}
}

// Creates an iterator which iterates over all the elements before provided index in order.
func (a *Array[T]) Head(end int) Iter[T] {
	return &arrayIter[T]{array: a, index: 0, end: end}
}

// Creates an iterator which iterates over the specified range of indexes in order.
func (a *Array[T]) Range(start, end int) Iter[T] {
	return &arrayIter[T]{array: a, index: start, end: end}
}

type revArrayIter[T any] struct {
	array *Array[T]
	index int
}

func (a *revArrayIter[T]) Next() T {
	value := a.array.At(a.index)
	a.index--

	return value
}

func (a *revArrayIter[T]) HasNext() bool {
	return a.index >= 0
}

// Creates an iterator which iterates over all the elements in order.
func (a *Array[T]) Reverse() Iter[T] {
	return &revArrayIter[T]{array: a, index: a.Len() - 1}
}
