package main

import "fmt"

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue add to the queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)

}

// Dequeue removes from a queue
func (q *Queue) Dequeue() int {
	value := q.items[0]
	q.items = q.items[1:len(q.items)]
	return value
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(3)
	myQueue.Enqueue(44)
	myQueue.Enqueue(52)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
