package main

import "fmt"

func main() {
	var num int
	fmt.Print("Input any number: ")

	// Taking input from user
	fmt.Scanln(&num)

	// Validating even or odd

	if isEven(num) {
		//print that it’s even
		fmt.Println(num, "is even")
	} else {

		//print that it’s odd
		fmt.Println(num, "is odd")
	}
}

func isEven(num int) bool {
	return num%2 == 0 // If num is divisible by 2 with no remainder, it’s even.
}
