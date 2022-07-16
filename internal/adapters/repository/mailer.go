package repository

import (
	"context"
	"github.com/decadevs/lunch-api/internal/ports"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

//type Claims struct {
//	UserEmail string `json:"email"`
//	jwt.StandardClaims
//}

//
//type Service struct{}

type Mail struct {
	//Client *mailgun.MailgunImpl
}

func NewMail() ports.MailerRepository {
	return &Mail{}
}

// SendMail  METHOD THAT  WILL BE USED TO SEND EMAILS TO USERS
func (s *Mail) SendMail(subject, body, recipient, Private, Domain string) error {
	//privateAPIKey := Private
	//yourDomain := Domain

	mg := mailgun.NewMailgun(Domain, Private)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	// Create a new message with template
	m := mg.NewMessage("Lunch Wallet <Lunch@Decadev.com>", subject, "")
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
