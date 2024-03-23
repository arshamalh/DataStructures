package heap

import (
	"fmt"
)

type MaxBinaryHeap struct {
	values []int
}

func NewMaxBinaryHeap() *MaxBinaryHeap {
	return &MaxBinaryHeap{
		values: []int{},
	}
}

func (bh *MaxBinaryHeap) Insert(value int) *MaxBinaryHeap {
	current_idx := len(bh.values)
	bh.values = append(bh.values, value)
	parent_idx := (current_idx - 1) / 2

	// Bubble Up
	for bh.values[parent_idx] < value {
		bh.values[parent_idx], bh.values[current_idx] = bh.values[current_idx], bh.values[parent_idx]
		current_idx = parent_idx
		parent_idx = (current_idx - 1) / 2
	}
	return bh
}

func (bh *MaxBinaryHeap) Remove() (value int) {
	value = bh.values[0]
	last_index := len(bh.values) - 1
	bh.values[0] = bh.values[last_index]
	bh.values = bh.values[:last_index]

	// Sink down
	n := 0
	next_idx := bh.nextIndex(n)
	for next_idx != 0 && bh.values[n] < bh.values[next_idx] {
		bh.values[n], bh.values[next_idx] = bh.values[next_idx], bh.values[n]
		n = next_idx
		next_idx = bh.nextIndex(n)
	}

	return value
}

// Returns the index of the maximum of the next row.
// If the current node is a leaf, returns 0
func (bh *MaxBinaryHeap) nextIndex(n int) (idx int) {
	idx_left := 2*n + 1
	idx_right := idx_left + 1
	if idx_right >= len(bh.values) {
		if idx_left >= len(bh.values) {
			return 0
		}
		return idx_left
	}
	if bh.values[idx_left] > bh.values[idx_right] {
		return idx_left
	}
	return idx_right
}

func (bh *MaxBinaryHeap) String() string {
	return fmt.Sprint(bh.values)
}
