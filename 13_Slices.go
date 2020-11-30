package main

import (
	"fmt"
)

func main() {
	var x [5] int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3] // when we do not specify the size on the LHS of array then it is slice

	fmt.Println(s)
	fmt.Println(len(s)) // get length
	fmt.Println(cap(s)) // get capacity


	// creating a new Slice
	var a []int = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	//adding elements to slice
	a = append(a, 10)
	fmt.Println(a)
}

