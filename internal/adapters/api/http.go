package api

import "github.com/decadevs/lunch-api/internal/ports"

type HTTPHandler struct {
	UserService   ports.UserService
	MailerService ports.MailerService
	AWSService    ports.AWSService
}

func NewHTTPHandler(userService ports.UserService, mailerService ports.MailerService, AWSService ports.AWSService) *HTTPHandler {
	return &HTTPHandler{
		UserService:   userService,
		MailerService: mailerService,
		AWSService:    AWSService,
	}
}
