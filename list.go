package iter

// List is a dynamicly sized and ordered collection of elements.
// List is an implementation of a a doubly linked list.
// This allows for elements to removed and added in constant time.
// Items in this list are not indexed and iteration is required to find
// specific elements.
type List[V any] struct {
	front, back *Node[V]
}

// Initializes a new List with the option to specify iterables to use as
// elements within the List. (This will advance the iterators to completion.)
func NewList[V any](inputs ...Iter[V]) *List[V] {
	list := &List[V]{}

	for _, input := range inputs {
		for input.HasNext() {
			list.PushNodeBack(&Node[V]{Value: input.Next()})
		}
	}

	return list
}

// Node is an instance of a element within the doubly linked list.
// A node has ref to both the next and previous elements in the List.
// If either Prev or Next are nil, there are no more elements to traverse.
type Node[V any] struct {
	Value      V
	Prev, Next *Node[V]
}

// Provides the node at the front of the list.
func (l *List[V]) Front() *Node[V] {
	return l.front
}

// Provides the node at the end of the list.
func (l *List[V]) Back() *Node[V] {
	return l.back
}

// Pushes Element of value V at the end of the list.
func (l *List[V]) PushBack(v V) *Node[V] {
	node := &Node[V]{
		Value: v,
	}

	l.PushNodeBack(node)

	return node
}

// Pushes Element of value V at the front of the list.
func (l *List[V]) PushFront(v V) *Node[V] {
	node := &Node[V]{
		Value: v,
	}

	l.PushNodeFront(node)

	return node
}

// Pushes node to front of the list. Preforming this operation unlinks the node
// from whatever prior list the node was apart of.
func (l *List[V]) PushNodeFront(node *Node[V]) {
	l.Remove(node)

	node.Next = l.front
	node.Prev = nil

	if l.front != nil {
		l.front.Prev = node
	} else {
		l.back = node
	}

	l.front = node
}

// Pushes node to back of the list. Preforming this operation unlinks the node
// from whatever prior list the node was apart of.
func (l *List[V]) PushNodeBack(node *Node[V]) {
	l.Remove(node)

	node.Next = nil
	node.Prev = l.back

	if l.back != nil {
		l.back.Next = node
	} else {
		l.front = node
	}

	l.back = node
}

// Removes the node at the front of the list and returns the value of the
// removed node.
func (l *List[V]) PopFront() V {
	if l.front == nil {
		var v V

		return v
	}

	current := l.front
	l.Remove(current)

	return current.Value
}

// Removes the node at the end of the list and returns the value of the
// removed node.
func (l *List[V]) PopBack() V {
	if l.back == nil {
		var v V

		return v
	}

	current := l.back
	l.Remove(current)

	return current.Value
}

// Removes the node from the list.
func (l *List[V]) Remove(node *Node[V]) {
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node == l.back {
		l.back = node.Prev
	}

	if node == l.front {
		l.front = node.Next
	}
}

type linkedListIter[V any] struct {
	current *Node[V]
}

func (l *linkedListIter[V]) Next() V {
	if !l.HasNext() {
		var v V

		return v
	}

	value := l.current.Value
	l.current = l.current.Next

	return value
}

func (l *linkedListIter[V]) HasNext() bool {
	return l.current != nil
}

// Creates an iterator which iterates over all the elements in order.
func (l *List[V]) Iter() Iter[V] {
	return &linkedListIter[V]{current: l.front}
}

type reverseLinkedListIter[V any] struct {
	current *Node[V]
}

func (l *reverseLinkedListIter[V]) Next() V {
	if !l.HasNext() {
		var v V

		return v
	}

	value := l.current.Value
	l.current = l.current.Prev

	return value
}

func (l *reverseLinkedListIter[V]) HasNext() bool {
	return l.current != nil
}

// Creates an iterator which iterates over all the elements in order.
func (l *List[V]) Reverse() Iter[V] {
	return &reverseLinkedListIter[V]{current: l.back}
}
