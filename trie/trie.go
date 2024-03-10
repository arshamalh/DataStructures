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
	currentNode.words["*"] = nil
}

// TODO: It's a search prefix, or if isn't, write the complete one
func (root *trie) Search(word string) (string, error) {
	currentNode := root
	result := ""
	for _, charRune := range word {
		char := string(charRune)
		if val, ok := currentNode.words[char]; ok { // TODO: nil check on current node?
			result += char
			currentNode = val
		} else {
			return "", ErrWordDoesNotExist
		}
	}
	return result, nil
}

func (root *trie) searchPrefix(prefix string) *trie {
	currentNode := root
	for _, charRune := range prefix {
		char := string(charRune)
		if val, ok := currentNode.words[char]; ok { // TODO: nil check on current node?
			currentNode = val
		} else {
			break
		}
	}
	return currentNode
}

func (root *trie) GetAllTheWords(prefix string) []string {
	listOfWords := make([]string, 0)
	if root == nil {
		return listOfWords
	}
	for charKey, childTrie := range root.words {
		childWords := childTrie.GetAllTheWords(charKey)
		if len(childWords) == 0 {
			listOfWords = append(listOfWords, prefix)
			continue
		}

		for _, word := range childWords {
			if word == "*" {
				word = ""
			}
			listOfWords = append(listOfWords, prefix+word)
		}
	}
	return listOfWords
}

func (t *trie) AutoComplete(prefix string) ([]string, error) {
	if prefix == "" { // There is nothing to return for empty prefix but there is no error too.
		return []string{}, nil
	}

	currentNode := t.searchPrefix(prefix)
	return currentNode.GetAllTheWords(prefix), nil
}
