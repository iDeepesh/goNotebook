package channels

import (
	"fmt"
	"sync"
)

//ExecuteOneKFactorialsInParallel - Runs 1k concurrent go routines to calculate 9m factorials
func ExecuteOneKFactorialsInParallel() {
	fmt.Println("Inside channels.ExecuteOneKFactorials")
	defer fmt.Println("Completed channels.ExecuteOneKFactorials")

	//Create input stream channel
	in := gen()

	//Fan out from 1 to 1000 channels
	cArr := make([]<-chan int, 1000)
	for i := range cArr {
		cArr[i] = factorials(in)
	}
	fmt.Println("Created channels:", len(cArr))

	//Fan in from 1000 to 1 channels
	f := mergeOneK(cArr)

	var counter int
	for range f {
		counter++
	}
	fmt.Println("Number of factorials processed", counter)
}

func mergeOneK(ca []<-chan int) <-chan int {
	fmt.Println("Received the number of channels:", len(ca))
	out := make(chan int)
	var wg sync.WaitGroup

	doMerge := func(in <-chan int) {
		for n := range in {
			out <- n
		}
		wg.Done()
	}

	for i := range ca {
		wg.Add(1)
		go doMerge(ca[i])
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			for j := 1000; j < 10000; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorials(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- facto(n)
		}
		close(out)
	}()
	return out
}

func facto(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
