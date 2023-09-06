package bst

import (
	"ds/queue"
)

type TraverseType string

const (
	// Breadth first search
	BFS TraverseType = "bfs"

	// Depth first search
	DFS TraverseType = "dfs"
)

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint |
		~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float32 | ~float64 | ~string
}

type Node[T ordered] struct {
	Value T
	Right *Node[T]
	Left  *Node[T]
}

func newNode[T ordered](value T, right, left *Node[T]) *Node[T] {
	return &Node[T]{
		Value: value,
		Right: right,
		Left:  left,
	}
}

type BST[T ordered] struct {
	Root *Node[T]
}

func New[T ordered]() *BST[T] {
	return &BST[T]{}
}

func (bst *BST[T]) Insert(value T) *BST[T] {
	new_node := newNode(value, nil, nil)

	if bst.Root == nil {
		bst.Root = new_node
		return bst
	}

	node := bst.Root
	for node != nil {
		if node.Value < value {
			if node.Right == nil {
				node.Right = new_node
				break
			}
			node = node.Right
		} else if value < node.Value {
			if node.Left == nil {
				node.Left = new_node
				break
			}
			node = node.Left
		} else {
			return bst
		}
	}
	return bst
}

func (bst *BST[T]) RecursiveInsert(value T) *BST[T] {
	new_node := newNode(value, nil, nil)

	if bst.Root == nil {
		bst.Root = new_node
		return bst
	}

	recInsert(bst.Root, new_node)
	return bst
}

func recInsert[T ordered](check_node *Node[T], insert_node *Node[T]) {
	if check_node.Value < insert_node.Value {
		if check_node.Right != nil {
			recInsert(check_node.Right, insert_node)
		} else {
			check_node.Right = insert_node
		}
	} else {
		if check_node.Left != nil {
			recInsert(check_node.Left, insert_node)
		} else {
			check_node.Left = insert_node
		}
	}
}

func (bst *BST[T]) Find(value T) *Node[T] {
	if bst.Root == nil {
		return nil
	}

	for node := bst.Root; node != nil; {
		if node.Value < value {
			node = node.Right
		} else if value < node.Value {
			node = node.Left
		} else {
			return node
		}
	}
	return nil
}

func (bst *BST[T]) Traverse(method TraverseType) []T {
	if method == "bfs" {
		return breathFirstSearch(bst)
	} else if method == "dfs" {
		return depthFirstSearch(bst)
	}
	return nil
}

func breathFirstSearch[T ordered](bst *BST[T]) []T {
	var node *Node[T]
	visited := []T{}
	queue := queue.New[*Node[T]]()
	queue.Push(bst.Root)
	for queue.Length != 0 {
		node = queue.Pop().Value
		visited = append(visited, node.Value)
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
	}
	return visited
}

func depthFirstSearch[T ordered](bst *BST[T]) []T {
	return lookingRecursive(bst.Root)
}

func lookingRecursive[T ordered](node *Node[T]) (visited []T) {
	if node.Left != nil {
		new_visited := lookingRecursive(node.Left)
		visited = append(visited, new_visited...)
	}
	// Moving this line will change
	visited = append(visited, node.Value)
	if node.Right != nil {
		new_visited := lookingRecursive(node.Right)
		visited = append(visited, new_visited...)
	}
	return visited
}
