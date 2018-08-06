package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iDeepesh/goNotebook/algorithms/dynamicprogramming/fibonacci"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please enter the number to calculate fibonacci or q to quit:")
		s.Scan()
		t := s.Text()
		if strings.Compare(t, "q") == 0 {
			break
		}

		n, _ := strconv.Atoi(t)

		fmt.Print("Type of dyniamic programming to apply Top down (t) or Bottom up (b):")
		s.Scan()
		dp := s.Text()
		var fn int
		if strings.Compare(dp, "t") == 0 {
			fn = fibonacci.FibonacciTopDownDynamicProgramming(n)
		} else {
			fn = fibonacci.FibonacciBottomUpDynamicProgramming(n)
		}

		fmt.Printf("The fibonacci for %d is %d\n", n, fn)
	}
}
