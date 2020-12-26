package main

import (
	"fmt"
	"strings"
)

// AlphabetSize is the number of the Alphabet
const AlphabetSize = 26

// Node represents a node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents the trie and it has the root Node
type Trie struct {
	root *Node
}

// InitTrie to create a new Trie
func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert will create new node entries in the trie
func (t *Trie) Insert(w string) {
	// convert words to lowercase to avoid running out of the array size
	w = strings.ToLower(w)
	// get the length of the word
	wordLength := len(w)
	// set currentNode to the root to allow looping through the children and consistently change node
	currentNode := t.root
	// begin for loop
	for i := 0; i < wordLength; i++ {
		// get the index position of the character by minusing it from the rune 'a'
		charIndex := w[i] - 'a'
		// verify the value of the currentIndex if nil it means it is an empty index, so create a new node in the index
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		// change the currentNode to the current child node and change isEnd to false to tell that it is not a leaf
		currentNode = currentNode.children[charIndex]
		currentNode.isEnd = false
	}

	// set isEnd to true for the current node
	currentNode.isEnd = true
}

// Search will find a word in the Trie
func (t *Trie) Search(w string) bool {
	// convert words to lowercase to avoid running out of the array size
	w = strings.ToLower(w)
	// get the length of the word
	wordLength := len(w)
	// set currentNode to the root to allow looping through the children and consistently change node
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		// get the index position of the character by minusing it from the rune 'a'
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			// if the current index is empty that means the word does not exist to return false
			return false
		}
		// change the currentNode to the children of the node, this will only happen if the current index of the currentNode child is not nil
		currentNode = currentNode.children[charIndex]
	}

	// if there was not empty child then the word exists in the trie and i am not checking if it isEnd because if a word is in the trie but still has children it should still be said to be inside the trie
	return true
}

// SearchReturn will search for a word and return the last node of that word
func (t *Trie) SearchReturn(w string) (bool, *Node) {
	// The same as Search but it returns the last currentNode
	w = strings.ToLower(w)
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false, nil
		}
		currentNode = currentNode.children[charIndex]
	}

	// currentNode.isEnd = true
	return true, currentNode
}

// FindSimilar will help find words that  are similar to the word or that completes it
func (t *Trie) FindSimilar(w string) []string {
	// change word to lowercase
	w = strings.ToLower(w)
	// create a slice to hold all the final similar words with a length of 0 to avoid the empty space problem
	similar := make([]string, 0)
	// Search the trie to find the word
	isPrefix, currentNode := t.SearchReturn(w)

	// if the word exists then find its similars
	if isPrefix {
		// Find similar from this node down
		wordSlice := currentNode.FindSimilar()
		// loop through the word and merge the similars with the original words
		for _, word := range wordSlice {
			mergedWord := w + word
			similar = append(similar, mergedWord)
		}

	}
	return similar

}

// FindSimilar will help nodes find all children down to its leaves
/*
	The point of this function is to loop through the children,
	add to the slice, merge the characters of the children's children
	into one slice and return to the parent to merge with its own slice
*/
func (n *Node) FindSimilar() []string {
	// create a slice to hold the final arranged similar and set to length 0 to avoid the empty space problem
	globalSimilar := make([]string, 0)

	// Loop through the children of the node
	for index, node := range n.children {
		// check if the current node is not nill
		if node != nil {
			// get the index of the current node by adding it to the rune 'a'
			charIndex := index + 'a'
			// use the index to get which character is supposed to be in that position
			character := fmt.Sprintf("%c", charIndex)

			// check if the current node is the end if so push the character into the globalSimilar (that holds the final slice)
			if node.isEnd {
				globalSimilar = append(globalSimilar, character)
			} else {
				// if not find the similar of the current node
				wordSlice := node.FindSimilar()
				// loop through the similar of the current node merge it with the character and push to the globalSimilar
				for _, word := range wordSlice {
					mergedWord := character + word
					globalSimilar = append(globalSimilar, mergedWord)
				}
			}

		}
	}

	return globalSimilar
}

func main() {
	myTrie := InitTrie()

	wordsToAdd := []string{"john", "josh", "joshua", "jamil", "jonah", "jibola"}

	for _, word := range wordsToAdd {
		myTrie.Insert(word)
	}

	// wordsToSearch := []string{"john", "josh", "jibola", "Joshua", "joshua"}
	// for _, word := range wordsToSearch {
	// 	fmt.Println(myTrie.Search(word))
	// }

	fmt.Println(myTrie.FindSimilar("j"))
}
