package btree

type BTree struct {
	degree int
	root   *Node
}

type Node struct {
	keys map[int]*Node
	leaf bool
}

func New(degree int) *BTree {
	return &BTree{
		root:   nil,
		degree: degree,
	}
}
