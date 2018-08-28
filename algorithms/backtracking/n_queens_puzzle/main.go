package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	stop := false
	for !stop {
		fmt.Print("Please enter the number of queens or (q) for quitting: ")
		s.Scan()
		t := s.Text()
		if strings.Compare(t, "q") == 0 {
			stop = true
			break
		}

		n, _ := strconv.Atoi(t)

		fmt.Print("Do you want to get all solutions, y or n: ")
		s.Scan()
		t = s.Text()
		if strings.Compare(t, "n") == 0 {
			sol := getFirstSolution(n)
			printSolution(sol)
		} else {
			sols := getAllSolutions(n)
			printAllSolutions(sols)
		}
	}
}

func getAllSolutions(n int) [][]int {
	sols := make([][]int, 0)

	s := make([]int, n)
	for i := range s {
		s[i] = -1
	}

	for i := 0; i < n; i++ {
		s[0] = i
		recordAllSolutions(1, s, &sols)
	}

	return sols
}

func recordAllSolutions(i int, s []int, ss *[][]int) {
	if i >= len(s) {
		sol := make([]int, len(s))
		for j := range s {
			sol[j] = s[j]
		}
		*ss = append(*ss, sol)
		return
	}

	for j := 0; j < len(s); j++ {
		// fmt.Printf("i=%d, s=%+v\n", i, s)
		s[i] = j
		if !reject(i, s) {
			recordAllSolutions(i+1, s, ss)
		}
	}

	s[i] = -1
}

func printAllSolutions(s [][]int) {
	for i := range s {
		fmt.Println("Solution #", i+1)
		printSolution(s[i])
	}
}

func getFirstSolution(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = -1
	}

	for i := 0; i < n; i++ {
		s[0] = i
		if findSolution(1, s) {
			break
		}
	}

	return s
}

func findSolution(i int, s []int) bool {
	if i >= len(s) {
		return true
	}

	for j := 0; j < len(s); j++ {
		s[i] = j
		if !reject(i, s) {
			if findSolution(i+1, s) {
				return true
			}
		}
	}

	s[i] = -1
	return false
}

func reject(i int, s []int) bool {
	// fmt.Printf("i=%d, %+v\n", i, s)
	return rowThreat(i, s) || diagonalThreat(i, s)
}

func rowThreat(i int, s []int) bool {
	for j := 0; j < i; j++ {
		if s[j] == s[i] {
			// fmt.Println("Row Threat")
			return true
		}
	}

	return false
}

func diagonalThreat(i int, s []int) bool {
	for j := 0; j < i; j++ {
		d := s[i] - s[j]
		if d < 0 {
			d *= -1
		}

		if d == i-j {
			// fmt.Println("Diagonal Threat")
			return true
		}
	}

	return false
}

func printSolution(s []int) {
	for _, i := range s {
		printLine(i, len(s))
	}
}

func printLine(i, n int) {
	for j := 0; j < n; j++ {
		if j != i {
			fmt.Print("|_")
		} else {
			fmt.Print("|x")
		}
	}
	fmt.Println("|")
}
