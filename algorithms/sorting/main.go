package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/iDeepesh/goNotebook/algorithms/sorting/sorts"
)

func main() {
	fmt.Println("Inside sorting main.")
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Please enter number of elements in array: ")
		s.Scan()
		n, e := strconv.Atoi(s.Text())
		if e != nil {
			fmt.Println("Please enter valid integer value !!!!")
		}
		fmt.Print("Please enter max number of digits for elements: ")
		s.Scan()
		m, e := strconv.Atoi(s.Text())
		if e != nil {
			fmt.Println("Please enter valid integer value !!!!")
		}

		a := generateIntegerArray(n, int(math.Pow10(m)))

		fmt.Print("Which sorting - b (bubble), s (selection), i (insertion), m (merge), q (quick), h (heap), r (radix): ")
		s.Scan()

		sort(a, s.Text(), m)

		fmt.Print("Do you want to sort another array, y (yes), n (no)? ")
		s.Scan()
		if strings.Compare(s.Text(), "n") == 0 {
			break
		}
	}
}

func sort(a []int, t string, m int) {

	switch t {
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
	case "r":
		sorts.RadixSort(a, m)
		return
	}

	fmt.Println("Error - Unsupported sort requested........")
}

func generateIntegerArray(n, m int) []int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	a := make([]int, n)
	for n > 0 {
		n--
		a[n] = r.Intn(m - 1)
	}

	fmt.Println(a)
	return a
}
