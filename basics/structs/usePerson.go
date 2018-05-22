package structs

import (
	"encoding/json"
	"fmt"
	"github.com/iDeepesh/goNotebook/basics/structs/persons"
	"os"
	"strings"
)

//ExecuteStructUsage - an example of working with nested structs
func ExecuteStructUsage() {
	fmt.Println("Inside persons.ExecuteStructUsage")
	defer fmt.Println("Completed persons.ExecuteStructUsage")
	p1 := persons.Person{}
	p1.Name.Name.First = "Jane"
	p1.Name.Name.Last = "Doe"
	p1.Name.Middle = "WhoKnows"
	p1.Age = 25
	p1.Name.SetInitials("Miss")
	fmt.Println("Person printing outside of persons package, note missing initials field:", p1)
	p1.PrintPerson()

	p2 := persons.Person{Name: persons.FullName{Name: persons.SimpleName{First: "John", Last: "Doe"}, Middle: "Unkown"}, Age: 25}
	fmt.Println(p2)
}

//ExecuteMarshaling - an example of marshaling a struct into a string and unmarshaling string to struct
func ExecuteMarshaling() {
	fmt.Println("Inside persons.ExecuteMarshaling")
	defer fmt.Println("Completed persons.ExecuteMarshaling")
	n := persons.SimpleName{First: "John", Last: "Doe"}
	bs, e := json.Marshal(n)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	fmt.Println(bs)

	s := string(bs)
	fmt.Println(s)

	var n1 persons.SimpleName
	err := json.Unmarshal(bs, &n1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(n1)
}

//ExecuteEncoding - an example of encoding a struct into json and decoding json into struct
func ExecuteEncoding() {
	fmt.Println("Inside persons.ExecuteEncoding")
	defer fmt.Println("Completed persons.ExecuteEncoding")
	n := persons.SimpleName{First: "Jane", Last: "Doe"}
	enc := json.NewEncoder(os.Stdout)
	e := enc.Encode(n)
	if e != nil {
		fmt.Println("Error in encoding", e)
		os.Exit(1)
	}

	bs, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error in marshalling", err)
		os.Exit(1)
	}

	d := json.NewDecoder(strings.NewReader(string(bs)))
	var n1 persons.SimpleName
	er := d.Decode(&n1)
	if er != nil {
		fmt.Println("Error in decoding", er)
		os.Exit(1)
	}

	fmt.Println(n1)
}
