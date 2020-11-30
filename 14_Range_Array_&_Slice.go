package main

import (
	"fmt"
)

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}

	// range
	// for each 
	for i, element := range a { // i will be index and element will be value which returned by range a
		fmt.Printf("%d: %d\n", i, element)
	}

	for i, element := range a {
		for j, ele := range a{
			if ele == element && i != j{
				fmt.Println("NO")
			}
		}
	}
}

