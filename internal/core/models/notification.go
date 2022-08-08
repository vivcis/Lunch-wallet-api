package models

import "time"

type Notification struct {
	Message string     `json:"message"`
	Year    int        `json:"year"`
	Month   time.Month `json:"month"`
	Day     int        `json:"day"`
}
