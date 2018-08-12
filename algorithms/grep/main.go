package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iDeepesh/goNotebook/algorithms/grep/kmp"
)

func main() {
	fmt.Print("Please enter a string to search for: ")

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	p := s.Text()

	pf := kmp.GetPartialFunction(p)
	fmt.Println("Here is the Partial/Failure function table for:", p)
	for i, v := range pf {
		fmt.Printf("[%q:%d] ", p[i], v)
	}
	fmt.Println()
}
