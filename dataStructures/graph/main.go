package main

import (
	"github.com/iDeepesh/goNotebook/dataStructures/graph/adjecencylist"
)

func main() {
	g := adjecencylist.CreateGraph()

	adjecencylist.DepthFirstTraversal(g[0])

	adjecencylist.Reset(g)
	adjecencylist.BredthFirstTraversal(g[0])
}
