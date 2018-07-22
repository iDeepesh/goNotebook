package main

import (
	"github.com/iDeepesh/goNotebook/dataStructures/tree/bst"
)

func main() {
	ia := []int{2, 9, 1, 6, 4, 8, 0, 3, 7, 5}
	t := bst.Tree{}

	for _, i := range ia {
		t.Add(i)
	}

	t.Print()

	t.Delete(8)
	t.Print()
}
