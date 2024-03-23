package priorityqueue

import "fmt"

type Node[T any, P ~int | ~uint | float64] struct {
	Value    T
	Priority P
}

type PriorityQueue[T any, P ~int | ~uint | float64] struct {
	values []*Node[T, P]
}

func New[T any, P ~int | ~uint | float64]() *PriorityQueue[T, P] {
	return &PriorityQueue[T, P]{
		values: []*Node[T, P]{},
	}
}

func (pq *PriorityQueue[T, P]) Insert(value T, priority P) *PriorityQueue[T, P] {
	current_idx := len(pq.values)
	pq.values = append(pq.values, &Node[T, P]{
		Priority: priority,
		Value:    value,
	})
	parent_idx := (current_idx - 1) / 2

	// Bubble Up
	for pq.values[parent_idx].Priority > priority {
		pq.values[parent_idx], pq.values[current_idx] = pq.values[current_idx], pq.values[parent_idx]
		current_idx = parent_idx
		parent_idx = (current_idx - 1) / 2
	}
	return pq
}

func (pq *PriorityQueue[T, P]) Remove() (priority P, value T, ok bool) {
	if len(pq.values) == 0 {
		return 0, *new(T), false
	}
	node := pq.values[0]
	last_index := len(pq.values) - 1
	pq.values[0] = pq.values[last_index]
	pq.values = pq.values[:last_index]

	// Sink down
	n := 0
	next_idx := pq.nextIndex(n)
	for next_idx != 0 && pq.values[n].Priority > pq.values[next_idx].Priority {
		pq.values[n], pq.values[next_idx] = pq.values[next_idx], pq.values[n]
		n = next_idx
		next_idx = pq.nextIndex(n)
	}

	return node.Priority, node.Value, true
}

// Returns the index of the maximum of the next row.
// If the current node is a leaf, returns 0
func (pq *PriorityQueue[T, P]) nextIndex(n int) (idx int) {
	idx_left := 2*n + 1
	idx_right := idx_left + 1
	if idx_right >= len(pq.values) {
		if idx_left >= len(pq.values) {
			return 0
		}
		return idx_left
	}
	if pq.values[idx_left].Priority < pq.values[idx_right].Priority {
		return idx_left
	}
	return idx_right
}

func (pq *PriorityQueue[T, P]) String() string {
	values := []Node[T, P]{}
	for _, item := range pq.values {
		values = append(values, *item)
	}
	return fmt.Sprint(values)
}
