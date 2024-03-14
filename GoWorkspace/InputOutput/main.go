package main

// global variable
var url = "https://www.google.com"

func main() {
	println("Hello World!")
	var x int = 5
	println(x)
	var name string = "Paul"
	println(name)
	const y = 10
	println(y)

	// function scope
	var text string = "Hello"
	println(text)

	otherText := "World"
	println(otherText)

	otherNumber := 10
	println(otherNumber)

	println(url)

	PrintData()

	id := 1
	firstName := "Paul"
	age := 25
	println(id, firstName, age)
}