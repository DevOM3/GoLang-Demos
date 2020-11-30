package main

import (
    "fmt"
)


func main() {
    var arr [3] int 

    arr[0] = 21
    fmt.Println(arr)

    arr1 := [4]int{1, 2, 3, 4}
    fmt.Println(arr1)

    arr2d := [2][2]int{{1, 2}, {2, 1}}
    fmt.Println(arr2d)
    fmt.Println(len(arr2d))
}

