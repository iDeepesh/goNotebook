package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Node - graph node
type Node struct {
	id  int
	adj []*Edge
}

//Edge - graph edge
type Edge struct {
	u, v *Node
	wt   int
}

func main() {
	fmt.Println("Welcome to Travelling Salesman demonstration.....")
	s := bufio.NewScanner(os.Stdin)
	stop := false
	for !stop {
		fmt.Print("Please indicate type of TSP calculation - (bf) brute force, (bb) branch & bound, (b) both:")
		s.Scan()
		c := s.Text()
		bf, bb := true, true
		if strings.Compare(c, "bf") == 0 {
			bb = false
		} else if strings.Compare(c, "bb") == 0 {
			bf = false
		}

		fmt.Print("Please enter the no of nodes (max 10 for brute force) for sample graph:")
		s.Scan()
		n, e := strconv.Atoi(s.Text())
		if e != nil {
			fmt.Println("Please enter a valid integer. Error:", e)
			return
		}

		g := createCompleteGraph(n)

		if bf {
			t1 := time.Now()
			p1, c1 := findTSP(g)
			d1 := time.Since(t1).Nanoseconds() / 1000

			fmt.Println("\nTime taken by brute force TSP:", d1)
			fmt.Println("Cost from the brute force TSP:", c1)
			printPath(p1)
		}

		if bb {
			t2 := time.Now()
			p2, c2 := findTSPBB(g)
			d2 := time.Since(t2).Nanoseconds() / 1000

			fmt.Println("\nTime taken by branch and bound TSP:", d2)
			fmt.Println("Cost from the branch and bound TSP:", c2)
			printPathWithEdges(p2)
		}

		fmt.Print("\nDo you want to try again - (y) Or (n): ")
		s.Scan()
		if strings.Compare(s.Text(), "n") == 0 {
			stop = true
		}
	}
}

func findTSPBB(g []*Node) ([]*Edge, int) {
	return rTSPBB(g[0], g[0], []*Edge{}, len(g), math.MinInt32)
}

func rTSPBB(s, f *Node, p []*Edge, nn, c int) ([]*Edge, int) {
	cc := 0
	for _, e := range p {
		cc += e.wt
	}

	refCost := c
	curCost := math.MaxInt32
	var path []*Edge

	for _, e := range s.adj {
		var tc int
		var tp []*Edge
		if refCost > math.MinInt32 && cc+e.wt > refCost {
			// branch and bound
			continue
		} else if e.v.id == f.id && len(p) == nn-1 {
			//base case
			tp = append(p, e)
			tc = cc + e.wt
		} else if isInPath(e.v, p) {
			continue
		} else {
			//recursive case
			tp, tc = rTSPBB(e.v, f, append(p, e), nn, refCost)
		}
		if tc < curCost {
			curCost = tc
			path = tp
			refCost = tc
		}
	}

	return path, curCost
}

func isInPath(n *Node, a []*Edge) bool {
	for _, e := range a {
		if e.u.id == n.id || e.v.id == n.id {
			return true
		}
	}
	return false
}

func findTSP(g []*Node) ([]*Node, int) {
	return rTSP(g[0], g[0], []*Node{}, len(g))
}

func rTSP(s, f *Node, p []*Node, nn int) ([]*Node, int) {
	var path []*Node
	cost := math.MaxInt32
	for _, e := range s.adj {
		var p1 []*Node
		var c int
		if e.v.id == f.id && len(p) == nn-1 {
			c = e.wt
			p1 = []*Node{s, e.v}
		} else if traversed(e.v, p) {
			continue
		} else {
			p1, c = rTSP(e.v, f, append(p, s), nn)
			c += e.wt
			p1 = append([]*Node{s}, p1...)
		}
		if c < cost {
			cost = c
			path = p1
		}
	}

	return path, cost
}

func traversed(n *Node, a []*Node) bool {
	for i := range a {
		if a[i].id == n.id {
			return true
		}
	}

	return false
}

func createCompleteGraph(n int) []*Node {
	fmt.Println("Creating a complete graph with element count of", n)

	g := make([]*Node, 0)

	for i := 0; i < n; i++ {
		g = append(g, &Node{i, nil})
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	l := len(g)
	for i := 0; i < l; i++ {
		adj := make([]*Edge, 0)
		for j := 0; j < n; j++ {
			if i != j {
				adj = append(adj, &Edge{g[i], g[j], r.Intn(10) + 1})
			}
		}
		g[i].adj = adj
	}

	printGraph(g)
	return g
}

func printPathWithEdges(e []*Edge) {
	fmt.Printf("Path: %d", e[0].u.id)
	for i := range e {
		fmt.Printf(" -> %d", e[i].v.id)
	}
	fmt.Println()
}

func printPath(g []*Node) {
	fmt.Printf("Path: %d", g[0].id)
	for i := range g {
		if i != 0 {
			fmt.Printf(" -> %d", g[i].id)
		}
	}
	fmt.Println()
}

func printGraph(g []*Node) {
	for _, n := range g {
		fmt.Printf("Node: %d, Edges:", n.id)
		printEdges(n.adj)
	}
}

func printEdges(eList []*Edge) {
	for _, e := range eList {
		fmt.Printf("%d-%d:%d, ", e.u.id, e.v.id, e.wt)
	}
	fmt.Println()
}
