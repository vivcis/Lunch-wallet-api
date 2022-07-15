package api

import "github.com/decadevs/lunch-api/internal/ports"

type HTTPHandler struct {
	UserService   ports.UserService
	MailerService ports.MailerService
}

func NewHTTPHandler(userService ports.UserService, mailerService ports.MailerService) *HTTPHandler {
	return &HTTPHandler{
		UserService:   userService,
		MailerService: mailerService,
	}
}
