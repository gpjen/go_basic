package main

import "fmt"

type Stack struct {
	Arr      []uint
	Capacity int
	Length   int
}

func NewStack(cap int) *Stack {
	return &Stack{
		Arr:      make([]uint, cap),
		Capacity: cap,
		Length:   0,
	}
}

func (s *Stack) Push(val uint) {
	if s.Length == s.Capacity {
		fmt.Println("Stack Overflow...")
		return
	}

	s.Arr[s.Length] = val
	s.Length++
}

func (s *Stack) Pop() {
	if s.Length == 0 {
		fmt.Println("Stack is empty...")
		return
	}

	s.Length--
	popped := s.Arr[s.Length]
	s.Arr[s.Length] = 0

	fmt.Println("Popped :", popped)

}

func (s *Stack) Print() {

	fmt.Println("Length :", s.Length)

	for idx, value := range s.Arr {
		fmt.Println(idx, " :", value)
	}
}

func main() {
	fmt.Println("---- Stack -----")

	mystack := NewStack(5)

	mystack.Push(10)
	mystack.Push(20)
	mystack.Push(30)

	mystack.Pop()
	mystack.Pop()
	mystack.Pop()
	mystack.Pop()
	mystack.Pop()
	mystack.Pop()

	mystack.Push(40)
	mystack.Push(50)
	mystack.Push(60)

	mystack.Print()
}
