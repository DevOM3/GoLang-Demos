package main

import (
	"fmt"
)


// creating a function having actual parameter as a pointer
func pointerWala(p *int) {
	*p = 21
}

func main() {
	var no int = 10 //  creating variable
	p := &no // creating a pointer to the no

	fmt.Println(p) // printing the address of the no
	fmt.Println(*p) // printing the value in pointer p

	fmt.Println(no) // value of no before
	pointerWala(&no) // passing pointer as a parameter
	fmt.Println(no) // after changing value using pointer
}
