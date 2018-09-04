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

		fmt.Print("Please enter the rank index of the element to select: ")
		s.Scan()
		rank, e := strconv.Atoi(s.Text())
		if e != nil {
			fmt.Println("Please enter valid integer value !!!!")
		}

		result := quickSelect(a, rank)
		fmt.Printf("\nThe element at rank %d will be %d\n\n", rank, result)

		sorts.QuickSort(a)

		fmt.Print("Do you want to try another selection, y (yes) or n (no)?")
		s.Scan()
		if strings.Compare(s.Text(), "n") == 0 {
			break
		}
	}
}

func quickSelect(a []int, rank int) int {
	l := len(a) - 1

	c, i := 0, 0
	for i < l {
		if a[i] < a[l] {
			a[c], a[i] = a[i], a[c]
			i++
			c++
		} else {
			i++
		}
	}
	a[c], a[l] = a[l], a[c]

	if c == rank {
		return a[c]
	}

	if c > rank {
		return quickSelect(a[:c], rank)
	} else {
		return quickSelect(a[c+1:], rank-c-1)
	}
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
