package adjecencylist

import (
	"fmt"
	"math/rand"
)

//Node - struct fo Adjacency List for graphs
type Node struct {
	data int
	adj  []*Edge
}

type Edge struct {
	u *Node
	v *Node
}

type nodeInfo struct {
	visited   bool
	arrival   int
	departure int
}

//CreateCompleteGraph - as the name suggests
func CreateCompleteGraph(n int) []*Node {
	fmt.Println("Creating a complete graph with element count of", n)

	g := make([]*Node, 0)

	for i := 0; i < n; i++ {
		g = append(g, &Node{i, nil})
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

	g := make([]*Node, 0)

	for i := 0; i < n; i++ {
		g = append(g, &Node{i, nil})
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

//DepthFirstTraversal - as the name suggests
func DepthFirstTraversal(g []*Node, ri int) {
	fmt.Print("Printing Depth First Traversal:")

	gInfo := make(map[int]*nodeInfo, 0)
	for _, n := range g {
		gInfo[n.data] = &nodeInfo{false, -1, -1}
	}

	r := g[ri]
	tE := make([]*Edge, 0)
	t := 0

	var dfs func(*Node)
	dfs = func(n *Node) {
		if gInfo[n.data].visited {
			return
		}

		gInfo[n.data].visited = true
		gInfo[n.data].arrival = t
		t++
		fmt.Printf("%d, ", n.data)

		for _, a := range n.adj {
			if gInfo[a.v.data].visited {
				continue
			}
			tE = append(tE, a)
			dfs(a.v)
		}

		gInfo[n.data].departure = t
		t++
	}

	dfs(r)
	fmt.Println()

	fmt.Println("Printing the Tree Edges")
	printEdges(tE)

	fE, bE, cE := getEdgesByType(g, gInfo, tE)

	fmt.Println("Printing the Forward Edges")
	printEdges(fE)

	fmt.Println("Printing the Backward Edges")
	printEdges(bE)

	fmt.Println("Printing the Cross Edges")
	printEdges(cE)
}

// Returns - Forward edges, Backward edges, Cross edges
func getEdgesByType(g []*Node, gI map[int]*nodeInfo, tE []*Edge) ([]*Edge, []*Edge, []*Edge) {
	fE := make([]*Edge, 0)
	bE := make([]*Edge, 0)
	cE := make([]*Edge, 0)

	for _, n := range g {
		for _, e := range n.adj {
			if findEdge(tE, e) {
				continue
			}

			uI := gI[e.u.data]
			vI := gI[e.v.data]

			if uI.arrival < vI.arrival && uI.departure > vI.departure {
				fE = append(fE, e)
			} else if uI.arrival > vI.arrival && uI.departure < vI.departure {
				bE = append(bE, e)
			} else if vI.arrival < vI.departure && vI.departure < uI.arrival && uI.arrival < uI.departure {
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

	gInfo := make(map[int]*nodeInfo, 0)
	for _, n := range g {
		gInfo[n.data] = &nodeInfo{false, -1, -1}
	}

	r := g[ri]
	tE := make([]*Edge, 0)
	var q []*Node
	q = append(q, r)
	gInfo[r.data].visited = true

	for len(q) > 0 {
		n := q[0]
		fmt.Printf("%d, ", n.data)

		for _, e := range n.adj {
			if gInfo[e.v.data].visited {
				continue
			}

			q = append(q, e.v)
			tE = append(tE, e)
			gInfo[e.v.data].visited = true
		}
		q = q[1:]
	}

	fmt.Println()

	fmt.Println("Printing the Tree Edges")
	printEdges(tE)
}
