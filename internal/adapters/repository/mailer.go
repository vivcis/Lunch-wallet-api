package repository

import (
	"context"
	"github.com/decadevs/lunch-api/internal/ports"
	"github.com/dgrijalva/jwt-go"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

type Claims struct {
	UserEmail string `json:"email"`
	jwt.StandardClaims
}

//
//type Service struct{}

type Mail struct {
	Client *mailgun.MailgunImpl
}

//func (m Mail) SendgMail(subject, body, to, Private, Domain string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m Mail) GenerateNonAuthToken(UserEmail string, secret string) (*string, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m Mail) DecodeToken(token, secret string) (string, error) {
//	//TODO implement me
//	panic("implement me")
//}

func NewMail() ports.MailerRepository {
	return &Mail{}
}

// SendMail  METHOD THAT  WILL BE USED TO SEND EMAILS TO USERS
func (s *Mail) SendMail(subject, body, recipient, Private, Domain string) error {
	privateAPIKey := Private
	yourDomain := Domain

	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Create a new message with template
	m := mg.NewMessage("Oja Ecommerce <Oja@Decadev.gon>", subject, "")
	m.SetTemplate("i_template")

	// Add recipients
	err := m.AddRecipient(recipient)
	if err != nil {
		return err
	}

	// Add the variables recipient be used by the template
	err = m.AddVariable("title", subject)
	if err != nil {
		return err
	}
	err = m.AddVariable("body", body)
	if err != nil {
		return err
	}

	// Send the message with a 10 second timeout
	_, _, err = mg.Send(ctx, m)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
