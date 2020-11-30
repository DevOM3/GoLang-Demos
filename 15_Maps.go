package main

import (
	"fmt"
)

func main() {
	var mp map[string]int = map[string]int{
		"om": 17,
		"bhargavi": 8,
		"mummy": 37,
		"pappa": 40,
	}

	fmt.Println(mp)

	// another way to make a map
	mp1 := make(map[string]int)
	mp1["apple"] = 5
	fmt.Println(mp1)
	delete(mp, "om")
	fmt.Println(mp)


	// checking if key exists
	val,ok := mp["om"]
	fmt.Println(val, ok)
}

