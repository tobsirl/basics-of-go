package main

import "frontendmasters.com/go/server/data"

func main() {
	paul := data.Instructor{Id: 1, FirstName: "John", LastName: "Doe"}

	simon := data.NewInstructor("Simon", "Holmes", 100)
	keith := data.NewInstructor("Keith", "Powell", 100)

	println(paul.Print())
	println(simon.Print())
	println(keith.Print())
}