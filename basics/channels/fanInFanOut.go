package channels

import (
	"time"
	"sync"
	"fmt"
	"math/rand"
)

type intTuple struct{
	number int
	squared int
}

func ExecuteFanInFanOut(){
	fmt.Println("Inside channels.ExecuteFanInFanOut")
	defer fmt.Println("Completed channels.ExecuteFanInFanOut")

	//Input stream
	c1 := generator(1,2,3,4,5,6,7,8,9,10)

	//Fan out
	c2 := square(c1)
	c3 := square(c1)
	c4 := square(c1)

	//Fan in
	c5 := merge(c2,c3,c4)
	for n := range c5{
		fmt.Println("Squared numbers", n)
	}
}

func generator(n ...int) <-chan int{
	out := make(chan int)

	go func(){
		for _, nn := range n{
			out <- nn
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan intTuple {
	out := make(chan intTuple)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	go func(){
		for n := range in{
			time.Sleep(time.Duration(r.Intn(5))*time.Millisecond)
			out <- intTuple{n, n*n}
		}
		close(out)
	}()

	return out
}

func merge(in ...<-chan intTuple) <-chan intTuple{
	out := make(chan intTuple)
	var wg sync.WaitGroup
	wg.Add(len(in))

	for _, c := range in{
		go func(cIn <-chan intTuple){
			for n := range cIn{
				out<-n
			}
			wg.Done()
		}(c)
	}

	go func(){
		wg.Wait()
		close(out)
	}()

	return out
}