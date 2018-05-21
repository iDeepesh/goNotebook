package persons

import (
	"fmt"
)

type Xyz struct {
	Abc string
}

type SimpleName struct {
	First string
	Last  string
}

type FullName struct {
	Name     SimpleName
	Middle   string
	initials string
}

type Person struct {
	Name FullName
	Age  int
}

func (fn FullName) SetInitials(in string) {
	fn.initials = in
}

func (p Person) PrintPerson() {
	fmt.Println("Person printing inside of persons package, note missing initials field:", p)
}
