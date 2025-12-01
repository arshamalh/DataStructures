package priorityqueue_test

import (
	"testing"

	"github.com/arshamalh/DataStructures/priorityqueue"
	"github.com/stretchr/testify/assert"
)

func TestMinPQ(t *testing.T) {
	assert := assert.New(t)

	type Node struct {
		Name     string
		Priority int
	}
	input := []Node{
		{"Bizhan", 5},
		{"Arsham", 1},
		{"Abtin", 4},
		{"Delaram", 3},
		{"Atash", 2},
	}
	expectedOutput := []Node{
		{"Arsham", 1},
		{"Atash", 2},
		{"Delaram", 3},
		{"Abtin", 4},
		{"Bizhan", 5},
	}

	mpq := priorityqueue.New[string, int]()
	for _, person := range input {
		mpq.Insert(person.Name, person.Priority)
	}

	got := make([]Node, 0)
	for {
		priority, name, ok := mpq.Remove()
		if !ok {
			break
		}
		got = append(got, Node{
			Name:     name,
			Priority: priority,
		})
	}

	assert.Equal(expectedOutput, got)
}
