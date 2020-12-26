package main

import "fmt"

const arraySize = 7

// HashTable structure is an array of linked
type HashTable struct {
	array [arraySize]*bucket
}

// Bucket structure is a linked list
type bucket struct {
	head *bucketNode
}

// bucketNode is the nodes of the linked list
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert
func (h *HashTable) Insert(k string) {

	keyIndex := hashFunction(k)

	keyBucket := h.array[keyIndex]
	keyBucket.insert(k)

}

// Search
func (h *HashTable) Search(k string) bool {
	keyIndex := hashFunction(k)
	return h.array[keyIndex].search(k)
}

// // Delete
func (h *HashTable) Delete(k string) bool {
	keyIndex := hashFunction(k)
	h.array[keyIndex].delete(k)
}

// insert
func (b *bucket) insert(k string) {

	if b.search(k) {
		fmt.Println("Key already exists")
		return
	}
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

// search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) delete(k string) bool {
	currentNode := b.head

	if currentNode == nil {
		return false
	}

	if currentNode.key == k {
		b.head = b.head.next
		return true
	}

	for currentNode != nil {
		if currentNode.next.key == k {
			currentNode.next = currentNode.next.next
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

// hash
func hashFunction(k string) int {
	sum := 0

	for _, v := range k {
		sum += int(v)
	}
	return sum % arraySize
}

// Init will initialize a new hash table
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func main() {
	// myHash := Init()

	// fmt.Println(myHash)

	testBucket := &bucket{}
	testBucket.insert("PROMISE")
	testBucket.insert("RANDY")
	fmt.Println(testBucket.delete("RANDY"))
	fmt.Println(testBucket.search("RANDY"))
}
