package simplesmtp

import (
	"fmt"
	"net/smtp"
	"strconv"
)

const (
	DEFAULT_SMTP_HOST = "smtp.gmail.com"
	DEFAULT_SMTP_PORT = 587
	msg               = "From: %s \n" +
		"To: %s \n" +
		"Subject: %s \n\n %s"
)

type (
	SimpleSmtp struct {
		Email    string      // your email address
		Password string      // your password of smtp account
		Host     string      // smtp host, default is "smtp.gmail.com"
		Port     int         // smtp port, default is 587
		To       []string    // email address to send
		Subject  string      // email subject
		Body     interface{} // email body
	}

	response struct {
		Status  bool
		Message string
	}
)

func (s *SimpleSmtp) Send() {
	defer func() {
		if err := recover(); err != nil {
			errMsg := response{Status: false, Message: err.(error).Error()}
			fmt.Println(errMsg)
		}
	}()

	if s.Host == "" {
		s.Host = DEFAULT_SMTP_HOST
	}
	if s.Port == 0 {
		s.Port = DEFAULT_SMTP_PORT
	}

	for _, email := range s.To {
		if err := smtp.SendMail(s.Host+":"+strconv.Itoa(s.Port), smtp.PlainAuth("", s.Email, s.Password, s.Host), s.Email, []string{email}, []byte(fmt.Sprintf(msg, s.Email, email, s.Subject, s.Body))); err != nil {
			panic(err)
		}
		fmt.Println(response{Status: true, Message: "success send email to " + email})
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
