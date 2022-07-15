package api

import (
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
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
	//err = h.Mail.SendMail("forgot Password", html, buyer.Email, privateAPIKey, yourDomain)
	sendErr := u.MailerService.SendMail("forgot password", html, beneficiary.Email, privateAPIKey, yourDomain)

}

func (u HTTPHandler) FoodBeneficiaryResetPassword(c *gin.Context) {}
