package main

import "fmt"

// global variable
// var url = "https://www.google.com"

// // functions receiving reference
// func calculateTax(price float32) (float32, float32) {

// 	return price * 0.075, price * 0.025
// }

func birthday(age *int)  {
	fmt.Printf("The pointer is %v and the value is %d\n", age, *age)
	if (*age>140) {
		panic("Too old to be true!")
	}
	*age++
}

func main() {
	// stateTax, _ := calculateTax(100)


	// println(stateTax)

	// println("Hello World!")
	// var x int = 5
	// println(x)
	// var name string = "Paul"
	// println(name)
	// const y = 10
	// println(y)

	// // function scope
	// var text string = "Hello"
	// println(text)

	// otherText := "World"
	// println(otherText)

	// otherNumber := 10
	// println(otherNumber)

	// println(url)

	// PrintData()

	// id := 1
	// firstName := "Paul"
	// defer fmt.Println("Bye!")

	// age := 25
	// birthday(&age)
	// birthday(&age)
	// birthday(&age)
	// fmt.Println(age)
	// println(id, firstName, age)

	// Collections
	// var numbers [5]int

	// Slices: similar to a dynamic length array, but they are actually chunks of an array
	// var names []string

	// Control Structures
	// if else
	if user := "Paul"; user == "Paul" {
		println("Hello Paul!")
	} else {
		println("Hello Stranger!")
		
	}

	day := "Monday"

	// switch
	switch day {
	case "Monday":
		println("It's Monday!")
	case "Tuesday":
		println("It's Tuesday!")
	default:
		println("It's not Monday or Tuesday!")
	} 

	switch {
	case day == "Monday":

	case day == "Tuesday":

	default:
	}

	// Classic for loop
	for i := 0; i < 10; i++ {
		println(i)
	}

	// For range, similar to "for in" in other languages
	for index := range []string{"Paul", "John", "Mary"} {
		println(index)
	}

	// For range, similar to "foreach" in other languages
	for key, value := range map[string]string{"name": "Paul", "age": "25"} {
		println(key, value)
	}
}

// func save() {
// 	println("Saving data...")
// }

// func saveString(text string) {
// 	println(text)
// }

// func add (x int, y int) int {
// 	return x + y
// }

// func addAndSubtract (x int, y int) (int, int) {
// 	return x + y, x - y
// }

