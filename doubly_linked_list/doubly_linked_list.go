package doublylinkedlist

import "fmt"

type Node struct {
	Value any
	Next  *Node
	Prev  *Node
}

func newNode(value any, next, prev *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
		Prev:  prev,
	}
}

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dl *DoublyLinkedList) isZeroOrOneNode() (bool, *Node) {
	if dl.Length == 0 {
		return true, nil
	} else if dl.Length == 1 {
		value := dl.Head
		dl.Head = nil
		dl.Tail = nil
		dl.Length--
		return true, value
	}
	return false, nil
}

func (dl *DoublyLinkedList) Push(value any) *DoublyLinkedList {
	var node *Node
	if dl.Length == 0 {
		node = newNode(value, nil, nil)
		dl.Head = node
	} else {
		node = newNode(value, nil, dl.Tail)
		dl.Tail.Next = node
	}
	dl.Tail = node
	dl.Length++
	return dl
}

func (dl *DoublyLinkedList) Pop() *Node {
	if ok, node := dl.isZeroOrOneNode(); ok {
		return node
	}
	node := dl.Tail
	dl.Tail = dl.Tail.Prev
	dl.Tail.Next = nil
	dl.Length--
	return node
}

func (dl *DoublyLinkedList) Shift() *Node {
	if ok, node := dl.isZeroOrOneNode(); ok {
		return node
	}
	node := dl.Head
	dl.Head = dl.Head.Next
	dl.Length--
	return node
}

func (dl *DoublyLinkedList) UnShift(value any) *DoublyLinkedList {
	var node *Node
	if dl.Length == 0 {
		node = newNode(value, nil, nil)
		dl.Tail = node
	} else {
		node = newNode(value, dl.Head, nil)
		dl.Head.Prev = node
	}
	dl.Head = node
	dl.Length++
	return nil
}

func (dl *DoublyLinkedList) Set(index int, value any) error {
	node := dl.Get(index)
	if node == nil {
		return fmt.Errorf("wrong index, min 0 and max %d", dl.Length-1)
	}
	node.Value = value
	return nil
}

func (dl *DoublyLinkedList) Get(index int) *Node {
	if index >= dl.Length || index < 0 {
		return nil
	}
	if index <= dl.Length/2 {
		node := dl.Head
		for i := 0; node != nil; i++ {
			if i == index {
				return node
			}
			node = node.Next
		}
	} else {
		node := dl.Tail
		for i := dl.Length - 1; node != nil; i-- {
			if i == index {
				return node
			}
			node = node.Prev
		}
	}
	return nil
}
func (dl *DoublyLinkedList) Insert(index int, value any) *DoublyLinkedList
func (dl *DoublyLinkedList) Remove(index int) (value any)
func (dl *DoublyLinkedList) Reverse() *DoublyLinkedList
func (dl *DoublyLinkedList) Traverse() chan *Node
func (dl *DoublyLinkedList) String() string {
	printable_list := ""

	node := dl.Head
	for node != nil {
		printable_list += fmt.Sprint(node.Value)
		node = node.Next
		if node != nil {
			printable_list += " <=> "
		}
	}
	return printable_list
}
