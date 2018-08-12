package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the number of lists to generate: ")
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	fmt.Print("Number of elements per list: ")
	s.Scan()
	m, _ := strconv.Atoi(s.Text())

	aa := generateLists(n, m)

	minRange := findMinRange(aa)
	fmt.Printf("Min Range:\nIndex: %+v\n", minRange)
	printListByIndex(aa, minRange)
}

func findMinRange(aa [][]int) []int {
	r := make([]int, len(aa))
	ind := make([]int, len(aa))
	printListByIndex(aa, ind)

	for {
		stop := true
		mInd := -1
		for i := 0; i < len(aa); i++ {
			if len(aa[i]) > ind[i]+1 {
				stop = false
				if mInd < 0 {
					mInd = i
				} else if aa[mInd][ind[mInd]+1] > aa[i][ind[i]+1] {
					mInd = i
				}
			}
		}

		if stop {
			break
		}

		ind[mInd]++
		printListByIndex(aa, ind)

		min, max := findMinMax(aa, ind)
		minR, maxR := findMinMax(aa, r)

		if max-min < maxR-minR {
			copy(r, ind)
		}
	}
	return r
}

func printListByIndex(aa [][]int, ind []int) {
	fmt.Print("[")
	for i, n := range ind {
		fmt.Printf(" %d,", aa[i][n])
	}
	fmt.Print(" ]")
	min, max := findMinMax(aa, ind)
	fmt.Printf(" min: %d, max %d, range:%d\n", min, max, max-min)
}

func findMinMax(aa [][]int, ind []int) (int, int) {
	min, max := math.MaxInt32, math.MinInt32

	for i, j := range ind {
		if aa[i][j] > max {
			max = aa[i][j]
		}

		if aa[i][j] < min {
			min = aa[i][j]
		}
	}

	return min, max
}

func printLists(aa [][]int) {
	for _, a := range aa {
		fmt.Printf("%+v\n", a)
	}
}

func generateLists(n, m int) [][]int {
	aa := make([][]int, n)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range aa {
		for j := 0; j < m; j++ {
			aa[i] = append(aa[i], r.Intn(100))
		}
		sort.Ints(aa[i])
	}

	printLists(aa)
	return aa
}
