package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/iDeepesh/goNotebook/dataStructures/list/skip"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Creating Skiplist. What is the max number of levels in list: ")
	s.Scan()
	ml, _ := strconv.Atoi(s.Text())

	sl := skip.NewSkipList(ml)

	fmt.Print("How many elements to add in list: ")
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := 0; i < n; i++ {
		sl.Insert(r.Intn(1000))
	}

	sl.Print()

	for {
		fmt.Print("What do you want to do - find (f), insert (i), delete (d) an element or quit (q)?")
		s.Scan()
		selection := s.Text()
		if strings.Compare(selection, "q") == 0 {
			break
		}
		switch selection {
		case "f":
			fmt.Print("Enter the value of the element to find?")
			s.Scan()
			f, _ := strconv.Atoi(s.Text())
			if sl.Find(f) == nil {
				fmt.Println("Can't fine the element with value of", f)
			} else {
				fmt.Println("Found an element with value of", f)
			}
		case "i":
			fmt.Print("Enter the value of the element to insert:")
			s.Scan()
			i, _ := strconv.Atoi(s.Text())
			if sl.Insert(i) == nil {
				fmt.Println("There is already an element with value of", i)
			} else {
				fmt.Println("Inserted the element with value of", i)
			}
		case "d":
			fmt.Print("Enter the value of the element to delete: ")
			s.Scan()
			d, _ := strconv.Atoi(s.Text())
			if sl.Delete(d) == nil {
				fmt.Println("Can't fine the element with value of", d)
			} else {
				fmt.Println("Found & deleted an element with value of", d)
			}
		}
		sl.Print()
	}
}
