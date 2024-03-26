package main

import (
	"fmt"

	"frontendmasters.com/go/server/data"
)

func main() {

	paul := data.Instructor{Id: 1, FirstName: "John", LastName: "Doe"}

	simon := data.NewInstructor("Simon", "Holmes", 100)
	keith := data.NewInstructor("Keith", "Powell", 100)

	goCourse := data.Course{Id: 1, Name: "Go Fundamentals", Slug: "go-fundamentals", Legacy: false, Duration: 3.2, Instructor: paul}

	swiftWS := data.NewWorkShop("Swift Fundamentals", simon)

	fmt.Printf("%+v\n", swiftWS)

	println(paul.Print())
	println(simon.Print())
	println(keith.Print())
	fmt.Printf("%+v\n", goCourse)
}
