package heap

import (
	"fmt"
)

type GenericBinaryHeap[T any] struct {
	cmpVar  int
	compare CompareFunc[T]
	values  []T
}

// Makes a min binary heap by default, but it can be converted to max binary heap if we pass in isMax = True argument
func NewGenericBinaryHeap[T any](compareFunc CompareFunc[T], isMax bool) *GenericBinaryHeap[T] {
	cmpVar := 1
	if isMax {
		cmpVar = -1
	}

	return &GenericBinaryHeap[T]{
		cmpVar:  cmpVar,
		compare: compareFunc,
		values:  make([]T, 0),
	}
}

func (bh *GenericBinaryHeap[T]) Push(value T) *GenericBinaryHeap[T] {
	current_idx := len(bh.values)
	bh.values = append(bh.values, value)
	parent_idx := (current_idx - 1) / 2

	// Bubble Up
	// In case of MinBinaryHeap, bh.cmpVar is 1 and shows a greater than b (value from argument)
	for bh.compare(bh.values[parent_idx], value) == bh.cmpVar {
		bh.values[parent_idx], bh.values[current_idx] = bh.values[current_idx], bh.values[parent_idx]
		current_idx = parent_idx
		parent_idx = (current_idx - 1) / 2
	}

	return bh
}

func (bh *GenericBinaryHeap[T]) Pop() (value T, ok bool) {
	if len(bh.values) == 0 {
		return value, false
	}
	value = bh.values[0]
	last_index := len(bh.values) - 1
	bh.values[0] = bh.values[last_index]
	bh.values = bh.values[:last_index]

	// Sink down
	n := 0
	next_idx := bh.nextIndex(n)
	for next_idx != 0 && bh.compare(bh.values[n], bh.values[next_idx]) == bh.cmpVar {
		bh.values[n], bh.values[next_idx] = bh.values[next_idx], bh.values[n]
		n = next_idx
		next_idx = bh.nextIndex(n)
	}

	return value, true
}

// Returns the index of the maximum of the next row.
// If the current node is a leaf, returns 0
func (bh *GenericBinaryHeap[T]) nextIndex(n int) (idx int) {
	idx_left := 2*n + 1
	idx_right := idx_left + 1
	if idx_right >= len(bh.values) {
		if idx_left >= len(bh.values) {
			return 0
		}
		return idx_left
	}
	if bh.compare(bh.values[idx_left], bh.values[idx_right]) == -bh.cmpVar {
		return idx_left
	}
	return idx_right
}

func (bh *GenericBinaryHeap[T]) String() string {
	return fmt.Sprint(bh.values)
}

func (bh *GenericBinaryHeap[T]) Len() int {
	return len(bh.values)
}
