package queue

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func newNode[T any](value T, next *Node[T]) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  next,
	}
}

type Queue[T any] struct {
	Length int
	head   *Node[T]
	tail   *Node[T]
}

// Add to the end, remove from the beginning
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) isZeroOrOneNode() (bool, *Node[T]) {
	if q.Length == 0 {
		return true, nil
	} else if q.Length == 1 {
		value := q.head
		q.head = nil
		q.tail = nil
		q.Length--
		return true, value
	}
	return false, nil
}

func (q *Queue[T]) Push(value T) *Queue[T] {
	new_node := newNode(value, nil)
	if q.Length == 0 {
		q.head = new_node
	} else {
		q.tail.Next = new_node
	}
	q.tail = new_node
	q.Length++
	return q
}

func (q *Queue[T]) Pop() *Node[T] {
	if ok, node := q.isZeroOrOneNode(); ok {
		return node
	}
	node := q.head
	q.head = q.head.Next
	q.Length--
	return node
}

func (q *Queue[T]) String() string {
	printable_list := ""
	node := q.head
	for node != nil {
		printable_list = fmt.Sprint(node.Value) + " => " + printable_list
		node = node.Next
	}
	return printable_list
}
