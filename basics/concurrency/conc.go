package concurrency

import (
	"sync/atomic"
	"time"
	"sync"
	"fmt"
	"math/rand"
)

func ExecuteConcurrentWaitGroup(){
	fmt.Println("Inside concurrency.ExecuteConcurrentWaitGroup")
	defer fmt.Println("Completed concurrency.ExecuteConcurrentWaitGroup")

	var c int
	var wg sync.WaitGroup

	wg.Add(2)
	go incrementor("foo", &c, &wg)
	go incrementor("bar", &c, &wg)
	wg.Wait()

	fmt.Println("Finale counter value:", c)
}

func ExecuteConcurrentMutex(){
	fmt.Println("Inside concurrency.ExecuteConcurrentMutex")
	defer fmt.Println("Completed concurrency.ExecuteConcurrentMutex")

	var c int
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(2)
	go mutexIncrementor("foo", &c, &wg, &m)
	go mutexIncrementor("bar", &c, &wg, &m)
	wg.Wait()

	fmt.Println("Finale counter value:", c)
}

func ExecuteConcurrentAtomic(){
	fmt.Println("Inside concurrency.ExecuteConcurrentAtomic")
	defer fmt.Println("Completed concurrency.ExecuteConcurrentAtomic")

	var c int32
	var wg sync.WaitGroup

	wg.Add(2)
	go atomicIncrementor("foo", &c, &wg)
	go atomicIncrementor("bar", &c, &wg)
	wg.Wait()

	fmt.Println("Finale counter value:", c)
}

func incrementor(s string, c *int, wg *sync.WaitGroup){
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	for i:=0;i<1000;i++{
		*c++
		// fmt.Println(s, "index:", i, "counter:", *c)
		time.Sleep(time.Duration(r.Intn(5)) * time.Millisecond)
	}

	wg.Done()
}

func mutexIncrementor(s string, c *int, wg *sync.WaitGroup, m *sync.Mutex){
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	for i:=0;i<1000;i++{
		m.Lock()
		*c++
		// fmt.Println(s, "index:", i, "counter:", *c)
		m.Unlock()
		time.Sleep(time.Duration(r.Intn(5)) * time.Millisecond)
	}

	wg.Done()
}

func atomicIncrementor(s string, c *int32, wg *sync.WaitGroup){
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	for i:=0;i<10;i++{
		fmt.Println(s, "index:", i, "counter:", atomic.AddInt32(c, 1))
		time.Sleep(time.Duration(r.Intn(5)) * time.Millisecond)
	}

	wg.Done()
}