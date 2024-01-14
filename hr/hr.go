package hr

type Person struct {
	Name string
	Age  int
}

// ## struct Embeding and anoymous field
type Employee struct {
	Person
	Designation string
}

// type name and object if lower case for local , upper case for global
