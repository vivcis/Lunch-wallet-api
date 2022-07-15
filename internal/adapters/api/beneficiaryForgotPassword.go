package api

import (
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func (u HTTPHandler) FoodBeneficiaryForgotPassword(c *gin.Context) {
	var forgotPassword models.ResetPasswordRequest

	err := c.ShouldBindJSON(&forgotPassword)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please fill all fields"})
		return
	}
	beneficiary, berr := u.UserService.FindFoodBenefactorByEmail(forgotPassword.Email)
	if berr != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	link := "http://localhost:8080/api/v1/user/beneficiaryresetpassword/" + beneficiary.ID
	body := "Here is your reset <a href='" + link + "'>link</a>"
	html := "<strong>" + body + "</strong>"

	//initialize email sent out
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")

	sendErr := u.MailerService.SendMail("forgot password", html, beneficiary.Email, privateAPIKey, yourDomain)

	//if email was sent return 200 status code
	if sendErr == nil {
		c.JSON(200, gin.H{"message": "please check your email for password reset link"})
		return
	} else {
		log.Println(sendErr)
		c.JSON(500, gin.H{"message": "something went wrong while trying to send you a mail, please try again"})
		c.Abort()
		return
	}
}

func (u HTTPHandler) FoodBeneficiaryResetPassword(c *gin.Context) {}
