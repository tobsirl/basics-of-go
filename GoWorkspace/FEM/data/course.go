package data

import "fmt"

type Duration float32 // in hours

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (c Course) SignUp() bool {
	return true
}

func (c Course) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Slug: %s, Legacy: %t, Duration: %f, Instructor: %s", c.Id, c.Name, c.Slug, c.Legacy, c.Duration, c.Instructor.Print())
}
