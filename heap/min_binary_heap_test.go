package heap_test

import (
	"testing"

	"github.com/arshamalh/DataStructures/heap"
	"github.com/stretchr/testify/assert"
)

func TestMinBinaryHeap(t *testing.T) {
	assert := assert.New(t)
	input := []int{5, 1, 6, 2, 4, 3}
	expectedOutput := []int{1, 2, 3, 4, 5, 6}
	bh := heap.NewMinBinaryHeap()
	for _, n := range input {
		bh.Push(n)
	}

	got := make([]int, 0)
	for {
		gotNumber, ok := bh.Pop()
		if !ok {
			break
		}
		got = append(got, gotNumber)
	}

	assert.Equal(expectedOutput, got)
}

func TestMinBinaryHeapLimited(t *testing.T) {
	assert := assert.New(t)
	input := []int{5, 1, 6, 7, 2, 9, 4, 3}
	expectedOutput := []int{1, 2, 3, 4, 5, 6}
	bh := heap.NewMinBinaryHeap()
	for _, n := range input {
		bh.Push(n)
	}

	got := make([]int, 0)
	for {
		gotNumber, ok := bh.Pop()
		if !ok {
			break
		}
		got = append(got, gotNumber)
	}

	assert.Equal(expectedOutput, got)
}
