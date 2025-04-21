package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// basic declaration
	// const name string = "Mahmudul Hasan"
	// var age int = 17
	// var isStudent bool = false

	// fmt.Println(name, age, isStudent)

	// int
	// var x int
	// var y string
	// var z bool
	// fmt.Println(x, y, z)

	// Short Declaration
	// sum := 5 + 3
	// message := "Go is fun!"
	// fmt.Println(sum, message)

	// Type Conversion
	// a := 10
	// b := 3.5

	// result := float64(a) * b
	// fmt.Println(result)

	// var x, y = 40, 50

	// x = x + y
	// y = x - y
	// x = x - y

	// fmt.Println(x, y)

	var score float64
	var grade string

	// Create a new reader to read input with spaces
	reader := bufio.NewReader(os.Stdin)

	// Ask for name (handles spaces properly)
	fmt.Println("What's your name?")
	studentName, _ := reader.ReadString('\n')
	studentName = strings.TrimSpace(studentName) // Remove the newline character at the end

	// Ask for score
	fmt.Println("What is your score?")
	fmt.Scan(&score)

	// Calculate grade based on score
	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	// Output the result
	fmt.Printf("%s scored %.2f and got a grade of %s\n", studentName, score, grade)

}
