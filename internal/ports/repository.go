package ports

import (
	"errors"

	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetByID(id string) (*models.User, error)
	FindUserByFullName(fullname string) (*models.User, error)
}

type Router interface {
	GET(uri string, fn func(ctx *gin.Context))
	POST(uri string, fn func(ctx *gin.Context))
	PUT(uri string, fn func(ctx *gin.Context))
	DELETE(uri string, fn func(ctx *gin.Context))
	SERVE() error
	GROUP(path string) *gin.RouterGroup
}

// FindUserByFullName finds a user by the fullname
func (pdb *PostgresDb) FindUserByFullName(fullname string) (*models.User, error) {
	user := &models.User{}

	if err := pdb.DB.Where("username = ?", fullname).First(user).Error; err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("user inactive")
	}
	return user, nil
}