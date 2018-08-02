package adjecencylist

import (
	"fmt"
)

type nodeInfo struct {
	node      *Node
	visited   bool
	arrival   int
	departure int
	distance  int
	prev      *Node
	index     int
}

type nodeInfoPQ []*nodeInfo

func (nq nodeInfoPQ) Len() int {
	return len(nq)
}

func (nq nodeInfoPQ) Less(i, j int) bool {
	return nq[i].distance < nq[j].distance
}

func (nq nodeInfoPQ) Swap(i, j int) {
	nq[i], nq[j] = nq[j], nq[i]
	nq[i].index = i
	nq[j].index = j
}

func (nq *nodeInfoPQ) Pop() interface{} {
	old := *nq
	l := len(*nq)
	x := old[l-1]
	x.index = -1
	*nq = old[:l-1]
	return x
}

func (nq *nodeInfoPQ) Push(x interface{}) {
	*nq = append(*nq, x.(*nodeInfo))
}

func (nq *nodeInfoPQ) print() {
	for _, ni := range *nq {
		fmt.Printf("Node: %d, Distance: %d, Index: %d\n", ni.node.data, ni.distance, ni.index)
	}
}
func getGraphInfo(g []*Node) map[int]*nodeInfo {
	gInfo := make(map[int]*nodeInfo, 0)
	for _, n := range g {
		gInfo[n.data] = &nodeInfo{n, false, -1, -1, -1, nil, -1}
	}
	return gInfo
}
