package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iDeepesh/goNotebook/dataStructures/trie/retrieve"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	t := retrieve.NewTrie()

	stop := false
	for !stop {
		fmt.Print("Enter (a) Add, (c) Check, (d) delete or (q) Quit: ")
		s.Scan()
		switch s.Text() {
		case "q":
			stop = true
		case "a":
			fmt.Print("Please enter the lowercase string to add to Trie: ")
			s.Scan()
			t.Add(strings.ToLower(s.Text()))
			t.Print()
		case "c":
			fmt.Print("Please enter the lowercase string to check in Trie: ")
			s.Scan()
			fmt.Println("Checked the trie for string. Status found: ", t.IsMember(strings.ToLower(s.Text())))
		case "d":
			fmt.Print("Please enter the lowercase string to delete from Trie: ")
			s.Scan()
			t.Remove(strings.ToLower(s.Text()))
			t.Print()
		}
	}
}
