package main

import (
	"fmt"
)

// creating an Interface
type I interface {
	showData()
}

// creating a Student
type Student struct {
	age  int
	name string
}

// a Student function
func (s Student) showData() {
	fmt.Println("Age: ", s.age)
	fmt.Println("Name: ", s.name)
}

func main() {
	// calling an interface
	i := []I{Student{17, "OM"}}
	i[0].showData()
}
