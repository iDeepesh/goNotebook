package interfaces

import (
	"container/heap"
	"fmt"
)

type minHeap []int
type maxHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[0 : l-1]
	return x
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	// we are returning more than because it is min heap
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[0 : l-1]
	return x
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

//ExecuteMinHeapTest - as the name suggests
func ExecuteMinHeapTest() {
	fmt.Println("Inside ExecuteMinHeapTest")
	min := minHeap{54, 38, 72, 98, 15, 8, 84, 21, 46, 63}
	heap.Init(&min)
	add := []int{50, 10, 40, 70, 90, 30, 80, 20, 60}
	fmt.Println(min)
	for _, a := range add {
		heap.Push(&min, a)
		fmt.Printf("Added: %d, New Heap: %d\n", a, min)
	}

	l := len(min)
	for i := 0; i < l; i++ {
		fmt.Printf("Popped: %d, New heap: ", heap.Pop(&min))
		fmt.Println(min)
	}
}

//ExecuteMaxHeapTest - as the name suggests
func ExecuteMaxHeapTest() {
	fmt.Println("Inside ExecuteMaxHeapTest")
	m := maxHeap{54, 38, 72, 98, 15, 8, 84, 21, 46, 63}
	heap.Init(&m)
	add := []int{50, 10, 40, 70, 90, 30, 80, 20, 60}
	fmt.Println(m)
	for _, a := range add {
		heap.Push(&m, a)
		fmt.Printf("Added: %d, New Heap: %d\n", a, m)
	}

	l := len(m)
	for i := 0; i < l; i++ {
		fmt.Printf("Popped: %d, New heap: ", heap.Pop(&m))
		fmt.Println(m)
	}
}
