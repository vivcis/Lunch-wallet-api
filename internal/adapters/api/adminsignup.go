package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"os"
)

// CreateUser godoc
// @Summary      Create User
// @Description  creates a user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param user body models.Admin true "Add user"
// @Success      201  {string}  string "success"
// @Failure      400  {string}  string "error"
// @Failure      404  {string}  string "error"
// @Failure      500  {string}  string "error"
// @Router       /user/adminsignup [post]
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

	validPassword := user.IsValid(user.Password)
	if !validPassword {
		helpers.JSON(c, "Enter strong password", 400, nil, []string{"password must have upper, lower case, number, special character and length not less than 8 characters"})
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

// VerifyEmail godoc
// @Summary      Verify Email
// @Description  verifies an admin email
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param token path string true "Token string"
// @Success      200  {string}  string "success"
// @Failure      400  {string}  string "error"
// @Failure      404  {string}  string "error"
// @Failure      500  {string}  string "error"
// @Router       /user/adminverifyemail [patch]
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
