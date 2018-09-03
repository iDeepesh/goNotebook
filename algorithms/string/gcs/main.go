package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type lcsCell struct {
	length             int
	left, top, topLeft bool
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Please enter first string: ")
		s.Scan()
		s1 := s.Text()

		fmt.Print("Please enter second string: ")
		s.Scan()
		s2 := s.Text()

		lcs := getLCSTable(s1, s2)
		printLCSTable(s1, s2, lcs)

		gcs := traceLCS(s1, s2, lcs)
		for _, gcss := range gcs {
			fmt.Println(gcss)
		}

		fmt.Print("Do you want to try again, (y) yes or (n) no?")
		s.Scan()
		if strings.Compare(s.Text(), "n") == 0 {
			break
		}
	}
}

func traceLCS(s1, s2 string, lcs [][]lcsCell) []string {
	return rTraceLCS(s1, s2, lcs, len(lcs)-1, len(lcs[0])-1)
}

func rTraceLCS(s1, s2 string, lcs [][]lcsCell, i int, j int) []string {
	if i == 0 || j == 0 {
		return []string{""}
	}

	var r string
	if s1[i-1] == s2[j-1] {
		r = string(s1[i-1])
		gcss := rTraceLCS(s1, s2, lcs, i-1, j-1)
		for k := range gcss {
			gcss[k] = gcss[k] + r
		}
		return gcss
	}

	gcss := []string{}
	if lcs[i][j].top {
		gcss = append(gcss, rTraceLCS(s1, s2, lcs, i-1, j)...)
	}
	if lcs[i][j].left {
		gcss = append(gcss, rTraceLCS(s1, s2, lcs, i, j-1)...)
	}
	return gcss
}

func getLCSTable(r, c string) [][]lcsCell {
	lcs := make([][]lcsCell, len(r)+1)
	for i := range lcs {
		lcs[i] = make([]lcsCell, len(c)+1)
		fillLCSTableRow([]rune(r), []rune(c), i, lcs)
	}

	return lcs
}

func fillLCSTableRow(r, c []rune, i int, lcs [][]lcsCell) {
	if i == 0 {
		return
	}

	for j := range lcs[i] {
		fillLCSTableCell(r, c, i, j, lcs)
	}
}

func fillLCSTableCell(r, c []rune, i, j int, lcs [][]lcsCell) {
	if j == 0 {
		return
	}

	if r[i-1] == c[j-1] {
		lcs[i][j].topLeft = true
		lcs[i][j].length = lcs[i-1][j-1].length + 1
	} else {
		if lcs[i-1][j].length == lcs[i][j-1].length {
			lcs[i][j].length = lcs[i-1][j].length
			lcs[i][j].top = true
			lcs[i][j].left = true
		} else if lcs[i-1][j].length > lcs[i][j-1].length {
			lcs[i][j].length = lcs[i-1][j].length
			lcs[i][j].top = true
		} else {
			lcs[i][j].length = lcs[i][j-1].length
			lcs[i][j].left = true
		}
	}
}

func printLCSTable(s1, s2 string, lcs [][]lcsCell) {
	fmt.Println("\nHere is the LCS table for strings:")
	fmt.Print("      |")
	fmt.Print("      |")
	for i := range s2 {
		fmt.Printf("   %s  |", string(s2[i]))
	}
	fmt.Println()
	fmt.Print("-------")
	fmt.Print("-------")
	for range s2 {
		fmt.Print("-------")
	}
	fmt.Println()
	for i, r := range lcs {
		if i > 0 {
			fmt.Printf("   %s  |", string(s1[i-1]))
		} else {
			fmt.Print("      |")
		}
		for _, c := range r {
			if c.left {
				fmt.Print("<")
			} else {
				fmt.Print(" ")
			}
			if c.topLeft {
				fmt.Print("\\")
			} else {
				fmt.Print(" ")
			}
			if c.top {
				fmt.Print("^")
			} else {
				fmt.Print(" ")
			}
			fmt.Printf("%2d |", c.length)
		}
		fmt.Println()
		fmt.Print("-------")
		for range r {
			fmt.Print("-------")
		}
		fmt.Println()
	}
}
