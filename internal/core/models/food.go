package models

import (
	"time"
)

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
	Images    []Image    `json:"images" gorm:"oneToMany"`
	Kitchen
}

type Image struct {
	Model
	ProductId uint   `json:"product_id"`
	Url       string `json:"url"`
}

type Kitchen struct {
	Name string `json:"name"`
}
