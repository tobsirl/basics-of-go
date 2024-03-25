package data

import "time"

type Workshop struct {
	Course // embedding
	Date  time.Time
}