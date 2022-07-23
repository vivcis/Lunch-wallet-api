package api

import (
	"os"

	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
)

// FoodBeneficiarySignUp creates a new food benefactor
func (u HTTPHandler) FoodBeneficiarySignUp(c *gin.Context) {
	var user *models.FoodBeneficiary
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helpers.JSON(c, "Unable to bind request", 400, nil, []string{"unable to bind request: validation error"})
		return
	}

	validDecagonEmail := user.ValidateDecagonEmail()
	if !validDecagonEmail {
		helpers.JSON(c, "Enter valid decagon email", 400, nil, []string{err.Error()})
		return
	}

	_, Emailerr := u.UserService.FindFoodBenefactorByEmail(user.Email)
	if Emailerr == nil {
		helpers.JSON(c, "Email already exists", 400, nil, []string{"email exists"})
		return
	}
	if err = user.HashPassword(); err != nil {
		helpers.JSON(c, "Unable to hash password", 400, nil, []string{err.Error()})
		return
	}
	_, err = u.UserService.CreateFoodBenefactor(user)
	if err != nil {
		helpers.JSON(c, "Unable to create user", 400, nil, []string{"unable to create user"})
		return
	}
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")
	link := "http://localhost:8081/api/v1/user/beneficiaryverifyemail/" + user.Token
	body := "Click this <a href='" + link + "'>link</a> to verify your email."

	sendErr := u.MailerService.SendMail("Email verification", body, user.Email, privateAPIKey, yourDomain)
	if sendErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error,please try again"})
		return
	}
	helpers.JSON(c, "Please check your email to verify your account", 201, nil, nil)

}

func (u *HTTPHandler) BeneficiaryVerifyEmail(c *gin.Context) {
	id := c.Param("id")
	_, foodErr := u.UserService.FoodBeneficiaryEmailVerification(id)
	if foodErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error,please try again"})
		return
	}
	helpers.JSON(c, "Congratulations, your email is now verified", 200, nil,
		[]string{"Congratulations, your email is now verified"})
}

// KitchenStaffSignUp creates a new kitchen staff
func (u *HTTPHandler) KitchenStaffSignUp(c *gin.Context) {
	staff := &models.KitchenStaff{}
	err := c.ShouldBindJSON(staff)
	if err != nil {
		helpers.JSON(c, "Unable to bind request", 400, nil, []string{"unable to bind request: validation error"})
		return
	}

	validDecagonEmail := staff.ValidateDecagonEmail()
	if !validDecagonEmail {
		helpers.JSON(c, "Enter valid decagon email", 400, nil, []string{err.Error()})
		return
	}

	_, err = u.UserService.FindKitchenStaffByEmail(staff.Email)
	if err == nil {
		helpers.JSON(c, "Email exist", 400, nil, []string{"email exists"})
		return
	}

	if err = staff.HashPassword(); err != nil {
		helpers.JSON(c, "Unable to hash password", 400, nil, []string{err.Error()})
		return
	}
	_, err = u.UserService.CreateKitchenStaff(staff)
	if err != nil {
		helpers.JSON(c, "Unable to create user", 400, nil, []string{err.Error()})
		return
	}
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")
	link := "http://localhost:8081/api/v1/user/kitchenstaffverifyemail/" + staff.ID
	body := "Click this <a href='" + link + "'>link</a> to verify your email."

	sendErr := u.MailerService.SendMail("Email verification", body, staff.Email, privateAPIKey, yourDomain)
	if sendErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error,please try again"})
		return
	}
	helpers.JSON(c, "Please check your email to verify your account", 201, nil, nil)
}

func (u *HTTPHandler) KitchenStaffVerifyEmail(c *gin.Context) {
	id := c.Param("id")
	_, foodErr := u.UserService.KitchenStaffEmailVerification(id)
	if foodErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error,please try again"})
		return
	}
	helpers.JSON(c, "Congratulations, your email is now verified", 200, nil,
		[]string{"Congratulations, your email is now verified"})
}

func (u HTTPHandler) AdminSignUp(c *gin.Context) {
	var user *models.Admin
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helpers.JSON(c, "Unable to bind request", 400, nil, []string{"unable to bind request: validation error"})
		return
	}

	validDecagonEmail := user.ValidateDecagonEmail()
	if !validDecagonEmail {
		helpers.JSON(c, "Enter valid decagon email", 400, nil, []string{err.Error()})
		return
	}

	_, Emailerr := u.UserService.FindAdminByEmail(user.Email)
	if Emailerr == nil {
		helpers.JSON(c, "Email already exists", 400, nil, []string{"email exists"})
		return
	}
	if err = user.HashPassword(); err != nil {
		helpers.JSON(c, "Unable to hash password", 400, nil, []string{err.Error()})
		return
	}
	_, err = u.UserService.CreateAdmin(user)
	if err != nil {
		helpers.JSON(c, "Unable to create user", 400, nil, []string{"unable to create user"})
		return
	}
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")
	link := "http://localhost:8081/api/v1/user/adminverifyemail/" + user.ID
	body := "Click this <a href='" + link + "'>link</a> to verify your email."

	sendErr := u.MailerService.SendMail("Email verification", body, user.Email, privateAPIKey, yourDomain)
	if sendErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error,please try again"})
		return
	}
	helpers.JSON(c, "Please check your email to verify your account", 201, nil, nil)

}

func (u *HTTPHandler) AdminVerifyEmail(c *gin.Context) {
	id := c.Param("id")
	_, foodErr := u.UserService.AdminEmailVerification(id)
	if foodErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error,please try again"})
		return
	}
	helpers.JSON(c, "Congratulations, your email is now verified", 200, nil,
		[]string{"Congratulations, your email is now verified"})
}
