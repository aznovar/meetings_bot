package models

import "time"

type Meeting struct {
	ID           int
	Title        string
	Date         time.Time
	Participants string
	Summary      string
}
