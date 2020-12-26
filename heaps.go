package main

// MaxHeap struct has a slice that holds an array
type MaxHeap struct {
	array []int
}

// Insert add an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key and removes it from the heap

func main() {

}
