package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/iDeepesh/goNotebook/dataStructures/list/skip"
)

func main() {
	fmt.Println("Skiplist demonstration")
	sl := skip.NewSkipList(5)

	// s := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(s)

	for i := 0; i < 20; i++ {
		sl.Insert(i) //r.Intn(100)) //math.MaxInt32 - 1))
	}

	sl.Print()

	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Which element to delete: ")
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	sl.Delete(n)
	sl.Print()
}
