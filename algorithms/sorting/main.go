package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iDeepesh/goNotebook/algorithms/sorting/sorts"
)

func main() {
	fmt.Println("Inside sorting main.")

	var a []int
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter a number: ")
	for s.Scan() {
		t := s.Text()
		if strings.Compare(t, "sort") == 0 {
			break
		}
		n, e := strconv.Atoi(t)
		if e != nil {
			fmt.Println("Please enter valid integer value !!!!")
		} else {
			a = append(a, n)
		}
		fmt.Print("Enter another number or sort to start sorting: ")
	}

	fmt.Print("Which sorting - b (bubble), s (selection), i (insertion), m (merge), q (quick), h (heap): ")
	s.Scan()

	switch t := s.Text(); t {
	case "b":
		sorts.BubbleSort(a)
		return
	case "s":
		sorts.SelectionSort(a)
		return
	case "i":
		sorts.InsertionSort(a)
		return
	case "m":
		sorts.MergeSort(a)
		return
	case "q":
		sorts.QuickSort(a)
		return
	case "h":
		sorts.HeapSort(a)
		return
	}

	fmt.Println("Error - Unsupported sort requested........")
}
