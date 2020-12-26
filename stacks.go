package main

import "fmt"

type Stack struct {
	items []int
}

// Push will add value
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove value
func (s *Stack) Pop() int {
	value := s.items[len(s.items)-1]
	s.items = s.items[1 : len(s.items)-1]
	return value
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(3)
	myStack.Push(10)
	myStack.Push(20)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)

}
