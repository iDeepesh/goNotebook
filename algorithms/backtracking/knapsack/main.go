package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type item struct {
	w, v int
}

type key struct {
	c, i int
}

type dpVal struct {
	ks    []item
	total item
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	stop := false
	for !stop {
		fmt.Print("Please enter the max weight/value of the item: ")
		s.Scan()
		w, _ := strconv.Atoi(s.Text())
		fmt.Print("Please enter the number of available items: ")
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		items := generateItems(n, w)
		fmt.Println(items)

		fmt.Print("\nPlease enter the capacity of the knapsack: ")
		s.Scan()
		c, _ := strconv.Atoi(s.Text())

		fmt.Println("\nFilling knapsack ...........................")
		count := 0
		t1 := time.Now()
		ks, totals, count := fillKnapsack(c, 0, items, count)
		d1 := time.Since(t1)
		fmt.Printf("Here is the Knapsack with max value:%d and weight: %d\n", totals.v, totals.w)
		fmt.Printf("Total recursive calls: %d and total duration: %d\n", count, d1.Nanoseconds()/1000)
		fmt.Println(ks)

		fmt.Println("\nFilling knapsack with Dynamic Programming ...........................")
		dpTable := make(map[key]dpVal)
		countDP := 0
		t2 := time.Now()
		ksDP, totalsDP, countDP := fillKnapsackDP(c, 0, items, dpTable, countDP)
		d2 := time.Since(t2)
		fmt.Printf("Here is the Knapsack with max value:%d and weight: %d\n", totalsDP.v, totalsDP.w)
		fmt.Printf("Total recursive calls with Dynamic Programming: %d and total duration: %d\n", countDP, d2.Nanoseconds()/1000)
		fmt.Println(ksDP)

		fmt.Print("\nNow what? (c) continue or (q) quit: ")
		s.Scan()
		if strings.Compare(s.Text(), "q") == 0 {
			stop = true
		}
	}
}

func fillKnapsackDP(c, i int, items []item, dpTable map[key]dpVal, count int) ([]item, item, int) {
	k := key{c, i}
	if dp, ok := dpTable[k]; ok {
		return dp.ks, dp.total, count
	}

	count++

	if i >= len(items) || c <= 0 {
		return make([]item, 0), item{0, 0}, count
	}

	ks, total, woCount := fillKnapsackDP(c, i+1, items, dpTable, count)
	count = woCount

	if c >= items[i].w {
		wCurrent, wTotal, wCount := fillKnapsackDP(c-items[i].w, i+1, items, dpTable, count)
		count = wCount
		wCurrent = append(wCurrent, items[i])
		wTotal.w, wTotal.v = wTotal.w+items[i].w, wTotal.v+items[i].v

		if wTotal.v > total.v {
			ks, total = wCurrent, wTotal
		}
	}

	dpTable[k] = dpVal{ks, total}
	return ks, total, count
}

func fillKnapsack(c, i int, items []item, count int) ([]item, item, int) {
	count++

	if i >= len(items) || c <= 0 {
		return make([]item, 0), item{0, 0}, count
	}

	woCurrent, woTotal, woCount := fillKnapsack(c, i+1, items, count)
	count = woCount

	if c-items[i].w < 0 {
		return woCurrent, woTotal, count
	}

	wCurrent, wTotal, wCount := fillKnapsack(c-items[i].w, i+1, items, count)
	count = wCount
	wCurrent = append(wCurrent, items[i])
	wTotal.w, wTotal.v = wTotal.w+items[i].w, wTotal.v+items[i].v

	if wTotal.v > woTotal.v {
		return wCurrent, wTotal, count
	} else {
		return woCurrent, woTotal, count
	}
}

func generateItems(n, w int) []item {
	fmt.Printf("Generating %d random items with max weight/value of %d ........................ \n", n, w)
	items := make([]item, n)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range items {
		items[i].w = r.Intn(w)
		items[i].v = r.Intn(w)
	}

	return items
}
