package main

import (
	"fmt"
)


// creating a struct 
type Student struct {
	name string
	grades []int
	age int
}

// a getter struct function
func (s Student) getAge() int {
	return s.age
}

// a setter struct function
// a function fro which we want to change the value of the object must have *Struct
func (s *Student) setAge(age int) {
	s.age = age
}


func main() {
	s1 := Student{"OM", []int{99, 97, 96, 95, 99}, 17}
	a := s1.getAge()

	fmt.Println(s1)
	fmt.Println(a)

	s1.setAge(21)
	fmt.Println(s1)

}

