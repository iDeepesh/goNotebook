package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//ExecuteChannelRange - an example of using for with range for channels. Uses WaitGroup for managing lifecycle of channel
func ExecuteChannelRange() {
	fmt.Println("Inside channels.ExecuteChannelRange")
	defer fmt.Println("Completed channels.ExecuteChannelRange")

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var wg sync.WaitGroup
	wg.Add(10)

	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func(n int) {
			time.Sleep(time.Duration(r.Intn(5)) * time.Millisecond)
			c <- n
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

//ExecuteChannelSemaphores - an example of using semaphores for managing lifecycle of channel
func ExecuteChannelSemaphores() {
	fmt.Println("Inside channels.ExecuteChannelSemaphores")
	defer fmt.Println("Completed channels.ExecuteChannelSemaphores")

	c := make(chan int)
	done := make(chan bool)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 10; i++ {
		go func(n int) {
			time.Sleep(time.Duration(r.Intn(5)) * time.Millisecond)
			c <- n
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

//ExecuteChannelConcurrentPolling - an example of mulitple routines polling concurrently from a channel
func ExecuteChannelConcurrentPolling() {
	fmt.Println("Inside channels.ExecuteChannelConcurrentPolling")
	defer fmt.Println("Completed channels.ExecuteChannelConcurrentPolling")

	c := make(chan int)
	done := make(chan bool)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 10; i++ {
		go func(n int) {
			time.Sleep(time.Duration(r.Intn(5)) * time.Millisecond)
			c <- n
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			<-done
		}
		close(c)
	}()

	// poll := func(i int) {
	go func(i int) {
		for n := range c {
			fmt.Println("Poller:", i, "Value:", n)
		}
	}(1)

	// go poll(1)
	// go poll(2)
}

//ExecuteDirectionalChannels - an example of directional channel
func ExecuteDirectionalChannels() {
	fmt.Println("Inside channels.ExecuteDirectional channels")
	defer fmt.Println("Completed channels.ExecuteDirectional channels")

	c := func() <-chan int {
		c := make(chan int)
		go func(c chan<- int) {
			for i := 0; i < 10; i++ {
				c <- i
			}
			close(c)
		}(c)
		return c
	}()

	r := func(c <-chan int) <-chan int {
		var sum int
		for n := range c {
			sum += n
		}
		rC := make(chan int)
		go func(n int, c chan<- int) {
			c <- n
			close(c)
		}(sum, rC)
		return rC
	}(c)

	for sum := range r {
		fmt.Println("The sum is", sum)
	}
}
