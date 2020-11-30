package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a number")
	scanner.Scan()
	val,_ := strconv.ParseInt(scanner.Text(), 10, 64)

	// first method
	switch val {
	case 1:
		fmt.Println("1")
		break
	case 2, 21:
		fmt.Println("2")
		break
	default:
		fmt.Println("default")
	}

	switch{
		case val < 10:
			fmt.Println("less than 10")
			break
		case val > 10 && val <21:
			fmt.Println("Between 10 to 21")
			break
		default:
			fmt.Println("Greater than 21")
	}
}

