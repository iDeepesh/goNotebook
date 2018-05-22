package interfaces

import (
	"fmt"
)

//ExecuteAssertions - an example of assertion (establishing instance type from interface type)
func ExecuteAssertions() {
	fmt.Println("Inside interfaces.ExecuteAssertions")
	defer fmt.Println("Completed interfaces.ExecuteAssertions")

	c := circle{4}

	var s shape
	s = c

	printCircle(s.(circle)) //Needs assertion to make sure that it is circle
}

func printCircle(c circle) {
	fmt.Printf("This is a cirlce. %T\n", c)
}
