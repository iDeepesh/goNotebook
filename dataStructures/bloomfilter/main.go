package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iDeepesh/goNotebook/datastructures/bloomfilter/bf"
)

func main() {

	b := bf.New(1024)
	s := bufio.NewScanner(os.Stdin)
	stop := false
	for !stop {
		fmt.Print("Enter (a) to add, (c) to check for an item in bloom filter. (q) to quit: ")
		s.Scan()
		switch s.Text() {
		case "q":
			stop = true
		case "a":
			fmt.Print("Please enter the string to add into the bloom filter: ")
			s.Scan()
			c := b.Add(s.Text())
			fmt.Println("Added successfully. Total number of items recorded:", c)
		case "c":
			fmt.Print("Please enter the string to check the bloom filter for: ")
			s.Scan()
			if b.Exists(s.Text()) {
				fmt.Println("The above string may be in bloom filter.")
			} else {
				fmt.Println("The above string is definitely not in bloom filter.")
			}
		}
	}
}
