package main

import "fmt"

var x = 10

func main() {
	x := 1
	{
		x := 5
		fmt.Println("Inner x:", x)
	}
	fmt.Println("Back to global x:", x)
}
