package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type potKey struct {
	s, e int
}

func main() {
	fmt.Print("How many pots of gold in game: ")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	fmt.Print("Max number of coins in pot: ")
	s.Scan()
	sz, _ := strconv.Atoi(s.Text())

	gPots := generateGoldPots(n, sz)

	counter := 0
	coins := fetchMaxCoins(gPots, &counter)
	fmt.Printf("Recursion: Fetched %d coins as the max coins. Recursion count: %d\n", coins, counter)

	dpMap := make(map[potKey]int)
	counterDP := 0
	coinsDP := fetchMaxCoinsDP(gPots, 0, len(gPots)-1, &counterDP, dpMap)
	// fmt.Printf("Memoization table: %+v\n", dpMap)
	fmt.Printf("Dynamic programming: Fetched %d coins as the max coins. Recursion count: %d\n", coinsDP, counterDP)
}

func fetchMaxCoinsDP(pots []int, s, e int, c *int, dp map[potKey]int) int {
	if v, ok := dp[potKey{s, e}]; ok {
		return v
	}

	*c++

	if s > e {
		return 0
	} else if s == e {
		return pots[s]
	}

	f := pots[s] + min(fetchMaxCoinsDP(pots, s+2, e, c, dp), fetchMaxCoinsDP(pots, s+1, e-1, c, dp))
	l := pots[e] + min(fetchMaxCoinsDP(pots, s+1, e-1, c, dp), fetchMaxCoinsDP(pots, s, e-2, c, dp))

	dp[potKey{s, e}] = max(f, l)
	return max(f, l)
}
func fetchMaxCoins(pots []int, c *int) int {
	*c++

	l := len(pots)
	if l <= 0 {
		return 0
	} else if l == 1 {
		return pots[0]
	}

	s := pots[0] + min(fetchMaxCoins(pots[2:], c), fetchMaxCoins(pots[1:l-1], c))
	e := pots[l-1] + min(fetchMaxCoins(pots[1:l-1], c), fetchMaxCoins(pots[:l-2], c))

	return max(s, e)
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func generateGoldPots(n, sz int) []int {
	if sz < 0 {
		sz = math.MaxInt32
	}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	pots := make([]int, n)
	for i := range pots {
		pots[i] = r.Intn(sz)
	}

	fmt.Println("Pots array:", pots)
	return pots
}
