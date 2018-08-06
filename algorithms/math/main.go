package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iDeepesh/goNotebook/algorithms/math/gcd"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Do you want to find Greatest Common Divisor of two numbers - y/n?")
		s.Scan()
		if strings.Compare(s.Text(), "n") == 0 {
			break
		}
		fmt.Print("Enter the first number:")
		s.Scan()
		m, _ := strconv.Atoi(s.Text())
		fmt.Print("Enter the second number:")
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		fmt.Printf("The GCD of %d & %d is %d\n", m, n, gcd.FindGCD(m, n))
	}
}
