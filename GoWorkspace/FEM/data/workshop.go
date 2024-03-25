package data

import "time"

type Workshop struct {
	Course // embedding
	Date  time.Time
}

func NewWorkShop(name string, instructor Instructor) Workshop {
	return Workshop{
		Course: Course{
			Name:       name,
			Instructor: instructor,
		},
		Date: time.Now(),
	}
}