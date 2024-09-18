package main

import "fmt"

// Multiply multiplies two integers and returns the result.
func Multiply(a, b int) int {
    return a * b
}

func main() {
    fmt.Println("Multiplication:", Multiply(2, 3))
}
