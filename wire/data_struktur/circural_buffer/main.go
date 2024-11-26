package main

import "fmt"

type circularBuffer struct {
	arr        []uint
	readIndex  int
	writeIndex int
	capacity   int
	length     int
}

func NewCircularBuffer(cap int) *circularBuffer {
	return &circularBuffer{
		arr:        make([]uint, cap),
		capacity:   cap,
		readIndex:  0,
		writeIndex: 0,
		length:     0,
	}
}

func (c *circularBuffer) Push(value uint) {
	if c.length == c.capacity {
		fmt.Println("Circular buffer is full....")
		return
	}

	c.arr[c.writeIndex] = value
	c.writeIndex = (c.writeIndex + 1) % c.capacity
	c.length++
}

func (c *circularBuffer) Pop() {
	if c.length == 0 {
		fmt.Println("Circular buffer is empty...")
		return
	}

	c.arr[c.readIndex] = 0
	c.readIndex = (c.readIndex + 1) % c.capacity
	c.length--
}

func (c *circularBuffer) Print() {

	fmt.Println("length :", c.length)
	for idx, val := range c.arr {
		fmt.Println(idx, " :", val)
	}
}

func main() {
	fmt.Println("----- Circular Buffer ----")

	mycb := NewCircularBuffer(5)

	mycb.Push(10)
	mycb.Push(20)
	mycb.Push(30)
	mycb.Push(40)
	mycb.Push(50)
	mycb.Push(51)
	mycb.Pop()
	mycb.Pop()
	mycb.Pop()
	mycb.Pop()
	mycb.Pop()
	mycb.Pop()
	mycb.Push(10)
	mycb.Push(20)
	mycb.Push(30)

	mycb.Print()
}
