package adjecencylist

import (
	"fmt"
	"math/rand"
)

//Node - struct fo Adjacency List for graphs
type Node struct {
	data      int
	adj       []*Edge
	visited   bool
	arrival   int
	departure int
}

type Edge struct {
	u *Node
	v *Node
}

//CreateCompleteGraph - as the name suggests
func CreateCompleteGraph(n int) []*Node {
	fmt.Println("Creating a complete graph with element count of", n)
	// var g []*Node
	g := make([]*Node, 0)

	for i := 0; i < n; i++ {
		g = append(g, &Node{i, nil, false, -1, -1})
	}

	l := len(g)
	for i := 0; i < l; i++ {
		adj := make([]*Edge, 0)
		for j := 0; j < n; j++ {
			if i != j {
				adj = append(adj, &Edge{g[i], g[j]})
			}
		}
		g[i].adj = adj
	}

	printGraph(g)
	return g
}

func printGraph(g []*Node) {
	for _, n := range g {
		fmt.Printf("Node: %d, Edges:", n.data)
		for _, e := range n.adj {
			fmt.Printf(" %d-%d,", e.u.data, e.v.data)
		}
		fmt.Println()
	}
}

//CreateRandomGraph - as the name suggests
func CreateRandomGraph(n int) []*Node {
	fmt.Println("Creating a random graph with element count of", n)
	// var g []*Node
	g := make([]*Node, 0)

	for i := 0; i < n; i++ {
		g = append(g, &Node{i, nil, false, -1, -1})
	}

	l := len(g)
	for i := 0; i < l; i++ {
		adj := make([]*Edge, 0)
		index := rand.Perm(n)
		for j := 0; j < n/2; j++ {
			if i != index[j] {
				adj = append(adj, &Edge{g[i], g[index[j]]})
			}
		}
		g[i].adj = adj
	}

	printGraph(g)
	return g
}

//Reset - remove any traversal data to reset the graph
func Reset(g []*Node) {
	for _, n := range g {
		n.visited = false
		n.arrival = -1
		n.departure = -1
	}
}

//DepthFirstTraversal - as the name suggests
func DepthFirstTraversal(g []*Node, ri int) {
	fmt.Print("Printing Depth First Traversal:")

	r := g[ri]
	tE := make([]*Edge, 0)
	t := 0

	var dfs func(*Node)
	dfs = func(n *Node) {
		if n.visited {
			return
		}

		n.visited = true
		n.arrival = t
		t++
		fmt.Printf("%d, ", n.data)

		for _, a := range n.adj {
			if a.v.visited {
				continue
			}
			tE = append(tE, a)
			dfs(a.v)
		}

		n.departure = t
		t++
	}

	dfs(r)
	fmt.Println()

	fmt.Println("Printing the Tree Edges")
	printEdges(tE)

	fE, bE, cE := getEdgesByType(g, tE)

	fmt.Println("Printing the Forward Edges")
	printEdges(fE)

	fmt.Println("Printing the Backward Edges")
	printEdges(bE)

	fmt.Println("Printing the Cross Edges")
	printEdges(cE)
}

// Returns - Forward edges, Backward edges, Cross edges
func getEdgesByType(g []*Node, tE []*Edge) ([]*Edge, []*Edge, []*Edge) {
	fE := make([]*Edge, 0)
	bE := make([]*Edge, 0)
	cE := make([]*Edge, 0)

	for _, n := range g {
		for _, e := range n.adj {
			if findEdge(tE, e) {
				continue
			}

			if e.u.arrival < e.v.arrival && e.u.departure > e.v.departure {
				fE = append(fE, e)
			} else if e.u.arrival > e.v.arrival && e.u.departure < e.v.departure {
				bE = append(bE, e)
			} else if e.v.arrival < e.v.departure && e.v.departure < e.u.arrival && e.u.arrival < e.u.departure {
				cE = append(cE, e)
			}
		}
	}
	return fE, bE, cE
}

func findEdge(l []*Edge, e *Edge) bool {
	for _, le := range l {
		if le.u.data == e.u.data && le.v.data == e.v.data {
			return true
		}
	}

	return false
}

func printEdges(eList []*Edge) {
	for _, e := range eList {
		fmt.Printf("%d-%d, ", e.u.data, e.v.data)
	}
	fmt.Println()
}

//BredthFirstTraversal - as the name suggests
func BredthFirstTraversal(g []*Node, ri int) {
	fmt.Print("Printing Breadth First Traversal:")

	r := g[ri]
	tE := make([]*Edge, 0)
	var q []*Node
	q = append(q, r)
	r.visited = true

	for len(q) > 0 {
		n := q[0]
		fmt.Printf("%d, ", n.data)

		for _, e := range n.adj {
			if e.v.visited {
				continue
			}

			q = append(q, e.v)
			tE = append(tE, e)
			e.v.visited = true
		}
		q = q[1:]
	}

	fmt.Println()

	fmt.Println("Printing the Tree Edges")
	printEdges(tE)
}
