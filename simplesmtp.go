package simplesmtp

import (
	"fmt"
	"net/smtp"
	"strconv"
)

const (
	DEFAULT_SMTP_HOST = "smtp.gmail.com"
	DEFAULT_SMTP_PORT = 587
)

type ISimpleSmtp interface {
	Send()
}

type SimpleSmtp struct {
	Email    string      // email address
	Password string      // password of smtp account
	Host     string      // smtp host default is "smtp.gmail.com"
	Port     int         // smtp port default is 587
	To       []string    // email address to send
	Subject  string      // email subject
	Body     interface{} // email body
}

func (s *SimpleSmtp) Send() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if s.Host == "" {
		s.Host = DEFAULT_SMTP_HOST
	}
	if s.Port == 0 {
		s.Port = DEFAULT_SMTP_PORT
	}

	for _, email := range s.To {
		if err := smtp.SendMail(s.Host+":"+strconv.Itoa(s.Port), smtp.PlainAuth("", s.Email, s.Password, s.Host), s.Email, []string{email}, []byte(s.Body.(string))); err != nil {
			panic(err)
		}
	}
}

func NewSimpleSmtp(emailFrom string, password string, host string, port int, to []string, subject string, body interface{}) *SimpleSmtp {
	return &SimpleSmtp{
		Email:    emailFrom,
		Password: password,
		Host:     host,
		Port:     port,
		To:       to,
		Subject:  subject,
		Body:     body,
	}
}
