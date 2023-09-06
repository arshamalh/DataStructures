package stack

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

type Stack[T any] struct {
	head *Node[T]
	Len  int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (sl *Stack[T]) isZeroOrOneNode() (bool, *Node[T]) {
	if sl.Len == 0 {
		return true, nil
	} else if sl.Len == 1 {
		value := sl.head
		sl.head = nil
		sl.Len--
		return true, value
	}
	return false, nil
}

func (sl *Stack[T]) Push(value T) *Stack[T] {
	var new_node *Node[T]
	if sl.Len == 0 {
		new_node = newNode(value, nil)
	} else {
		new_node = newNode(value, sl.head)
	}
	sl.head = new_node
	sl.Len++
	return sl
}

func (sl *Stack[T]) Pop() *Node[T] {
	if ok, node := sl.isZeroOrOneNode(); ok {
		return node
	}
	node := sl.head
	sl.head = sl.head.Next
	sl.Len--
	return node
}

func (sl *Stack[T]) String() string {
	printable_list := ""
	node := sl.head
	for node != nil {
		printable_list += fmt.Sprint(node.Value)
		node = node.Next
		if node != nil {
			printable_list += " <= "
		}
	}
	return printable_list
}
