package main

import "fmt"

func CalculateSum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func AddAndMultiply(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	multiply := num1 * num2
	return sum, multiply
}

func SayHello(name string) {
	fmt.Printf("Welcome to programming world! %s\n", name)
}

func main() {

	// sum := CalculateSum(10, 20)
	// fmt.Println(sum)

	// sum, multiply := AddAndMultiply(10, 20)
	// fmt.Println(sum, multiply)

	// SayHello("Mike")

}
