package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println("This program evaluates mathematical expression with single digit numbers and no spaces in between.")
	fmt.Println("The supported list of operators is +, -, *, /, (, ).")
	fmt.Println("Example: 2+5*(6/2)-9")
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nPlease enter the expression to evaluate or (q) to quit: ")
		s.Scan()
		t := s.Text()
		if strings.Compare(t, "q") == 0 {
			break
		}

		iFix := []rune(t)
		pFix := convertToPostFix(iFix)
		fmt.Println("Here is the PostFix/RPN notation of expression:", string(pFix))
		fmt.Printf("The evaluation result is %s = %d\n", t, executeExpression(pFix))
	}
}

func executeExpression(pFix []rune) int {
	// * Create a stack to store operands (or values).
	// * Scan the given expression and do following for every scanned element.
	// 	* If the element is a number, push it into the stack
	// 	* If the element is a operator, pop operands for the operator from stack. Evaluate the operator and push the result back to the stack
	// * When the expression is ended, the number in the stack is the final answer

	s := []int{}
	for _, r := range pFix {
		if isNum(r) {
			s = append(s, int(r-'0'))
			continue
		}

		a, b := s[len(s)-2], s[len(s)-1]
		s = s[:len(s)-2]
		s = append(s, evaluate(a, b, r))
	}

	return s[0]
}

func evaluate(a, b int, o rune) int {
	var r int
	switch o {
	case '+':
		r = a + b
	case '-':
		r = a - b
	case '*':
		r = a * b
	case '/':
		r = a / b
	}

	return r
}

func convertToPostFix(in []rune) []rune {
	// * While there are tokens to be read:
	//     * Read a token.
	//     * If the token is a number, then add it to the output queue.
	//     * If the token is an operator, o1, then:
	//         * while there is an operator token, o2, at the top of the stack, and
	//             * o1 is left-associative and its precedence is *less than or equal* to that of o2,
	//                 * then pop o2 off the stack, onto the output queue;
	//         * push o1 onto the stack.
	//     * If the token is a left parenthesis, then push it onto the stack.
	//     * If the token is a right parenthesis:
	//         * Until the token at the top of the stack is a left parenthesis, pop operators off the stack onto the output queue.
	//         * Pop the left parenthesis from the stack, but not onto the output queue.

	out := []rune{}
	op := []rune{}

	for _, e := range in {
		if isNum(e) {
			out = append(out, e)
			continue
		}

		if e == ')' {
			for len(op) > 0 {
				o := op[len(op)-1]
				op = op[:len(op)-1]
				if o == '(' {
					break
				}

				out = append(out, o)
			}
			continue
		}

		for len(op) > 0 && op[len(op)-1] != '(' && precedence(op[len(op)-1]) >= precedence(e) {
			out = append(out, op[len(op)-1])
			op = op[:len(op)-1]
		}

		op = append(op, e)
	}

	for i := len(op) - 1; i >= 0; i-- {
		out = append(out, op[i])
	}

	return out
}

func precedence(r rune) int {
	switch r {
	case '+':
		return 1
	case '-':
		return 1
	case '*':
		return 2
	case '/':
		return 2
	}
	return math.MaxInt16
}

func isNum(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}

	return false
}
