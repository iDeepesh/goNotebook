package funcExpression

import (
	"fmt"
)

func ExecuteVariadicArgs(){
	fmt.Println("Insde funcExpression.ExecuteVariadicArgs")
	defer fmt.Println("Completed funcExpression.ExecuteVariadicArgs")

	ns := []int{23,93,2879,2390,843,2309,98734,237,34098}
	l := getLargest(ns...)
	fmt.Printf("The largest number is %d\n", l)
}

func ExecuteVariadicParams(){
	fmt.Println("Insde funcExpression.ExecuteVariadicParams")
	defer fmt.Println("Completed funcExpression.ExecuteVariadicParams")

	l := getLargest(23,93,2879,2390,843,2309,98734,237,34098)
	fmt.Printf("The largest number is %d\n", l)
}

func getLargest(nv ...int) int{
	var r int
	for _,n := range nv{
		if n > r{
			r = n
		}
	}
	return r
}