package persons

import (
	"fmt"
)

// Xyz - any struct
type Xyz struct {
	Abc string
}

//SimpleName - struct with just first and last names
type SimpleName struct {
	First string
	Last  string
}

//FullName - struct that adds middle name to first/last that it gets from SimpleName. Contains unexported field initials
type FullName struct {
	Name     SimpleName
	Middle   string
	initials string
}

//Person - struct with full name and age
type Person struct {
	Name FullName
	Age  int
}

//SetInitials - exported function to set unexported value in FullName struct
func (fn FullName) SetInitials(in string) {
	fn.initials = in
}

//PrintPerson - prints the person struct
func (p Person) PrintPerson() {
	fmt.Println("Person printing inside of persons package, note missing initials field:", p)
}
