package heap_test

import (
	"testing"

	"github.com/arshamalh/DataStructures/heap"
	"github.com/stretchr/testify/assert"
)

func TestMaxBinaryHeap(t *testing.T) {
	assert := assert.New(t)
	input := []int{5, 1, 6, 2, 4, 3}
	expectedOutput := []int{6, 5, 4, 3, 2, 1}
	bh := heap.NewMaxBinaryHeap()
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
