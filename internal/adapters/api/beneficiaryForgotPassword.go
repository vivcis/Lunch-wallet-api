package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

func (u HTTPHandler) FoodBeneficiaryForgotPassword(c *gin.Context) {
	var forgotPassword models.ForgotPassword

	err := c.ShouldBindJSON(&forgotPassword)
	if err != nil {
		helpers.JSON(c, "please fill all fields", 400, nil,
			[]string{"unable to bind request: validation error"})
		return
	}
	beneficiary, berr := u.UserService.FindFoodBenefactorByEmail(forgotPassword.Email)
	if berr != nil {
		helpers.JSON(c, "user not found", 404, nil, []string{"error: user not found"})
		return
	}
	link := "http://localhost:8080/api/v1/user/beneficiaryresetpassword/" + beneficiary.ID
	body := "Here is your reset <a href='" + link + "'>link</a>"
	html := "<strong>" + body + "</strong>"

	//initialize email sent out
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")

	sendErr := u.MailerService.SendMail("forgot password", html, beneficiary.Email, privateAPIKey, yourDomain)
	if sendErr != nil {
		log.Println(sendErr)
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}
	helpers.JSON(c, "message: please check your email for password reset link", 200, nil,
		[]string{"message: please check your email for password reset link"})
}

func (u HTTPHandler) FoodBeneficiaryResetPassword(c *gin.Context) {
	var reset models.ResetPassword
	err := c.ShouldBindJSON(&reset)
	if err != nil {
		helpers.JSON(c, "unable to bind json", 400, nil,
			[]string{"unable to bind request: validation error"})
		return
	}
	if reset.NewPassword != reset.ConfirmNewPassword {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "password mismatch"})
		return
	}
	id := c.Param("id")

	//user, userErr := u.UserService.FindUserById(id)
	//if userErr != nil {
	//	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
	//	return
	//}
	newPasswordHash, passErr := bcrypt.GenerateFromPassword([]byte(reset.NewPassword), bcrypt.DefaultCost)
	if passErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	_, Rerr := u.UserService.UserResetPassword(id, string(newPasswordHash))
	if Rerr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "password reset successful"})
}
