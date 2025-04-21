package main

import "fmt"

// standard or named function
func add(a, b int) int {
	return a + b
}

func main() {
	// anonymous function
	// immediately invoked function expression
	// IIFE
	func(a, b int) {
		c := a + b
		fmt.Println(c)
	}(10, 20)
}
