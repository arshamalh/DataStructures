package trie

import "fmt"

type trie struct {
	words map[string]*trie
}

var ErrWordDoesNotExist = fmt.Errorf("word does not exist")

func New() *trie {
	return &trie{
		words: make(map[string]*trie),
	}
}

func (root *trie) Insert(word string) {
	currentNode := root
	for _, charRune := range word {
		char := string(charRune)
		if val, ok := currentNode.words[char]; ok {
			currentNode = val
		} else {
			newNode := New()
			currentNode.words[char] = newNode
			currentNode = newNode
		}
	}
}

func (root *trie) Search(word string) (string, error) {
	currentNode := root
	result := ""
	for _, charRune := range word {
		char := string(charRune)
		if val, ok := currentNode.words[char]; ok {
			result += char
			currentNode = val
		} else {
			return "", ErrWordDoesNotExist
		}
	}
	return result, nil
}
