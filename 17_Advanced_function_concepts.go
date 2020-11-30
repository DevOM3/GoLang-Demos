package main

import (
    "fmt"
)

func test(){
    fmt.Println("Hello")
}


// function taking function as parameter 
func test2(myFunc func(int) int) {
    fmt.Println(myFunc(7))
}




// function closure
// function returning function 
func retWala(x int) func() {
    return func() {
        fmt.Println("Returned function")
    }
}




func main() {
    x := test // referencing a function 
    x()    

    // creating function inside function 
    fun := func (x int) int {
        fmt.Println("Hello inside")
        return x
    }

    fun(1)

    // new way to call a functio inside function
    tst := func (x int) int {
        return x*x
    }(2)
    fmt.Println(tst)

    // calling  function taking function as parameter
    test2(fun)


    // instant function call
    func(){
        fmt.Println("Test fun")
    }()

    // calling a funtion returning function
    retWala(0)()
}


