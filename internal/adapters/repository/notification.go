package repository

import (
	"errors"
	"github.com/decadevs/lunch-api/internal/core/models"
	"time"
)

func (p *Postgres) CreateNotification(notification models.Notification) error {
	return p.DB.Create(&notification).Error
}

func (p *Postgres) FindNotificationDate(year int, month time.Month, day int) ([]models.Notification, error) {
	var notification []models.Notification
	if err := p.DB.Where("year = ?", year).Where("month = ?", month).Where("day = ?", day).
		Find(&notification).Error; err != nil {
		return nil, errors.New("error getting notification")
	}
	return notification, nil
}
