package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 2
	fmt.Println(x < y)
	fmt.Println(x > y)
	fmt.Println(x == y)
	fmt.Println(x != y)
    fmt.Println(x + 1 != y) 
    
    x1 := "om"
    y1 := "Om"
    fmt.Println(x1 < y1) // comparison happens by using ASCII
    fmt.Println(x1 > y1)
    fmt.Println(x1 == y1)
}
