package main

import (
	"fmt"
)

func test(x int) {
	fmt.Println("hi", x)
}

// if both actual arguments are of same datatypes then we can specify them at once also 
func add(x, y int) {
	fmt.Println(x + y)
}

// to return we have to specify hat we want to return
func addRet(x, y int) int {
	return (x + y)
}

// returning multiple things
func ret2(x int, y float32) (int , float32) {
	return x, y
}

// auto returning function
func autoRet(x int, y int) (z, z1 int) {
	z = x + y
	z1 = x - y
	return
}

// using defer
func defferUse() {
	defer fmt.Println("Hello")  // defer will execute as soon as the last statement of the function is encountered
	fmt.Println("Last Statement")
}

func main() {
	test(11)
	add(1, 2)
	fmt.Println(addRet(1, 2))
	fmt.Println(ret2(1, 21.1))
	fmt.Println(autoRet(1, 2))
	defferUse()
}

