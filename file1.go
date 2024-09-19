package main

import "fmt"

// Function to add two numbers
func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(2, 3)
	fmt.Println("The sum of 2 and 3 is:", result)
}
