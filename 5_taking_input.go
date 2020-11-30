package main

// importing multiple packages as follows
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // creating a scanner
	fmt.Printf("When you were born (year) :- ")
	scanner.Scan()                                       // waiting
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) // getting the text and parsing it to integer of base 10 and 64 bit, it returns two values
	fmt.Printf("You are %d years Old", 2020-input)
}
