package models

type Food struct {
	Model
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	AdminName string  `json:"adminName"`
	Year      int     `json:"year"`
	Month     int     `json:"month"`
	Day       int     `json:"day"`
	Weekday   string  `json:"weekday"`
	Status    string  `json:"status"`
	Images    []Image `json:"images" gorm:"many2many:image"`
	Kitchen   string  `json:"kitchen"`
}

type Image struct {
	Model
	ProductId string `json:"product_id"`
	Url       string `json:"url"`
}
