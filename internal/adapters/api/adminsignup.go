package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"os"
)

func (u *HTTPHandler) AdminSignUp(c *gin.Context) {
	var user *models.Admin
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helpers.JSON(c, "Unable to bind request", 400, nil, []string{"unable to bind request: validation error"})
		return
	}

	validDecagonEmail := user.ValidAdminDecagonEmail()
	if !validDecagonEmail {
		helpers.JSON(c, "Enter valid decagon email", 400, nil, []string{"enter valid decagon email"})
		return
	}

	_, Emailerr := u.UserService.FindAdminByEmail(user.Email)
	if Emailerr == nil {
		helpers.JSON(c, "Email already exists", 400, nil, []string{"email exists"})
		return
	}
	if err = user.HashPassword(); err != nil {
		helpers.JSON(c, "Unable to hash password", 400, nil, []string{"unable to hash password"})
		return
	}
	_, err = u.UserService.CreateAdmin(user)
	if err != nil {
		helpers.JSON(c, "Unable to create user", 400, nil, []string{"unable to create user"})
		return
	}
	secretString := os.Getenv("JWT_SECRET")
	emailToken, _ := u.MailerService.GenerateNonAuthToken(user.Email, secretString)
	emailLink := os.Getenv("ADMIN_EMAIL")
	link := emailLink + *emailToken
	body := "Click this <a href='" + link + "'>link</a> to verify your email."
	html := "<strong>" + body + "</strong>"

	//initialize email sent out
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")
	sendErr := u.MailerService.SendMail("Email verification", html, user.Email, privateAPIKey, yourDomain)
	if sendErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}
	helpers.JSON(c, "Please check your email to verify your account", 201, nil, nil)

}

func (u *HTTPHandler) AdminVerifyEmail(c *gin.Context) {
	token := c.Query("token")
	secretString := os.Getenv("JWT_SECRET")
	userEmail, userr := u.MailerService.DecodeToken(token, secretString)
	if userr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}
	admin, err := u.UserService.FindAdminByEmail(userEmail)
	if err != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}

	_, foodErr := u.UserService.AdminEmailVerification(admin.ID)
	if foodErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error,please try again"})
		return
	}
	helpers.JSON(c, "Congratulations, your email is now verified", 200, nil,
		[]string{"Congratulations, your email is now verified"})
}
