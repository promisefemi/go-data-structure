package main

import "fmt"

// Node represents the components of the binary search
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
func (n *Node) Insert(k int) {
	if n.Key < k {
		//  Move Right
		if n.Right == nil {
			rightNode := &Node{Key: k}
			n.Right = rightNode
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// Move left
		if n.Left == nil {
			leftNode := &Node{Key: k}
			n.Left = leftNode
		} else {
			n.Left.Insert(k)
		}
	}

}

// Search
func (n *Node) Search(k int) bool {

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return false

}

func main() {
	tree := &Node{Key: 50}
	fmt.Println(tree)

}
