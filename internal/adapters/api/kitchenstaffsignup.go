package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

// CreateUser godoc
// @Summary      Create User
// @Description  creates a user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param staff body models.KitchenStaff true "Add user"
// @Success      201  {string}  string "success"
// @Failure      400  {string}  string "error"
// @Failure      404  {string}  string "error"
// @Failure      500  {string}  string "error"
// @Router       /user/kitchenstaffsignup [post]
func (u *HTTPHandler) KitchenStaffSignUp(c *gin.Context) {
	staff := &models.KitchenStaff{}
	err := c.ShouldBindJSON(staff)
	if err != nil {
		helpers.JSON(c, "Unable to bind request", 400, nil, []string{"unable to bind request: validation error"})
		return
	}

	if strings.Contains(staff.Email, "decagon.dev") || strings.Contains(staff.Email, "decagonhq.com") {
		helpers.JSON(c, "", 400, nil, []string{"enter a valid email"})
		return
	}
	
	validateEmail := staff.ValidateEmail()
	if !validateEmail {
		helpers.JSON(c, "Enter valid email", 400, nil, []string{"enter valid email"})
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

	secretString := os.Getenv("JWT_SECRET")
	emailToken, _ := u.MailerService.GenerateNonAuthToken(staff.Email, secretString)
	emailLink := os.Getenv("KITCHEN_STAFF_EMAIL")
	link := emailLink + *emailToken
	body := "Click this <a href='" + link + "'>link</a> to verify your email."
	html := "<strong>" + body + "</strong>"

	//initialize email sent out
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")
	sendErr := u.MailerService.SendMail("Email verification", html, staff.Email, privateAPIKey, yourDomain)
	if sendErr != nil {
		helpers.JSON(c, "internal server error, please try again again", 500, nil,
			[]string{"error: It is an internal server error, please try again"})
		return
	}

	helpers.JSON(c, "Please check your email to verify your account", 201, nil, nil)
}

// VerifyEmail godoc
// @Summary      Verify Email
// @Description  verifies a kitchen staff email
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param token path string true "Token string"
// @Success      200  {string}  string "success"
// @Failure      400  {string}  string "error"
// @Failure      404  {string}  string "error"
// @Failure      500  {string}  string "error"
// @Router       /user/kitchenstaffverifyemail/{token} [patch]
func (u *HTTPHandler) KitchenStaffVerifyEmail(c *gin.Context) {
	token := c.Param("token")
	secretString := os.Getenv("JWT_SECRET")
	userEmail, userr := u.MailerService.DecodeToken(token, secretString)
	if userr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error"})
		return
	}
	kitchenStaff, err := u.UserService.FindKitchenStaffByEmail(userEmail)
	if err != nil {
		helpers.JSON(c, "An internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}

	_, foodErr := u.UserService.KitchenStaffEmailVerification(kitchenStaff.ID)
	if foodErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error,please try again"})
		return
	}
	helpers.JSON(c, "Congratulations, your email is now verified", 200, nil,
		[]string{"Congratulations, your email is now verified"})
}
