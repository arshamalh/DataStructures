package heap_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/arshamalh/DataStructures/heap"
	"github.com/stretchr/testify/assert"
)

type Article struct {
	Title     string
	CreatedAt time.Time
}

func CompareArticles(a, b Article) int {
	if a.CreatedAt.After(b.CreatedAt) {
		return 1
	} else if a.CreatedAt.Before(b.CreatedAt) {
		return -1
	}
	return 0
}

var listOfArticles = []Article{
	{"Modern generation", time.Date(2003, 10, 3, 0, 0, 0, 0, time.UTC)},
	{"AI taking over the world", time.Date(2022, 10, 9, 0, 0, 0, 0, time.UTC)},
	{"A cow in the wild", time.Date(1999, 10, 1, 0, 0, 0, 0, time.UTC)},
	{"Happiness can be found", time.Date(2012, 10, 7, 0, 0, 0, 0, time.UTC)},
}

func TestGMaxBinaryHeap(t *testing.T) {
	assert := assert.New(t)

	expectedOutput := []Article{
		{"AI taking over the world", time.Date(2022, 10, 9, 0, 0, 0, 0, time.UTC)},
		{"Happiness can be found", time.Date(2012, 10, 7, 0, 0, 0, 0, time.UTC)},
		{"Modern generation", time.Date(2003, 10, 3, 0, 0, 0, 0, time.UTC)},
		{"A cow in the wild", time.Date(1999, 10, 1, 0, 0, 0, 0, time.UTC)},
	}

	bh := heap.NewGMaxBinaryHeap(CompareArticles)
	for _, n := range listOfArticles {
		bh.Push(n)
	}

	got := make([]Article, 0)
	for {
		gotArticle, ok := bh.Pop()
		if !ok {
			break
		}
		got = append(got, gotArticle)
	}

	assert.Equal(expectedOutput, got)
}

func TestGMinBinaryHeap(t *testing.T) {
	assert := assert.New(t)

	expectedOutput := []Article{
		{"A cow in the wild", time.Date(1999, 10, 1, 0, 0, 0, 0, time.UTC)},
		{"Modern generation", time.Date(2003, 10, 3, 0, 0, 0, 0, time.UTC)},
		{"Happiness can be found", time.Date(2012, 10, 7, 0, 0, 0, 0, time.UTC)},
		{"AI taking over the world", time.Date(2022, 10, 9, 0, 0, 0, 0, time.UTC)},
	}

	bh := heap.NewGMinBinaryHeap(CompareArticles)
	for _, n := range listOfArticles {
		bh.Push(n)
	}

	got := make([]Article, 0)
	for {
		gotArticle, ok := bh.Pop()
		if !ok {
			break
		}
		got = append(got, gotArticle)
	}

	assert.Equal(expectedOutput, got)
}

func TestGenericBinaryHeap(t *testing.T) {
	assert := assert.New(t)

	expectedOutputForMin := []Article{
		{"A cow in the wild", time.Date(1999, 10, 1, 0, 0, 0, 0, time.UTC)},
		{"Modern generation", time.Date(2003, 10, 3, 0, 0, 0, 0, time.UTC)},
		{"Happiness can be found", time.Date(2012, 10, 7, 0, 0, 0, 0, time.UTC)},
		{"AI taking over the world", time.Date(2022, 10, 9, 0, 0, 0, 0, time.UTC)},
	}

	expectedOutputForMax := []Article{
		{"AI taking over the world", time.Date(2022, 10, 9, 0, 0, 0, 0, time.UTC)},
		{"Happiness can be found", time.Date(2012, 10, 7, 0, 0, 0, 0, time.UTC)},
		{"Modern generation", time.Date(2003, 10, 3, 0, 0, 0, 0, time.UTC)},
		{"A cow in the wild", time.Date(1999, 10, 1, 0, 0, 0, 0, time.UTC)},
	}

	// Min Binary Heap
	minBH := heap.NewGenericBinaryHeap(CompareArticles, false)
	for _, n := range listOfArticles {
		minBH.Push(n)
	}

	gotMin := make([]Article, 0)
	for {
		gotArticle, ok := minBH.Pop()
		if !ok {
			break
		}
		gotMin = append(gotMin, gotArticle)
	}

	assert.Equal(expectedOutputForMin, gotMin)

	// Max Binary Heap
	maxBH := heap.NewGenericBinaryHeap(CompareArticles, true)
	for _, n := range listOfArticles {
		maxBH.Push(n)
	}

	gotMax := make([]Article, 0)
	for {
		gotArticle, ok := maxBH.Pop()
		if !ok {
			break
		}
		gotMax = append(gotMax, gotArticle)
	}

	assert.Equal(expectedOutputForMax, gotMax)
}

func TestLimitedGenericBinaryHeapMin(t *testing.T) {
	assert := assert.New(t)

	expectedOutputForMin := []Article{
		{"A cow in the wild", time.Date(1999, 10, 1, 0, 0, 0, 0, time.UTC)},
		{"Modern generation", time.Date(2003, 10, 3, 0, 0, 0, 0, time.UTC)},
	}

	// Min Binary Heap
	minBH := heap.NewLGBinaryHeap(CompareArticles, 2, false)
	for _, n := range listOfArticles {
		minBH.Push(n)
		fmt.Println(">>", minBH.String())

	}

	gotMin := make([]Article, 0)
	for {
		gotArticle, ok := minBH.Pop()
		if !ok {
			break
		}
		gotMin = append(gotMin, gotArticle)
	}

	for i, itemGotMax := range gotMin {
		assert.Equal(itemGotMax.CreatedAt.Year(), expectedOutputForMin[i].CreatedAt.Year())
		assert.Equal(itemGotMax.Title, expectedOutputForMin[i].Title)
	}
}

func TestLimitedGenericBinaryHeapMax(t *testing.T) {
	assert := assert.New(t)

	expectedOutputForMax := []Article{
		{"AI taking over the world", time.Date(2022, 10, 9, 0, 0, 0, 0, time.UTC)},
		{"Happiness can be found", time.Date(2012, 10, 7, 0, 0, 0, 0, time.UTC)},
	}

	// Max Binary Heap
	maxBH := heap.NewLGBinaryHeap(CompareArticles, 2, true)
	for _, n := range listOfArticles {
		maxBH.Push(n)
		fmt.Println(">>", maxBH.String())
	}

	gotMax := make([]Article, 0)
	for {
		gotArticle, ok := maxBH.Pop()
		if !ok {
			break
		}
		gotMax = append(gotMax, gotArticle)
	}

	for i, itemGotMax := range gotMax {
		assert.Equal(itemGotMax.CreatedAt.Year(), expectedOutputForMax[i].CreatedAt.Year())
		assert.Equal(itemGotMax.Title, expectedOutputForMax[i].Title)
	}
}
