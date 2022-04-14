package iter

import (
	"golang.org/x/exp/constraints"
)

// Set is an unordered collection with no duplicate elements.
// Elements must be primitives which belong to `constraints.Ordered`.
// Basic uses include membership testing and eliminating duplicate entries.
// This set can return an iterable interface while maintining O(1) time
// complexity for adding, removing, and finding elements.
type Set[K constraints.Ordered] struct {
	items map[K]*Node[K]
	nodes *List[K]
}

// Initializes a new Set with the option to specify iterables to use as
// elements within the Set. (This will advance the iterators to completion.)
func NewSet[K constraints.Ordered](inputs ...Iter[K]) *Set[K] {
	set := &Set[K]{
		items: map[K]*Node[K]{},
		nodes: NewList[K](),
	}

	for _, input := range inputs {
		for input.HasNext() {
			set.Add(input.Next())
		}
	}

	return set
}

// Checks if the element belongs to set.
func (s *Set[K]) Has(key K) bool {
	_, ok := s.items[key]

	return ok
}

// Adds an element to the set.
func (s *Set[K]) Add(key K) {
	if s.Has(key) {
		return
	}

	node := &Node[K]{
		Value: key,
	}
	s.items[key] = node
	s.nodes.PushNodeBack(node)
}

// Removes an element from the set.
func (s *Set[K]) Remove(key K) {
	if !s.Has(key) {
		return
	}

	node := s.items[key]
	s.nodes.Remove(node)
	delete(s.items, key)
}

// Provides the current size of the set.
func (s *Set[K]) Len() int {
	return len(s.items)
}

// Subtracts the current set from the provided set. This results in a brand new
// Set struct being created.
func (s *Set[K]) Sub(other *Set[K]) *Set[K] {
	newSet := NewSet(s.Iter())
	otherIter := other.Iter()

	for otherIter.HasNext() {
		newSet.Remove(otherIter.Next())
	}

	return newSet
}

// Combines the current set with the provided set. This results in a brand new
// Set struct being created.
func (s *Set[K]) Union(other *Set[K]) *Set[K] {
	newSet := NewSet(s.Iter())
	otherIter := other.Iter()

	for otherIter.HasNext() {
		newSet.Add(otherIter.Next())
	}

	return newSet
}

// Only fetches elements which belong to both current and provided sets.
// This results in band new Set struct being created.
func (s *Set[K]) Intersection(other *Set[K]) *Set[K] {
	newSet := NewSet[K]()
	otherIter := other.Iter()

	for otherIter.HasNext() {
		next := otherIter.Next()

		if s.Has(next) {
			newSet.Add(next)
		}
	}

	return newSet
}

type setIter[K constraints.Ordered] struct {
	current *Node[K]
}

func (s *setIter[K]) Next() K {
	value := s.current.Value
	s.current = s.current.Next

	return value
}

func (s *setIter[K]) HasNext() bool {
	return s.current != nil
}

// Creates an iterable struct which can be used to iterate over ever element
// in the set.
func (s *Set[K]) Iter() Iter[K] {
	return &setIter[K]{current: s.nodes.front}
}
