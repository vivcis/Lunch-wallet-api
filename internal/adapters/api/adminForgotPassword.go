package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

func (u HTTPHandler) AdminForgotPassword(c *gin.Context) {
	var forgotPassword models.ForgotPassword

	err := c.ShouldBindJSON(&forgotPassword)
	if err != nil {
		helpers.JSON(c, "please fill all fields", 400, nil,
			[]string{"unable to bind request: validation error"})
		return
	}
	admin, berr := u.UserService.FindAdminByEmail(forgotPassword.Email)
	if berr != nil {
		helpers.JSON(c, "user not found", 404, nil, []string{"error: user not found"})
		return
	}
	link := "http://localhost:8080/api/v1/user/adminresetpassword/" + admin.ID
	body := "Here is your reset <a href='" + link + "'>link</a>"
	html := "<strong>" + body + "</strong>"

	//initialize email sent out
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")

	sendErr := u.MailerService.SendMail("forgot password", html, admin.Email, privateAPIKey, yourDomain)
	if sendErr != nil {
		log.Println(sendErr)
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}
	helpers.JSON(c, "message: please check your email for password reset link", 200, nil,
		[]string{"message: please check your email for password reset link"})
}

func (u HTTPHandler) AdminResetPassword(c *gin.Context) {
	var reset models.ResetPassword
	err := c.ShouldBindJSON(&reset)
	if err != nil {
		helpers.JSON(c, "unable to bind json", 400, nil,
			[]string{"unable to bind request: validation error"})
		return
	}
	if reset.NewPassword != reset.ConfirmNewPassword {
		helpers.JSON(c, "password mismatch", 400, nil,
			[]string{"password mismatch"})
		return
	}
	id := c.Param("id")

	newPasswordHash, passErr := bcrypt.GenerateFromPassword([]byte(reset.NewPassword), bcrypt.DefaultCost)
	if passErr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}
	_, Rerr := u.UserService.AdminResetPassword(id, string(newPasswordHash))
	if Rerr != nil {
		helpers.JSON(c, "internal server error, please try again", 500, nil,
			[]string{"error: internal server error, please try again"})
		return
	}

	helpers.JSON(c, "message: password reset successful", 200, nil,
		[]string{"message: password reset successful"})
}
