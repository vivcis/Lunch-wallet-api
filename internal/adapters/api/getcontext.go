package api

import (
	"fmt"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
)

func (u HTTPHandler) GetAdminFromContext(c *gin.Context) (*models.Admin, error) {
	userI, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("error getting user from context")
	}
	user, ok := userI.(*models.Admin)
	if !ok {
		return nil, fmt.Errorf("an error occurred")
	}
	return user, nil
}

func (u HTTPHandler) GetBenefactorFromContext(c *gin.Context) (*models.FoodBeneficiary, error) {
	userI, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("error getting user from context")
	}
	user, ok := userI.(*models.FoodBeneficiary)
	if !ok {
		return nil, fmt.Errorf("an error occurred")
	}
	return user, nil
}

func (u HTTPHandler) GetKitchenStaffFromContext(c *gin.Context) (*models.KitchenStaff, error) {
	userI, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("error getting user from context")
	}
	user, ok := userI.(*models.KitchenStaff)
	if !ok {
		return nil, fmt.Errorf("an error occurred")
	}
	return user, nil
}
