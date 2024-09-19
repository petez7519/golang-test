package main

import "fmt"

// Function to multiply two numbers
func multiply(a int, b int) int {
	return a * b
}

func main2() {
	result := multiply(2, 3)
	fmt.Println("The product of 2 and 3 is:", result)
}
