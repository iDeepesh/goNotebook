package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iDeepesh/goNotebook/dataStructures/graph/adjecencylist"
)

func main() {
	fmt.Println("Welcome to graphs demonstration.....")
	fmt.Print("Please enter the no of nodes for sample graph:")

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, e := strconv.Atoi(s.Text())
	if e != nil {
		fmt.Println("Please enter a valid integer. Error:", e)
		return
	}

	fmt.Print("Do you want to generate a complete graph or a graph with random edges? Enter c for complete and r for random:")
	s.Scan()
	t := s.Text()
	var g []*adjecencylist.Node
	if strings.Compare(t, "c") == 0 {
		g = adjecencylist.CreateCompleteGraph(n)
	} else if strings.Compare(t, "r") == 0 {
		g = adjecencylist.CreateRandomGraph(n)
	} else {
		fmt.Println("Please enter a valid graph type to generate.")
	}

	fmt.Print("Traversal or find shortest path between two nodes. Choose t (traversal) or s (shortest path):")
	s.Scan()
	action := s.Text()
	if strings.Compare(action, "t") == 0 {
		fmt.Print("Please enter the index of node to start the traversal from:")
		s.Scan()
		r, e := strconv.Atoi(s.Text())
		if e != nil {
			fmt.Println("Please enter a valid integer. Error:", e)
			return
		} else if r < 0 || r > n-1 {
			fmt.Println("Please enter a valid index.")
			return
		}

		adjecencylist.DepthFirstTraversal(g, r)
		adjecencylist.BredthFirstTraversal(g, r)
	} else if strings.Compare(action, "s") == 0 {
		fmt.Print("Type the index of the from node:")
		s.Scan()
		from, e := strconv.Atoi(s.Text())
		if e != nil {
			fmt.Println("Please enter a valid integer. Error:", e)
			return
		} else if from < 0 || from > n-1 {
			fmt.Println("Please enter a valid index.")
			return
		}

		adjecencylist.PrintShortestPath(g, from)
	} else {
		fmt.Println("Please enter a valid action.")
	}
}
