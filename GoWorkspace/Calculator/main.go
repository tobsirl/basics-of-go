package main

import "fmt"

func main() {
	var operation string
	var number1 int
	var number2 int

	fmt.Println("Calculator Go 1.0")
	fmt.Println("<====================>")
	fmt.Println("Enter first number: ")
	fmt.Scanln(&number1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&number2)
	fmt.Println("Enter operation: ")
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		fmt.Println(number1 / number2)
	default:
		fmt.Println("Invalid operation")
		return
	}
}
