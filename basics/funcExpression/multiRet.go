package funcexpression

import (
	"fmt"
)

//ExecuteMultiRetFuncExpression - an example of multiple return values from function expression
func ExecuteMultiRetFuncExpression() {
	fmt.Println("Inside EfuncExpression.xecuteMutliRetFuncExpression")
	defer fmt.Println("Completed funcExpression.ExecuteMutliRetFuncExpression")

	he := func(n int) (int, bool) {
		return n / 2, (n&1 == 0)
	}
	for i := 1; i <= 10; i++ {
		h, e := he(i)
		fmt.Printf("Input = %d, half = %d, Even = %t\n", i, h, e)
	}
}

//ExecuteMultiRet - an example of multiple return values from function
func ExecuteMultiRet() {
	fmt.Println("Inside funcExpression.ExecuteMutliRet")
	defer fmt.Println("Completed funcExpression.ExecuteMutliRet")

	for i := 1; i <= 10; i++ {
		h, e := getHalfAndEven(i)
		fmt.Printf("Input = %d, half = %d, Even = %t\n", i, h, e)
	}
}

func getHalfAndEven(n int) (int, bool) {
	return n / 2, (n&1 == 0)
}
