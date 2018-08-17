package main

import (
	"bufio"
	"fmt"
	"os"
)

type dpIndex struct {
	i, j int
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter the first string: ")
	s.Scan()
	s1 := s.Text()

	fmt.Print("Please enter the second string: ")
	s.Scan()
	s2 := s.Text()

	dist := calculateEditDistance(s1, s2)
	fmt.Println("No of changes required to go from first to second string:", dist)
}

func calculateEditDistance(s1, s2 string) int {
	r1, r2 := []rune(s1), []rune(s2)
	ed, c := recursiveEditDistance(r1, r2, len(r1)-1, len(r2)-1, 0)
	fmt.Println("Total recursive calls:", c)

	dp := make(map[dpIndex]int)
	eddp, cdp := recursiveEditDistanceDP(r1, r2, len(r1)-1, len(r2)-1, 0, dp)
	fmt.Println("Total recursive calls with Dynamic Programming:", cdp)

	if ed != eddp {
		return -1
	}

	return ed
}

func recursiveEditDistanceDP(r1, r2 []rune, l1, l2, c int, dp map[dpIndex]int) (int, int) {
	if v, ok := dp[dpIndex{l1, l2}]; ok {
		return v, c
	}

	c++
	var dist int
	if l1 == 0 {
		dist = l2
	} else if l2 == 0 {
		dist = l1
	} else if r1[l1] == r2[l2] {
		dist, c = recursiveEditDistanceDP(r1, r2, l1-1, l2-1, c, dp)
	} else {
		d, cd := recursiveEditDistanceDP(r1, r2, l1-1, l2, c, dp)
		i, ci := recursiveEditDistanceDP(r1, r2, l1, l2-1, cd, dp)
		r, cr := recursiveEditDistanceDP(r1, r2, l1-1, l2-1, ci, dp)

		dist = min(d, i, r) + 1
		c = cr
	}
	dp[dpIndex{l1, l2}] = dist
	return dist, c
}

func recursiveEditDistance(r1, r2 []rune, l1, l2, c int) (int, int) {
	c++
	if l1 == 0 {
		return l2, c
	}

	if l2 == 0 {
		return l1, c
	}

	if r1[l1] == r2[l2] {
		return recursiveEditDistance(r1, r2, l1-1, l2-1, c)
	}

	d, cd := recursiveEditDistance(r1, r2, l1-1, l2, c)
	i, ci := recursiveEditDistance(r1, r2, l1, l2-1, cd)
	r, cr := recursiveEditDistance(r1, r2, l1-1, l2-1, ci)

	return 1 + min(d, i, r), cr
}

func min(i, j, k int) int {
	if i < j && i < k {
		return i
	}

	if j < i && j < k {
		return j
	}

	return k
}
