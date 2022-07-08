package ports

import "github.com/decadevs/lunch-api/internal/core/models"

type UserService interface {
	GetByID(id string) (*models.User, error)
}
