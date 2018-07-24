package bst

import (
	"fmt"
	"math"
)

//Node - struct representing single tree node
type Node struct {
	value int
	left  *Node
	right *Node
}

//Tree - struct representing a tree
type Tree struct {
	r *Node
}

//Add - add a new node to tree
func (t *Tree) Add(v int) {
	n := Node{v, nil, nil}
	t.r = add(t.r, n)
}

//Delete - delete a node from tree
func (t *Tree) Delete(v int) {
	t.r = deleteNode(t.r, v)
}

//CheckBST - checks if the tree is bst or not
func (t *Tree) CheckBST() bool {
	var checkBST func(*Node, int, int) bool
	checkBST = func(r *Node, min int, max int) bool {
		if r == nil {
			return true
		}

		if r.value > max || r.value <= min {
			return false
		}

		return checkBST(r.left, min, r.value) && checkBST(r.right, r.value, max)
	}

	return checkBST(t.r, math.MinInt32, math.MaxInt32)
}

//Print - print the tree in various forms
func (t *Tree) Print() {
	var findMaxLevel func(*Node) int
	findMaxLevel = func(r *Node) int {
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

	var printLevel func([]*Node, int, int)
	printLevel = func(nList []*Node, l int, maxLevel int) {
		initalSpaces := int(math.Pow(2, float64(maxLevel-l))) - 1
		separaterSpaces := int(math.Pow(2, float64(maxLevel-l+1))) - 1

		isAllElementsNil := true

		printSpace(initalSpaces)
		newList := []*Node{}
		for _, n := range nList {
			if n != nil {
				isAllElementsNil = false
				fmt.Print(n.value)
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

	maxLevel := findMaxLevel(t.r)
	nList := []*Node{t.r}
	printLevel(nList, 1, maxLevel)
}

//PrintAllOutputs - prints all traversals, etc.
func (t *Tree) PrintAllOutputs() {
	t.Print()
	fmt.Println("Checking if tree is BST or not: ", t.CheckBST())
	fmt.Println("Printing breadth first traversal for Tree....")
	breadthFirstTraversal(t.r)
	fmt.Println()
	fmt.Println("Printing pre-order traversal for Tree....")
	preOrderTraversal(t.r)
	fmt.Println()
	fmt.Println("Printing in-order traversal for Tree....")
	inOrderTraversal(t.r)
	fmt.Println()
	fmt.Println("Printing post-order traversal for Tree....")
	postOrderTraversal(t.r)
	fmt.Println()
	findMin(t.r)
	findMax(t.r)
	fmt.Println("Finding element 8: ", findNode(t.r, 8))
}

func add(r *Node, n Node) *Node {
	if r == nil {
		return &n
	} else if n.value > r.value {
		r.right = add(r.right, n)
		return r
	} else {
		r.left = add(r.left, n)
		return r
	}
}

func inOrderTraversal(r *Node) {
	if r == nil {
		return
	}

	inOrderTraversal(r.left)
	fmt.Print("\t", r.value)
	inOrderTraversal(r.right)
}

func preOrderTraversal(r *Node) {
	if r == nil {
		return
	}

	fmt.Print("\t", r.value)
	preOrderTraversal(r.left)
	preOrderTraversal(r.right)
}

func postOrderTraversal(r *Node) {
	if r == nil {
		return
	}

	postOrderTraversal(r.left)
	postOrderTraversal(r.right)
	fmt.Print("\t", r.value)
}

func breadthFirstTraversal(r *Node) {
	type smartNode struct {
		n *Node
		l int
	}

	var q []smartNode

	l := 0
	q = append(q, smartNode{r, l})

	for len(q) > 0 {
		s := q[0]

		if s.l > l {
			fmt.Println()
			l++
		}

		if s.n.left != nil {
			q = append(q, smartNode{s.n.left, s.l + 1})
		}

		if s.n.right != nil {
			q = append(q, smartNode{s.n.right, s.l + 1})
		}

		fmt.Print("\t", s.n.value)
		q = q[1:]
	}
}

func findMin(r *Node) *Node {
	if r == nil {
		fmt.Println("Error - empty tree")
	}

	if r.left == nil {
		fmt.Println("The min value is: ", r.value)
		return r
	}

	return findMin(r.left)
}

func findMax(r *Node) *Node {
	if r == nil {
		fmt.Println("Error - empty tree")
	}

	if r.right == nil {
		fmt.Println("The max value is: ", r.value)
		return r
	}

	return findMax(r.right)
}

func findNode(r *Node, v int) *Node {
	if r == nil {
		return nil
	}

	if r.value == v {
		return r
	}

	if r.value > v {
		return findNode(r.left, v)
	} else {
		return findNode(r.right, v)
	}
}

func deleteNode(r *Node, v int) *Node {
	if r == nil {
		return nil
	}

	if r.value > v {
		r.left = deleteNode(r.left, v)
		return r
	} else if r.value < v {
		r.right = deleteNode(r.right, v)
		return r
	}

	if r.left == nil && r.right == nil {
		return nil
	} else if r.left == nil {
		return r.right
	} else if r.right == nil {
		return r.left
	} else {
		n := findMin(r.right)
		r.value = n.value
		r.right = deleteNode(r.right, n.value)
		return r
	}
}
