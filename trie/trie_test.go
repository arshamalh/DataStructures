package trie_test

import (
	triePkg "ds/trie"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		title         string
		word          string
		expectedError error
	}{
		{
			title:         "empty strings",
			word:          "",
			expectedError: nil,
		},
		{
			title:         "valid string",
			word:          "Arsham",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			trie := triePkg.New()
			trie.Insert(tc.word)
			word, err := trie.Search(tc.word)
			assert.Equal(tc.expectedError, err)
			assert.Equal(tc.word, word)
		})
	}
}

const (
	Arthas   = "Arthas"
	Arsham   = "Arsham"
	Atousa   = "Atousa"
	Banana   = "Banana"
	Babak    = "Babak"
	Baba     = "Baba"     // Less chars than babak
	BabakAli = "BabakAli" // More chars than babak
)

func TestTrieMultipleWords(t *testing.T) {
	// Different orders as order should not matter
	namesInsert := []string{Arsham, Babak, Atousa, Arthas}
	namesSearch := []string{Arthas, Babak, Arsham, Atousa}
	assert := assert.New(t)
	trie := triePkg.New()

	for _, name := range namesInsert {
		trie.Insert(name)
	}

	for _, name := range namesSearch {
		found, err := trie.Search(name)
		assert.Nil(err)
		assert.Equal(name, found)
	}

	// "Baba" doesn't exist but it's a prefix of Babak,
	// so as we're just looking for prefixes,
	// it should not throw an error
	_, err := trie.Search(Baba)
	assert.Nil(err)

	_, err = trie.Search(Banana)
	assert.Equal(triePkg.ErrWordDoesNotExist, err)

	_, err = trie.Search(BabakAli)
	assert.Equal(triePkg.ErrWordDoesNotExist, err)
}
