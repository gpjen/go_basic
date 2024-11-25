package main

import "fmt"

type Queue struct {
	Arr    []string
	Cap    int
	Length int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		Arr:    make([]string, cap),
		Cap:    cap,
		Length: 0,
	}
}

func (q *Queue) Push(val string) {
	if val == "" {
		fmt.Println("Value is empty...")
		return
	}

	if q.Length == q.Cap {
		fmt.Println("Queue is full...")
		return
	}

	q.Arr[q.Length] = val
	q.Length++
}

func (q *Queue) Pop() {
	if q.Length == 0 {
		fmt.Println("Queue is empty..")
		return
	}

	popped := q.Arr[0]
	fmt.Println("Popped :", popped)

	// swich element
	for i := 1; i < q.Length; i++ {
		q.Arr[i-1] = q.Arr[i]
	}

	q.Length--
	q.Arr[q.Length] = ""

}

func (q *Queue) Print() {
	for idx, el := range q.Arr {
		fmt.Println(idx+1, ":", el)
	}
	fmt.Printf("Length : %d, Cap : %d/%d", q.Length, q.Cap, len(q.Arr))
}

func main() {
	fmt.Println("----- Queue -----")

	queue := NewQueue(3)
	queue.Push("10")
	queue.Push("20")
	queue.Push("30")
	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Push("40")
	queue.Push("50")
	queue.Push("60")
	queue.Push("70")
	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Push("80")
	queue.Print()
}
