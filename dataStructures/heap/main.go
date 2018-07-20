package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/iDeepesh/goNotebook/dataStructures/heap/minheap"
)

func main() {
	h := minheap.MinHeap{}

	s := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter command to play with heap - add, delete, exit: ")
	for s.Scan() {
		t := s.Text()
		switch t {
		case "exit":
			fmt.Println("Final Heap: ", h)
			return
		case "add":
			h = addElem(h, s)
		case "delete":
			h = deleteElem(h, s)
		}
		fmt.Print("Please enter another command to play with heap - add, delete, exit: ")
	}

}

func addElem(h minheap.MinHeap, s *bufio.Scanner) minheap.MinHeap {
	fmt.Print("Please enter a valid integer: ")
	s.Scan()
	n, e := strconv.Atoi(s.Text())
	if e != nil {
		fmt.Println("Invalid integer entered!!!")
		return h
	}
	return h.Add(n)
}

func deleteElem(h minheap.MinHeap, s *bufio.Scanner) minheap.MinHeap {

	fmt.Println("Current heap: ", h)
	fmt.Print("Please enter a valid index to delete: ")
	s.Scan()
	i, e := strconv.Atoi(s.Text())
	if e != nil || i < 0 || i > len(h)-1 {
		fmt.Println("Invalid index entered!!!")
		return h
	}
	return h.Delete(i)
}
