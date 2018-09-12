package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type fact struct {
	n int
	f int
}

//ExecuteFactorialPipeline - an example channel pipeline with mulitple concurrent routines adding requests
func ExecuteFactorialPipeline() {
	fmt.Println("Inside channels.ExecuteFactorialPipeline")
	defer fmt.Println("Completed channels.ExecuteFactorialPipeline")

	c := queue()
	out := factorial(c)

	for f := range out {
		fmt.Println("The factorial value is", f)
	}
}

func queue() <-chan int {
	c := make(chan int)
	go func() {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		nArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var wg sync.WaitGroup
		for _, n := range nArr {
			wg.Add(1)
			go func(nn int, wg *sync.WaitGroup) {
				time.Sleep(time.Duration(r.Intn(5)) * time.Millisecond)
				// fmt.Println("Queueing factorial for", nArr[nn])
				// c <- nArr[nn]
				// fmt.Println("Queueing factorial for", nn)
				c <- nn
				wg.Done()
			}(n, &wg)
		}
		wg.Wait()
		close(c)
	}()
	return c
}

func factorial(c <-chan int) <-chan fact {
	out := make(chan fact)
	go func() {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		var wg sync.WaitGroup
		for n := range c {
			wg.Add(1)
			go func(nn int, wg *sync.WaitGroup) {
				time.Sleep(time.Duration(r.Intn(5)) * time.Millisecond)
				// fmt.Println("Processing factorial for", nn)
				f := 1
				for i := nn; i > 0; i-- {
					f *= i
				}
				// fmt.Println("Queueing result for factorial of", nn)
				out <- fact{nn, f}
				wg.Done()
			}(n, &wg)
		}
		wg.Wait()
		close(out)
	}()
	return out
}
