package funcexpression

import (
	"fmt"
)

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/

//ExecuteCallback - an example of using func variable for callback
func ExecuteCallback() {
	fmt.Println("Inside funcExpression.ExecuteCallback")
	defer fmt.Println("Completed funcExpression.ExecuteCallback")

	cntr := counter()
	useCounter(cntr)
}

//ExecuteClosure - an example of using closure
func ExecuteClosure() {
	fmt.Println("Inside funcExpression.ExecuteClosure")
	defer fmt.Println("Completed funcExpression.ExecuteClosure")

	c := counter()
	fmt.Printf("First call to counter: %d\n", c())
	fmt.Printf("Second call to counter: %d\n", c())
	fmt.Printf("Third call to counter: %d\n", c())
}

func counter() func() int {
	var c int
	return func() int {
		c++
		return c
	}
}

func useCounter(c func() int) {
	fmt.Printf("First call to counter: %d\n", c())
	fmt.Printf("Second call to counter: %d\n", c())
	fmt.Printf("Third call to counter: %d\n", c())
}
