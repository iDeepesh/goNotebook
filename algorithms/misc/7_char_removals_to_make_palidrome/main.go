package main

import (
	"bufio"
	"fmt"
	"os"
)

type intTuple struct {
	i, j int
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Please enter the string or ctrl+c to quit: ")
		s.Scan()
		str := s.Text()

		n := kPalindromeCheck(str)

		if n < 0 {
			// should never come here
			fmt.Println("This string can't be converted to palidrome with any number of removals")
		} else {
			fmt.Printf("This string requires %d deletion to become palindrome\n", n)
		}
	}
}

func kPalindromeCheck(str string) int {
	a := []rune(str)
	l := len(a)

	n, c := recursiveKPalindromeCheck(a, 0, l-1, 0)
	fmt.Println("Total recursive invocation:", c)

	n1, c1 := recursiveKPalindromeCheckDP(a, 0, l-1, 0, make(map[intTuple]int))
	fmt.Println("Total recursive invocation with dynamic programming:", c1)

	if n != n1 {
		n = -1
	}

	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func recursiveKPalindromeCheckDP(a []rune, x, y, c int, dp map[intTuple]int) (int, int) {
	i, j := x, y
	if v, ok := dp[intTuple{i, j}]; ok {
		return v, c
	}

	c++
	n := 0
	for i < j {
		if a[i] != a[j] {
			n1, c1 := recursiveKPalindromeCheckDP(a, i+1, j, c, dp)
			n2, c2 := recursiveKPalindromeCheckDP(a, i, j-1, c1, dp)
			n += min(n1, n2)
			n++
			c = c2
			break
		}

		i++
		j--
	}

	dp[intTuple{x, y}] = n
	return n, c
}

func recursiveKPalindromeCheck(a []rune, i, j, c int) (int, int) {
	c++
	// fmt.Println(c, i, j)
	n := 0
	for i < j {
		if a[i] != a[j] {
			n1, c1 := recursiveKPalindromeCheck(a, i+1, j, c)
			n2, c2 := recursiveKPalindromeCheck(a, i, j-1, c1)
			n += min(n1, n2)
			n++
			c = c2
			break
		}

		i++
		j--
	}

	return n, c
}
