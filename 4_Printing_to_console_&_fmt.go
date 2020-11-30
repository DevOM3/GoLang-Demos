package main

import "fmt"

func main() {
	fmt.Printf("%T\n\n", 10) // %T to get value

	// boolen
	fmt.Printf("%t\n\n", true) // format true or false using %t

	// integers
	fmt.Printf("%v\n", 10)       // %v to print values
	fmt.Printf("%b\n", 100)      // %b to get binary
	fmt.Printf("%o\n", 100)      // %o to get octal
	fmt.Printf("%d\n", 100)      // to get decimal
	fmt.Printf("%x\n", 1056450)  // %x for small-hexadecimal
	fmt.Printf("%X\n\n", 167500) // %X for capital-hexadecimal

	// floats
	fmt.Printf("%e\n", 100.652653175347635)             // to get scientific notation lik1 10e-10
	fmt.Printf("%f\n", 100.52374537534)                 // %f or %F printig float
	fmt.Printf("%g\n\n", 100.6536526723625362563562353) // %g to get big exponent like double

	// Strings
	fmt.Printf("%s\n", "om")   // prints om
	fmt.Printf("%q\n\n", "om") // prints "om" (with quotations)

	// Width and precision
	fmt.Printf("%f\n", 10.65700)      // default width and precision
	fmt.Printf("%9f\n", 10.6273672)   // width 9 and default precision
	fmt.Printf("%.2f\n", 10.123)      // default width and 2 precision
	fmt.Printf("%9.2f\n", 10.89238923) // 9 width and 2 precision
	fmt.Printf("%9.f\n\n", 10.238237) //9 width and 0 precision

	// Padding
	fmt.Printf("%07d\n", 10) // to specify digit padding
	fmt.Printf("%-4f", 10.909) // 
}
