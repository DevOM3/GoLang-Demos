package main

import (
	"fmt"
)


func main() {
	var no1 int = 2
	var no2 int = 1
	ans := no1 + no2
	fmt.Println(ans)

	var no3 float32 = 14.21
	res := float32(no1) + no3 // converting no1 from int to float 
	fmt.Println(res)
}

