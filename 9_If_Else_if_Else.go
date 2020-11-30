package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("Enter your age")
    scanner.Scan()
    age,_ := strconv.ParseInt(scanner.Text(), 10, 64)
    

    // if-else if-else syntax
    if age == 18 {
        fmt.Println("You are 18")
    } else if age < 18 {
        fmt.Println("You are underage")
    } else{
        fmt.Println("You are an adult")
    }

}
