package heap

import (
	"fmt"
)

// Limited Generic Binary Heap
type LGBinaryHeap[T any] struct {
	// cmpVar will be 1 for min heap
	cmpVar    int
	compare   CompareFunc[T]
	values    []T
	maxLength int
}

// Makes a min binary heap by default, but it can be converted to max binary heap if we pass in isMax = True argument
func NewLGBinaryHeap[T any](compareFunc CompareFunc[T], maxLength int, isMax bool) *LGBinaryHeap[T] {
	cmpVar := 1
	if isMax {
		cmpVar = -1
	}

	if maxLength < 1 {
		maxLength = 1
	}

	return &LGBinaryHeap[T]{
		cmpVar:    cmpVar,
		compare:   compareFunc,
		values:    make([]T, 0),
		maxLength: maxLength,
	}
}

func (bh *LGBinaryHeap[T]) Push(value T) *LGBinaryHeap[T] {
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

	if bh.Len() > bh.maxLength {
		bh.ReversePop()
	}

	return bh
}

func (bh *LGBinaryHeap[T]) Pop() (value T, ok bool) {
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

// Removes and returns max element from MinHeap.
// Removes and return min element from MaxHeap
func (bh *LGBinaryHeap[T]) ReversePop() (value T, ok bool) {
	if len(bh.values) == 0 {
		return value, false
	}
	idx := len(bh.values) / 2
	endValue := bh.values[idx]
	for i, val := range bh.values {
		if bh.compare(val, endValue) == bh.cmpVar {
			endValue = val
			idx = i
		}
	}
	bh.values = append(bh.values[:idx], bh.values[idx+1:]...)
	return endValue, true
}

// Returns the index of the minimum (maximum) of the next row.
// If the current node is a leaf, returns 0
func (bh *LGBinaryHeap[T]) nextIndex(n int) (idx int) {
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

func (bh *LGBinaryHeap[T]) String() string {
	return fmt.Sprint(bh.values)
}

func (bh *LGBinaryHeap[T]) Len() int {
	return len(bh.values)
}
