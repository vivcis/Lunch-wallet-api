package models

import "time"

type Food struct {
	Model
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	AdminName string     `json:"adminName"`
	Year      int        `json:"year"`
	Month     time.Month `json:"month"`
	Day       int        `json:"day"`
	Weekday   string     `json:"weekday"`
	Status    string     `json:"status"`
}
