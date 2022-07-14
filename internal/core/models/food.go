package models

import "time"

type Food struct {
	Model
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	AdminName string    `json:"adminName"`
	Date      time.Time `json:"date"`
}
