package adjecencylist

import (
	"fmt"
	"math/rand"
)

//Node - struct fo Adjacency List for graphs
type Node struct {
	data      int
	adj       []*Node
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
		adj := make([]*Node, 0)
		for j := 0; j < n; j++ {
			if i != j {
				adj = append(adj, g[j])
			}
		}
		g[i].adj = adj
	}

	return g
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
		adj := make([]*Node, 0)
		index := rand.Perm(n)
		for j := 0; j < n/2; j++ {
			if i != index[j] {
				adj = append(adj, g[index[j]])
			}
		}
		g[i].adj = adj
	}

	return g
}

//Reset - remove any traversal data to reset the graph
func Reset(g []*Node) {
	for _, n := range g {
		n.visited = false
	}
}

//DepthFirstTraversal - as the name suggests
func DepthFirstTraversal(g []*Node, ri int) {
	fmt.Print("Printing DFT:")

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
		fmt.Print("\t", n.data)

		for _, a := range n.adj {
			if a.visited {
				continue
			}
			tE = append(tE, &(Edge{n, a}))
			dfs(a)
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

	for _, u := range g {
		for _, v := range u.adj {
			e := Edge{u, v}
			if findEdge(tE, &e) {
				continue
			}

			if u.arrival < v.arrival && u.departure > v.departure {
				fE = append(fE, &e)
			} else if u.arrival > v.arrival && u.departure < v.departure {
				bE = append(bE, &e)
			} else if v.arrival < v.departure && v.departure < u.arrival && u.arrival < u.departure {
				cE = append(cE, &e)
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
		fmt.Printf("\t%d-%d", e.u.data, e.v.data)
	}
	fmt.Println()
}

//BredthFirstTraversal - as the name suggests
func BredthFirstTraversal(r *Node) {
	fmt.Print("Printing BFT:")

	var q []*Node
	q = append(q, r)
	r.visited = true

	for len(q) > 0 {
		n := q[0]
		fmt.Print("\t", n.data)

		for _, a := range n.adj {
			if a.visited {
				continue
			}

			q = append(q, a)
			a.visited = true
		}
		q = q[1:]
	}

	fmt.Println()
}
