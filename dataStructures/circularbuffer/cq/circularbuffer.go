package cq

import (
	"fmt"
)

//CircularBuffer - data structure representing a circular buffer
type CircularBuffer struct {
	buffer []int
	s, l   int
}

//CreateBuffer - creates and initializes a CircularBuffer of given length
func CreateBuffer(l int) CircularBuffer {
	return CircularBuffer{make([]int, l), 0, 0}
}

//PrintCQ - prints the circular buffer
func (buf CircularBuffer) PrintCQ() {
	fmt.Printf("Start index: %d, Length: %d, Capacity: %d\n", buf.s, buf.l, len(buf.buffer))
	fmt.Println(buf.buffer)
}

//Push - push an item into the buffer
func (buf *CircularBuffer) Push(n int) {
	if buf.l == len(buf.buffer) {
		fmt.Println("Circular Buffer is full")
		return
	}

	buf.buffer[(buf.s+buf.l)%len(buf.buffer)] = n
	buf.l++
}

//Pop - pop the item at the head of the buffer
func (buf *CircularBuffer) Pop() int {
	if buf.l <= 0 {
		fmt.Println("Circular Buffer is empty")
		return -1
	}

	r := buf.buffer[buf.s]
	buf.s = (buf.s + 1) % len(buf.buffer)

	buf.l--
	return r
}
