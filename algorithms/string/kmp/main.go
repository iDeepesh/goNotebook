package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an string to search from: ")
		s.Scan()
		str := s.Text()
		printStringAsArray(str)

		fmt.Print("Please enter a string to search for: ")
		s.Scan()
		p := s.Text()

		i := searchPattern(str, p)
		fmt.Println("The pattern starts at the index of:", i)

		fmt.Print("Continue with another search, (y) yes, (q) quit: ")
		s.Scan()
		q := s.Text()
		if strings.Compare(q, "q") == 0 {
			break
		}
	}
}

func searchPattern(s, p string) int {
	pf := getPartialFunction(p)

	r := -1
	i, j := 0, 0
	for i < len(s) {
		// fmt.Printf("i=%d, s[i]= %q, j=%d, p[j]=%q\n", i, s[i], j, p[j])
		if s[i] == p[j] {
			if j == len(p)-1 {
				r = i - j
				break
			}
			i++
			j++
		} else if j != 0 {
			//In case of mismatch take the pattern index from previous element of pf table
			j = pf[j-1]
		} else { //if j == 0
			i++
		}
	}
	return r
}

// simple non efficient partial function table
func getPartialFunction(w string) []int {
	printStringAsArray(w)
	p := []rune(w)
	pf := make([]int, len(p))
	pf[0] = 0

	j, i := 0, 1
	for i < len(pf) {
		fmt.Printf("i=%d, p[i]= %q, j=%d, p[j]=%q\n", i, p[i], j, p[j])
		if p[i] == p[j] {
			pf[i] = j + 1
			i++
			j++
		} else if j != 0 {
			//In case of mismatch take the pattern index from previous element of pf table
			j = pf[j-1]
		} else { //if j == 0
			i++
		}
	}

	fmt.Println("Here is the Partial/Failure function table for:", w)
	printStringAsArray(w)
	for _, v := range pf {
		fmt.Printf(" %2d |", v)
	}
	fmt.Println()

	return pf
}

func printStringAsArray(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf(" %2d |", i)
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("  %s |", string(s[i]))
	}
	fmt.Println()
}
