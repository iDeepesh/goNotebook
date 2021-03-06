package sorts

import (
	"fmt"

	"github.com/iDeepesh/goNotebook/dataStructures/heap/minheap"
)

//RadixSort - as the name suggests
func RadixSort(a []int, m int) {
	fmt.Println("Input array: ", a)
	defer fmt.Println("Sorted array: ", a)
	d1 := 1
	d2 := 10

	for m > 0 {
		m--
		b := make([][]int, 10)
		for _, n := range a {
			i := (n % d2) / d1
			b[i] = append(b[i], n)
		}

		i := 0
		for j := range b {
			for _, n := range b[j] {
				a[i] = n
				i++
			}
		}
		d1 *= 10
		d2 *= 10
	}
}

//BubbleSort - as the name suggests
func BubbleSort(a []int) {
	fmt.Println("Input array: ", a)
	defer fmt.Println("Sorted array: ", a)

	for j := 1; j < len(a); j++ {
		for i := 0; i < len(a)-j; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
}

//SelectionSort - as the name suggests
func SelectionSort(a []int) {
	fmt.Println("Input array: ", a)
	defer fmt.Println("Sorted array: ", a)

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

//InsertionSort - as the name suggests
func InsertionSort(a []int) {
	fmt.Println("Input array: ", a)
	defer fmt.Println("Sorted array: ", a)

	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

//MergeSort - as the name suggests
func MergeSort(a []int) {
	fmt.Println("Input array: ", a)

	var mergeSort func([]int) []int
	mergeSort = func(a []int) []int {
		if len(a) == 1 {
			return a
		}

		m := len(a) / 2
		sa := mergeSort(a[:m])
		sb := mergeSort(a[m:])

		i := 0
		j := 0
		b := make([]int, len(a))
		for k := 0; k < len(b); k++ {
			if i < len(sa) && j < len(sb) {
				if sa[i] < sb[j] {
					b[k] = sa[i]
					i++
				} else {
					b[k] = sb[j]
					j++
				}
			} else if i < len(sa) {
				b[k] = sa[i]
				i++
			} else {
				b[k] = sb[j]
				j++
			}
		}
		return b
	}

	fmt.Println("Sorted array:", mergeSort(a))
}

//QuickSort - as the name suggests
func QuickSort(a []int) {
	fmt.Println("Input array: ", a)

	var quickSort func([]int)
	quickSort = func(a []int) {
		if len(a) <= 1 {
			return
		}

		q := 0
		l := len(a) - 1

		for i := 0; i < l; i++ {
			if a[i] <= a[l] {
				if i > q {
					a[i], a[q] = a[q], a[i]
				}
				q++
			}
		}
		a[l], a[q] = a[q], a[l]

		quickSort(a[:q])
		quickSort(a[q+1:])
	}

	quickSort(a)

	fmt.Println("Sorted array:", a)
}

//HeapSort - as the name suggests
func HeapSort(a []int) {
	fmt.Println("Input array: ", a)

	mH := minheap.MinHeap{}

	for _, e := range a {
		mH = mH.Add(e)
	}

	fmt.Println("Heap: ", mH)

	for i := len(mH) - 1; i >= 0; i-- {
		mH[0], mH[i] = mH[i], mH[0]
		mH.Balance(0, i-1)
	}

	fmt.Println("Sorted heap: ", mH)
}
