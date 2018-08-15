package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the size of the array: ")
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	fmt.Print("Please enter the maximum allowed value of elements: ")
	s.Scan()
	m, _ := strconv.Atoi(s.Text())

	a := generateArray(n, m)
	fmt.Println(a)
	i, max := splitArrayWithLargestDiff(a)
	fmt.Println("Maximum difference of sum for sub arrays is:", max)
	fmt.Println("Left sub-array:", a[:i+1])
	fmt.Println("Right sub-array:", a[i+1:])
}

func splitArrayWithLargestDiff(a []int) (int, int) {
	rSum := 0
	lSum := 0
	for _, e := range a {
		rSum += e
	}

	maxIndex := -1
	maxDiff := 0

	for i := 0; i < len(a)-1; i++ {
		e := a[i]
		lSum += e
		rSum -= e
		diff := lSum - rSum
		if diff < 0 {
			diff *= -1
		}
		if diff > maxDiff {
			fmt.Printf("MaxDiff: %d, LSum: %d, RSum: %d, Diff: %d, Index: %d\n", maxDiff, lSum, rSum, diff, i)
			maxDiff = diff
			maxIndex = i
		}
	}

	return maxIndex, maxDiff
}

func generateArray(n, m int) []int {
	a := make([]int, n)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range a {
		a[i] = r.Intn(m)
		if r.Int()%2 > 0 {
			a[i] *= -1
		}
	}

	return a
}
