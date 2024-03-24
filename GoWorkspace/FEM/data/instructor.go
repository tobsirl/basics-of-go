package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func (i Instructor) GetId() int {
	return i.Id
}

func (i Instructor) GetFirstName() string {
	return i.FirstName
}

func (i Instructor) GetLastName() string {
	return i.LastName
}

func (i Instructor) GetScore() int {
	return i.Score
}

func (i Instructor) Print() string {
	return fmt.Sprintf("Id: %d, FirstName: %s, LastName: %s, Score: %d", i.Id, i.FirstName, i.LastName, i.Score)
}

func NewInstructor(name string, lastname string, score int) Instructor {
	return Instructor{ FirstName: name, LastName: lastname, Score: score }
}
