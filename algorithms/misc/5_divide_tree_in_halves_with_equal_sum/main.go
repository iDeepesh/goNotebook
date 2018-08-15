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

type node struct {
	id    int
	value int
	left  *node
	right *node
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter number of nodes in tree: ")
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	fmt.Print("Do you want to set same value in all nodes? Enter value or -1 for random assignments: ")
	s.Scan()
	v, _ := strconv.Atoi(s.Text())

	r := generateRandomTree(n, v, 1000)
	printTree(r)
	end1, end2 := cutTreeInHalves(r)
	if end1 != nil && end2 != nil {
		fmt.Printf("Cut tree on edge joining nodes %+v and %+v\n", end1.id, end2.id)
	} else {
		fmt.Println("Can't cut this tree in two equal halves")
	}
}

func cutTreeInHalves(r *node) (*node, *node) {
	dp := make(map[node]int)
	return recursiveCutTreeInHalves(r, 0, dp)
}

func recursiveCutTreeInHalves(r *node, pSum int, dp map[node]int) (*node, *node) {
	lSum := findSum(r.left, dp)
	rSum := findSum(r.right, dp)
	rootSum := r.value + pSum

	if abs(lSum-rSum) == rootSum {
		// we found the edge to cut on
		if lSum > rSum {
			return r, r.left
		} else {
			return r, r.right
		}
	} else if abs(lSum-rSum) > rootSum {
		// we got to try more
		if lSum > rSum {
			return recursiveCutTreeInHalves(r.left, rootSum+rSum, dp)
		} else {
			return recursiveCutTreeInHalves(r.right, rootSum+lSum, dp)
		}
	} else {
		// cant be solved
		return nil, nil
	}
}

func abs(n int) int {
	if n < 0 {
		n = n * -1
	}
	return n
}

func findSum(r *node, dp map[node]int) int {
	if r == nil {
		return 0
	}

	if v, ok := dp[*r]; ok {
		return v
	}

	sum := r.value + findSum(r.left, dp) + findSum(r.right, dp)
	dp[*r] = sum
	return sum
}

func generateRandomTree(nodeCount, fixedVal, maxVal int) *node {
	fmt.Println("nodeCount at entry:", nodeCount)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	var v int
	if fixedVal != -1 {
		v = fixedVal
	} else {
		v = r.Intn(maxVal)
	}
	root := node{1, v, nil, nil}
	var q []*node
	q = append(q, &root)
	nodeCount--

	id := 2
	i := 0
	for i < len(q) && nodeCount > 0 {
		n := q[i]
		i++
		//flip a coin to decide to add left child
		if n.left == nil && r.Intn(10)&1 == 0 {
			if fixedVal != -1 {
				v = fixedVal
			} else {
				v = r.Intn(maxVal)
			}
			n.left = &node{id, v, nil, nil}
			q = append(q, n.left)
			nodeCount--
			id++
		}

		if nodeCount == 0 {
			continue
		}

		//flip a coin to decide to add right child
		if n.right == nil && r.Intn(10)&1 == 0 {
			if fixedVal != -1 {
				v = fixedVal
			} else {
				v = r.Intn(maxVal)
			}
			n.right = &node{id, v, nil, nil}
			q = append(q, n.right)
			nodeCount--
			id++
		}

		if nodeCount > 0 && i == len(q) {
			i = 0
		}
	}

	fmt.Println("nodeCount at exit:", nodeCount)
	return &root
}

func printTree(r *node) {
	var findMaxLevel func(*node) int
	findMaxLevel = func(r *node) int {
		if r == nil {
			return 0
		}

		leftLevel := findMaxLevel(r.left)
		rightLevel := findMaxLevel(r.right)

		if leftLevel > rightLevel {
			return leftLevel + 1
		} else {
			return rightLevel + 1
		}
	}

	printSpace := func(n int) {
		for i := 0; i < n; i++ {
			fmt.Print(" ")
		}
	}

	var printLevel func([]*node, int, int)
	printLevel = func(nList []*node, l int, maxLevel int) {
		initalSpaces := int(math.Pow(2, float64(maxLevel-l))) - 1
		separaterSpaces := int(math.Pow(2, float64(maxLevel-l+1))) - 1

		isAllElementsNil := true

		printSpace(initalSpaces)
		newList := []*node{}
		for _, n := range nList {
			if n != nil {
				isAllElementsNil = false
				fmt.Printf("%d,%d", n.id, n.value)
				newList = append(newList, n.left)
				newList = append(newList, n.right)
			} else {
				fmt.Print(" ")
				newList = append(newList, nil)
				newList = append(newList, nil)
			}
			printSpace(separaterSpaces)
		}

		fmt.Println("")

		if !isAllElementsNil {
			printLevel(newList, l+1, maxLevel)
		}
	}

	maxLevel := findMaxLevel(r)
	fmt.Println("max Level:", maxLevel)
	nList := []*node{r}
	printLevel(nList, 1, maxLevel)
}
