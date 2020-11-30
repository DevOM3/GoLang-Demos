package main

import (
	"fmt"
)


// creating a structure 
type Point struct { //type structure name struct keyword
	x int
	y int
}


// creating an embedded structure
type Circle struct {
	radius float32
	center *Point
}




// global funtion to change the value of the X
func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	// referencing a struct
	// we have to assing value to all the member variables
	var point Point = Point{11, 21}
	fmt.Println(point)

	// if we don't want to assign value to all member variables the it is as follows
	var point1 Point = Point{x: 21}
	fmt.Println(point1)

	// printing a single variable or assigning a value to it can be done using a . (dot) operator
	point1.x = 11
	point1.y = 10
	fmt.Println(point1)

	// creating a pointer for a structure 
	p := &Point{1, 2}
	fmt.Println(p)
	changeX(p)
	fmt.Println(p)
	p.y = 21 // changing value directly with pointer reference
	fmt.Println(p)


	// calling an embedded structure 
	circle := Circle{3.14, &Point{1, 4}}
	fmt.Println(circle)
	fmt.Println(circle.center)
	fmt.Println(circle.center.x)
}


