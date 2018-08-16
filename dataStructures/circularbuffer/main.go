package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/iDeepesh/goNotebook/dataStructures/circularbuffer/cq"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter the size of the circular buffer: ")
	s.Scan()
	l, _ := strconv.Atoi(s.Text())

	cb := cq.CreateBuffer(l)
	fmt.Println("Created circular buffer with length of ", l)

	for {
		fmt.Print("Please choose the operation - add (a), delete (d), quit (q): ")
		s.Scan()

		stop := false
		switch selection := s.Text(); selection {
		case "q":
			stop = true
		case "a":
			fmt.Print("Please enter the number to add to circular buffer:")
			s.Scan()
			n, _ := strconv.Atoi(s.Text())
			cb.Push(n)
		case "d":
			n := cb.Pop()
			fmt.Println("Here is the number removed from the head of the circular buffer", n)
		}

		cb.PrintCQ()
		if stop {
			break
		}
	}
}
