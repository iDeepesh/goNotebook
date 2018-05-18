package interfaces

import (
	"fmt"
	"sort"
)

type person []string

func (p person) Len() int {
	return len(p)
}

func (p person) Less(i, j int) bool {
	return p[i] <= p[j]
}

func (p person) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ExecutePersonSort() {
	fmt.Println("Inside interfaces.ExecutePersonSort")
	defer fmt.Println("Completed interfaces.ExecutePersonSort")

	fmt.Println("Sorting in ascending order")
	people := person{"Deepesh", "Neelima", "Adhya", "Bhavya", "Arya"}
	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)

	fmt.Println("Sorting in descending order")
	peeps := person{"Deepesh", "Neelima", "Adhya", "Bhavya", "Arya"}
	fmt.Println(peeps)
	sort.Sort(sort.Reverse(peeps))
	fmt.Println(peeps)
}

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
