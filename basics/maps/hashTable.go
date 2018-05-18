package maps

import (
	"bufio"
	"fmt"
	"github.com/iDeepesh/goNotebook/basics/io"
	"strings"
)

func ExecuteHashTableFromText()  {
	fmt.Println("Inside maps.ExecuteHashTableFromText")
	defer fmt.Println("Completed maps.ExecuteHashTableFromText")
	page := io.DoHttpGet("http://www.gutenberg.org/files/1342/1342-0.txt")
	createHashTableCounter(page)
	createHashTable(page)
}

func createHashTableCounter(text string) {
	// hT := map[int]int{}
	hT := make(map[int]int)
	for i:=0; i<10; i++{
		hT[i] = 0
	}
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		k := hashWord(scanner.Text())
		hT[k]++
	}
	fmt.Println(hT)
}

func createHashTable(text string) {
	// hT := map[int]map[string]int{}
	hT := make(map[int]map[string]int)
	for i:=0; i<12; i++{
		// hT[i] = map[string]int{}
		hT[i] = make(map[string]int)
	}
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		w := scanner.Text()
		k := hashWord(w)
		hT[k][w]++
	}

	var mostUsedWord string
	var mostUsedCount int
	for _,v := range hT{
		for key,value := range v{
			if value > mostUsedCount{
				mostUsedCount = value
				mostUsedWord = key
			}
		}
	}
	fmt.Printf("Most used work is \"%s\". Used at %d number of times.\n", mostUsedWord, mostUsedCount)
}

func hashWord(w string) int {
	var n int
	for r := range w{
		n =+ r
	}
	return n%10
}