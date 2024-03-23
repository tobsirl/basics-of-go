package main

import "frontendmasters.com/go/server/data"

func main() {
	user := data.Instructor{}
	user.FirstName = "Paul"



	println(user)
}