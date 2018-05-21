package channels

import (
	"time"
	"fmt"
	"math/rand"
)

type fact struct{
	n int
	f int
}

func ExecuteFactorialPipeline() {
	fmt.Println("Inside channels.ExecuteFactorialPipeline")
	defer fmt.Println("Completed channels.ExecuteFactorialPipeline")

	c := queue()
	out := factorial(c)

	for f := range out {
		fmt.Println("The factorial value is", f)
	}
}

func queue() <-chan int{
	c := make(chan int)
	go func(){
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		nArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _,n := range nArr {
			go func(nn int){
				time.Sleep(time.Duration(r.Intn(5))*time.Millisecond)
				// fmt.Println("Queueing factorial for", nArr[nn])
				// c <- nArr[nn]
				// fmt.Println("Queueing factorial for", nn)
				c <- nn
			}(n)
		}
		time.Sleep(6*time.Millisecond)
		close(c)
	}()
	return c
}


func factorial(c <-chan int) <-chan fact {
	out := make(chan fact)
	go func() {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		for n := range c {
			go func(nn int){
				time.Sleep(time.Duration(r.Intn(5))*time.Millisecond)
				// fmt.Println("Processing factorial for", nn)
				f := 1
				for i := nn; i > 0; i-- {
					f *= i
				}
				// fmt.Println("Queueing result for factorial of", nn)
				out <- fact{nn, f}
			}(n)
		}
		time.Sleep(6*time.Millisecond)
		close(out)
	}()
	return out
}