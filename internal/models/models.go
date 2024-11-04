package models

import "time"

type Daily struct {
	Id      int
	Workout string
	Date    time.Time
}
