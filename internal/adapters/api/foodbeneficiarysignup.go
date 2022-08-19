package api

import (
	"os"

	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary      Create User
// @Description  creates a user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param user body models.FoodBeneficiary true "Add user"
// @Success      201  {string}  string "success"
// @Failure      400  {string}  string "error"
// @Failure      404  {string}  string "error"
// @Failure      500  {string}  string "error"
// @Router       /user/beneficiarysignup [post]
func (u *HTTPHandler) FoodBeneficiarySignUp(c *gin.Context) {
	user := &models.FoodBeneficiary{}
	if err := u.decode(c, user); err != nil {
		helpers.JSON(c, "", 400, nil, err)
		return
	}

	validDecagonEmail := user.ValidateDecagonEmail()
	if !validDecagonEmail {
		helpers.JSON(c, "Enter valid decagon email", 400, nil, []string{"enter valid decagon email"})
		return
	}

	_, Emailerr := u.UserService.FindFoodBenefactorByEmail(user.Email)
	if Emailerr == nil {
		helpers.JSON(c, "Email already exists", 400, nil, []string{"email exists"})
		return
	}

	validPassword := user.IsValid(user.Password)
	if !validPassword {
		helpers.JSON(c, "Enter strong password", 400, nil, []string{"password must have upper, lower case, number, special character and length not less than 8 characters"})
		return
	}

	if err := user.HashPassword(); err != nil {
		helpers.JSON(c, "Unable to hash password", 400, nil, []string{err.Error()})
		return
	}

	_, err := u.UserService.CreateFoodBenefactor(user)
	if err != nil {
		helpers.JSON(c, "Unable to create user", 400, nil, []string{"unable to create user"})
		return
	}

	secretString := os.Getenv("JWT_SECRET")
	emailToken, _ := u.MailerService.GenerateNonAuthToken(user.Email, secretString)
	emailLink := os.Getenv("BENEFICIARY_EMAIL")
	link := emailLink + *emailToken
	body := "Click this <a href=' " + " " + link + " '>link</a> to verify your email."
	html := "<strong> " + body + " </strong>"

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
// @Description  verifies a food beneficiary email
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param token path string true "Token string"
// @Success      200  {string}  string "success"
// @Failure      400  {string}  string "error"
// @Failure      404  {string}  string "error"
// @Failure      500  {string}  string "error"
// @Router       /user/beneficiaryverifyemail [patch]
func (u *HTTPHandler) BeneficiaryVerifyEmail(c *gin.Context) {
	token := c.Query("token")
	secretString := os.Getenv("JWT_SECRET")
	userEmail, uerr := u.MailerService.DecodeToken(token, secretString)
	if uerr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}
	beneficiary, err := u.UserService.FindFoodBenefactorByEmail(userEmail)
	if err != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}

	_, foodErr := u.UserService.FoodBeneficiaryEmailVerification(beneficiary.ID)
	if foodErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error,please try again"})
		return
	}
	helpers.JSON(c, "Congratulations, your email is now verified", 200, nil,
		[]string{"Congratulations, your email is now verified"})
}
