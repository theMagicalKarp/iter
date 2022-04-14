package iter

import (
	"golang.org/x/exp/constraints"
)

// Map is an unordered collection of elements indexed by value.
// Elements must be indexed by primitives which belong to `constraints.Ordered`.
// This map can return an iterable interface while maintining O(1) time
// complexity for adding, removing, and finding elements.
type Map[K constraints.Ordered, V any] struct {
	items map[K]*Node[Tuple[K, V]]
	nodes *List[Tuple[K, V]]
}

// Initializes a new Map with the option to specify iterables to use as
// elements within the Map. (This will advance the iterators to completion.)
// These iterables must provide a Tuple as key/value paris, where "First"
// represents the key, while "Second" represents the value.
func NewMap[K constraints.Ordered, V any](inputs ...Iter[Tuple[K, V]]) *Map[K, V] {
	newMap := &Map[K, V]{
		items: map[K]*Node[Tuple[K, V]]{},
		nodes: NewList[Tuple[K, V]](),
	}

	for _, input := range inputs {
		for input.HasNext() {
			next := input.Next()
			newMap.Set(next.First, next.Second)
		}
	}

	return newMap
}

// Sets values V at index K.
func (m *Map[K, V]) Set(key K, value V) {
	if _, ok := m.items[key]; !ok {
		node := &Node[Tuple[K, V]]{
			Value: Tuple[K, V]{First: key, Second: value},
		}
		m.items[key] = node
		m.nodes.PushNodeBack(node)
	}

	m.items[key].Value.Second = value
}

// Removes value at index K.
func (m *Map[K, V]) Remove(key K) {
	if _, ok := m.items[key]; !ok {
		return
	}

	node := m.items[key]
	m.nodes.Remove(node)
	delete(m.items, key)
}

// Checks if values exists at index K.
func (m *Map[K, V]) Has(k K) bool {
	_, ok := m.items[k]

	return ok
}

// Get values exists at index K. (if element does not exist, "empty"
// value is returned).
func (m *Map[K, V]) Get(k K) V {
	node, ok := m.items[k]
	if !ok {
		var v V

		return v
	}

	return node.Value.Second
}

// Fetches size of map.
func (m *Map[K, V]) Len() int {
	return len(m.items)
}

// Creates an iterator which iterates over all the keys.
func (m *Map[K, V]) Keys() Iter[K] {
	return &keysIter[K, V]{current: m.nodes.front}
}

// Creates an iterator which iterates over all the key/value pairs.
func (m *Map[K, V]) Items() Iter[Tuple[K, V]] {
	return &itemsIter[K, V]{current: m.nodes.front}
}

// Creates an iterator which iterates over all the values.
func (m *Map[K, V]) Values() Iter[V] {
	return &valuesIter[K, V]{current: m.nodes.front}
}

type keysIter[K constraints.Ordered, V any] struct {
	current *Node[Tuple[K, V]]
}

func (m *keysIter[K, V]) Next() K {
	value := m.current.Value.First
	m.current = m.current.Next

	return value
}

func (m *keysIter[K, V]) HasNext() bool {
	return m.current != nil
}

type valuesIter[K constraints.Ordered, V any] struct {
	current *Node[Tuple[K, V]]
}

func (m *valuesIter[K, V]) Next() V {
	value := m.current.Value.Second
	m.current = m.current.Next

	return value
}

func (m *valuesIter[K, V]) HasNext() bool {
	return m.current != nil
}

type itemsIter[K constraints.Ordered, V any] struct {
	current *Node[Tuple[K, V]]
}

func (m *itemsIter[K, V]) Next() Tuple[K, V] {
	value := m.current.Value
	m.current = m.current.Next

	return value
}

func (m *itemsIter[K, V]) HasNext() bool {
	return m.current != nil
}
