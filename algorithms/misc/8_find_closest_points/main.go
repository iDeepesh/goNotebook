package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type point struct {
	x, y int
	d    float64
}

type theHeap []*point

func (hp *theHeap) Push(x interface{}) {
	item := x.(*point)
	*hp = append(*hp, item)
}

// By the time Pop is called, the head is already swapped to the end of the heap array. So delete last element.
func (hp *theHeap) Pop() interface{} {
	n := len(*hp)
	old := *hp
	item := old[n-1]
	*hp = old[0 : n-1]
	return item
}

func (hp theHeap) Len() int {
	return len(hp)
}

// Answers - if i should be before j
func (hp theHeap) Less(i, j int) bool {
	return hp[i].d > hp[j].d
}

func (hp theHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the number of points to generate: ")
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	pts := generatePoints(n, 100)
	fmt.Print("Please enter x-value for the point to find closest points from: ")
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	fmt.Print("Please enter y-value for the point to find closest points from: ")
	s.Scan()
	y, _ := strconv.Atoi(s.Text())
	fmt.Print("Please enter the number of closest points to find: ")
	s.Scan()
	m, _ := strconv.Atoi(s.Text())

	pt := point{x, y, 0}

	closePoints := findClosestPoints(pts, pt, m)

	maxClosest := -1.0
	for _, cp := range closePoints {
		if cp.d > maxClosest {
			maxClosest = cp.d
		}
	}

	c := 0
	for _, p := range pts {
		if p.d < maxClosest {
			c++
		}
	}
	if c > m {
		fmt.Println("Failed in finding the correct list of points. Check detailed output.")
	}

	fmt.Print("All points: ")
	printPoints(pts)
	fmt.Print("Closest points: ")
	printPointerPoints(closePoints)
}

func findClosestPoints(pts []point, pt point, m int) []*point {
	hpts := theHeap{}
	heap.Init(&hpts)

	for i := range pts {
		pts[i].calculateDistance(pt)
		if i < m {
			heap.Push(&hpts, &pts[i])
		} else {
			if pts[i].d < hpts[0].d {
				heap.Pop(&hpts)
				heap.Push(&hpts, &pts[i])
				// heap.Init(&hpts)
			}
		}
	}

	cpts := make([]*point, m)
	for i := m - 1; i >= 0; i-- {
		cpts[i] = (heap.Pop(&hpts)).(*point)
	}
	return cpts
}

func (pt *point) calculateDistance(p point) {
	x := pt.x - p.x
	y := pt.y - p.y
	pt.d = math.Sqrt(float64(x*x) + float64(y*y))
	// fmt.Printf("x2=%d, y2=%d, d2=%.2f, d=%.2f\n", x2, y2, py, pt.d)
}

func generatePoints(n, max int) []point {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	pts := make([]point, n)

	for i := 0; i < n; i++ {
		xs, ys := (r.Intn(max)&1) > 0, (r.Intn(max)&1) > 0
		x, y := r.Intn(max), r.Intn(max)
		if xs {
			x *= -1
		}
		if ys {
			y *= -1
		}
		pts[i] = point{x, y, 0}
	}

	printPoints(pts)
	return pts
}

func printPointerPoints(pts []*point) {
	fmt.Print("[")
	for _, p := range pts {
		fmt.Printf("{(%d,%d), %.2f},", p.x, p.y, p.d)
	}
	fmt.Println("]")
}

func printPoints(pts []point) {
	fmt.Print("[")
	for _, p := range pts {
		fmt.Printf("{(%d,%d), %.2f},", p.x, p.y, p.d)
	}
	fmt.Println("]")
}
