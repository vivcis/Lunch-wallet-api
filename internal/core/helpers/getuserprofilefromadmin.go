package helpers

import "github.com/decadevs/lunch-api/internal/core/models"

func GetUserProfileFromBeneficiary(user *models.FoodBeneficiary) models.UserProfile {
	var userprofile models.UserProfile

	userprofile.Email = user.Email
	userprofile.Stack = user.Stack
	userprofile.Avatar = user.Avatar
	userprofile.FullName = user.FullName
	userprofile.Location = user.Location

	return userprofile
}
