package interfaces

import (
	"fmt"
	"sort"
)

type persons []string

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j int) bool {
	return p[i] <= p[j]
}

func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//ExecutePersonSort - an example of sorting a struct
func ExecutePersonSort() {
	fmt.Println("Inside interfaces.ExecutePersonSort")
	defer fmt.Println("Completed interfaces.ExecutePersonSort")

	fmt.Println("Sorting in ascending order")
	people := persons{"Deepesh", "Neelima", "Adhya", "Bhavya", "Arya"}
	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)

	fmt.Println("Sorting in descending order")
	peeps := persons{"Deepesh", "Neelima", "Adhya", "Bhavya", "Arya"}
	fmt.Println(peeps)
	sort.Sort(sort.Reverse(peeps))
	fmt.Println(peeps)
}

//ExecuteStringSort - an example of sorting a slice of strings
func ExecuteStringSort() {
	fmt.Println("Inside interfaces.ExecuteStringSort")
	defer fmt.Println("Completed interfaces.ExecuteStringSort")

	fmt.Println("Sorting in ascending order")
	people := []string{"Deepesh", "Neelima", "Adhya", "Bhavya", "Arya"}
	fmt.Println(people)
	sort.Strings(people)
	fmt.Println(people)

	fmt.Println("Sorting in descending order")
	peeps := []string{"Deepesh", "Neelima", "Adhya", "Bhavya", "Arya"}
	fmt.Println(peeps)
	sort.Sort(sort.Reverse(sort.StringSlice(peeps)))
	fmt.Println(peeps)
}

//ExecuteIntSort - an example of sorting a slice of integers
func ExecuteIntSort() {
	fmt.Println("Inside interfaces.ExecuteIntSort")
	defer fmt.Println("Completed interfaces.ExecuteIntSort")

	fmt.Println("Sorting in ascending order")
	numbers := []int{82, 293, 28, 56, 10, 47, 50, 26, 98, 32, 46}
	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)

	fmt.Println("Sorting in descending order")
	nums := []int{82, 293, 28, 56, 10, 47, 50, 26, 98, 32, 46}
	fmt.Println(nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
}
