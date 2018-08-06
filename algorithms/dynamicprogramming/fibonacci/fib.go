package fibonacci

import (
	"fmt"
)

//FibonacciTopDownDynamicProgramming - as the name suggests
func FibonacciTopDownDynamicProgramming(n int) int {
	m := make(map[int]int, 0)
	return fibTopDown(n, m)
}

func fibTopDown(n int, m map[int]int) int {
	fn, f := m[n]
	if f {
		fmt.Printf("Found f%d from memoization table: %d\n", n, fn)
		return fn
	}

	if n == 1 || n == 2 {
		fn = 1
	} else {
		f1 := fibTopDown(n-1, m)
		f2 := fibTopDown(n-2, m)

		fn = f1 + f2
	}

	m[n] = fn
	fmt.Printf("------ Added f%d to memoization table: %d\n", n, fn)
	return fn
}

//FibonacciBottomUpDynamicProgramming - as the name suggests
func FibonacciBottomUpDynamicProgramming(n int) int {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 1

	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
		fmt.Printf("------ Added f%d to memoization table: %d\n", i, m[i])
	}

	return m[n]
}
