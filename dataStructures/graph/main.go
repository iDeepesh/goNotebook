package main

import (
	"fmt"

	"github.com/iDeepesh/goNotebook/dataStructures/graph/adjecencylist"
)

func main() {
	g := adjecencylist.CreateRandomGraph(10)
	adjecencylist.DepthFirstTraversal(g, 0)
	adjecencylist.Reset(g)
	adjecencylist.BredthFirstTraversal(g[0])

	fmt.Println()
	cg := adjecencylist.CreateCompleteGraph(10)
	adjecencylist.DepthFirstTraversal(cg, 0)
	adjecencylist.Reset(cg)
	adjecencylist.BredthFirstTraversal(cg[0])
}
