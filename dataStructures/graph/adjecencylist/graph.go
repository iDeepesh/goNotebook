package adjecencylist

import (
	"fmt"
)

//Node - struct fo Adjacency List for graphs
type Node struct {
	data    int
	adj     []*Node
	visited bool
}

//CreateGraph - creates an example graph
func CreateGraph() []*Node {
	var aList []*Node
	for i := 1; i < 11; i++ {
		aList = append(aList, &Node{i, nil, false})
	}

	aList[0].adj = []*Node{aList[2], aList[3], aList[5]}
	aList[1].adj = []*Node{aList[3], aList[4], aList[6]}
	aList[2].adj = []*Node{aList[4], aList[5], aList[7]}
	aList[3].adj = []*Node{aList[5], aList[6], aList[8]}
	aList[4].adj = []*Node{aList[6], aList[7], aList[9]}
	aList[5].adj = []*Node{aList[7], aList[8], aList[0]}
	aList[6].adj = []*Node{aList[8], aList[9], aList[1]}
	aList[7].adj = []*Node{aList[9], aList[0], aList[2]}
	aList[8].adj = []*Node{aList[0], aList[1], aList[3]}
	aList[9].adj = []*Node{aList[1], aList[2], aList[4]}

	return aList
}

//Reset - remove any traversal data to reset the graph
func Reset(g []*Node) {
	for _, n := range g {
		n.visited = false
	}
}

//DepthFirstTraversal - as the name suggests
func DepthFirstTraversal(r *Node) {
	fmt.Print("Printing DFT:")

	var dfs func(*Node)
	dfs = func(n *Node) {
		if n.visited {
			return
		}

		fmt.Print("\t", n.data)

		n.visited = true
		for _, a := range n.adj {
			dfs(a)
		}
	}

	dfs(r)
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
