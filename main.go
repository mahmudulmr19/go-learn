package main

import (
	"fmt"

	"example.com/mathlib"
)

var globalNumber int = 10

func main() {
	x := mathlib.Add(1, 2)
	fmt.Println(x)
	fmt.Println(globalNumber)
}

func init() {
	globalNumber = 20
}
