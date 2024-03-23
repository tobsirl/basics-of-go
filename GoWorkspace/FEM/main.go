package main

import "frontendmasters.com/go/server/data"

func main() {
	user := data.Instructor{Id: 1, FirstName: "John", LastName: "Doe"}
	user.FirstName = "Paul"



	println(user.Print())
	println(user.GetId())
	println(user.GetFirstName())
	println(user.GetLastName())
	println(user.GetScore())
}