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
	fmt.Print("Enter the number of unique elements in array: ")
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	fmt.Print("Enter the max value for elements in array: ")
	s.Scan()
	m, _ := strconv.Atoi(s.Text())

	a := generateArrayWithRepeatingNumbers(n, m)
	// fmt.Println(a)
	x, y := getNumbersWithOddReps(a)
	fmt.Printf("The numbers with odd number of repititions are: %d, %d\n", x, y)
}

func getNumbersWithOddReps(a []int) (int, int) {
	first := 0
	for _, n := range a {
		first = first ^ n
	}
	second := first

	on := 1
	for {
		if first&on > 0 {
			break
		}
		on = on << 1
	}

	for _, n := range a {
		if n&on > 0 {
			first = first ^ n
		} else {
			second = second ^ n
		}
	}

	return first, second
}

func generateArrayWithRepeatingNumbers(n, m int) []int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	a := make([]int, 0)
	odd := 0

	for i := 0; i < n; i++ {
		num := r.Intn(m)
		reps := r.Intn(10)
		if reps&1 == 1 {
			if odd >= 2 {
				reps++
			} else {
				odd++
				fmt.Print("ODD: ")
			}
		} else if odd < 2 {
			reps++
			odd++
		}
		fmt.Printf("Adding %d for %d number of times\n", num, reps)

		for j := 0; j < reps; j++ {
			a = append(a, num)
		}
	}

	ret := make([]int, len(a))
	p := rand.Perm(len(ret))
	for i, j := range p {
		ret[j] = a[i]
	}

	return ret
}
