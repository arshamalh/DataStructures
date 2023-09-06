package singlylinkedlist

import "fmt"

type Node struct {
	Value any
	Next  *Node
}

func newNode(value any, next *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}

type SinglyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func New() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (sl *SinglyLinkedList) isZeroOrOneNode() (bool, *Node) {
	if sl.Length == 0 {
		return true, nil
	} else if sl.Length == 1 {
		value := sl.Head
		sl.Head = nil
		sl.Tail = nil
		sl.Length--
		return true, value
	}
	return false, nil
}

func (sl *SinglyLinkedList) Push(value any) *SinglyLinkedList {
	new_node := newNode(value, nil)
	if sl.Length == 0 {
		sl.Head = new_node
	} else {
		sl.Tail.Next = new_node
	}
	sl.Tail = new_node
	sl.Length++
	return sl
}

func (sl *SinglyLinkedList) Pop() *Node {
	if ok, value := sl.isZeroOrOneNode(); ok {
		return value
	}
	node := sl.Head
	new_tail := node
	for node.Next != nil {
		new_tail = node
		node = node.Next
	}
	sl.Tail = new_tail
	sl.Tail.Next = nil
	sl.Length--
	return node
}

func (sl *SinglyLinkedList) Shift() *Node {
	if ok, node := sl.isZeroOrOneNode(); ok {
		return node
	}
	node := sl.Head
	sl.Head = sl.Head.Next
	sl.Length--
	return node
}

func (sl *SinglyLinkedList) UnShift(value any) *SinglyLinkedList {
	var new_node *Node
	if sl.Length == 0 {
		new_node = newNode(value, nil)
		sl.Tail = new_node
	} else {
		new_node = newNode(value, sl.Head)
	}
	sl.Head = new_node
	sl.Length++
	return sl
}

func (sl *SinglyLinkedList) Get(index int) *Node {
	if index >= sl.Length || index < 0 {
		return nil
	}
	node := sl.Head
	for i := 0; i < sl.Length; i++ {
		if i == index {
			return node
		}
		node = node.Next
	}
	return nil
}

func (sl *SinglyLinkedList) Set(index int, value any) error {
	node := sl.Get(index)
	if node == nil {
		return fmt.Errorf("wrong index, min 0 and max %d", sl.Length-1)
	}
	node.Value = value
	return nil
}

func (sl *SinglyLinkedList) Insert(index int, value any) *SinglyLinkedList {
	if index == 0 {
		return sl.UnShift(value)
	} else if index == sl.Length {
		return sl.Push(value)
	} else if index > sl.Length || index < 0 {
		return nil
	}

	back_node := sl.Get(index - 1)
	current_node := back_node.Next
	back_node.Next = newNode(value, current_node)
	sl.Length++
	return sl
}

func (sl *SinglyLinkedList) Remove(index int) *Node {
	if ok, value := sl.isZeroOrOneNode(); ok {
		return value
	} else if index == 0 {
		return sl.Shift()
	} else if index == sl.Length-1 {
		return sl.Pop()
	} else if index >= sl.Length || index < 0 {
		return nil
	}

	back_node := sl.Get(index - 1)
	current_node := back_node.Next
	next_node := current_node.Next
	back_node.Next = next_node
	sl.Length--
	return current_node
}

func (sl *SinglyLinkedList) Reverse() *SinglyLinkedList {
	node := sl.Head
	sl.Head, sl.Tail = sl.Tail, sl.Head
	var next, prev *Node
	for node != nil {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return sl
}

func (sl *SinglyLinkedList) Traverse(result chan *Node) {
	node := sl.Head
	for node != nil {
		result <- node
		node = node.Next
	}
	close(result)
}

func (sl *SinglyLinkedList) String() string {
	printable_list := ""
	node := sl.Head
	for node != nil {
		printable_list += fmt.Sprint(node.Value)
		node = node.Next
		if node != nil {
			printable_list += " => "
		}
	}
	return printable_list
}
