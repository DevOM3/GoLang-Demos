package main

import (
	"fmt"
)

func main() {
	// slices are mutable datatypes
	var x []int = []int{1, 2, 3, 4, 5}
	y := x
	// here x and y points to the same location so a change from one place will result in the change for other variable also

	// map is also of this category
	var mp map[string]int = map[string]int{"om": 17}
}

