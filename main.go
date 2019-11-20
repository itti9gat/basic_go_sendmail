package main

import (
	"log"
	"net/smtp"
)

type smtpServer struct {
	Host     string
	Port     string
	Username string
	Password string
}

type mailMessage struct {
	Sender   string
	Receiver []string
	Subject  string
	Body     string
}

func (s *smtpServer) Address() string {
	return s.Host + ":" + s.Port
}

func (s *smtpServer) SendMessage(m mailMessage) error {

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	for _, e := range m.Receiver {
		msg := []byte("To: " + e + "\r\n" +
			"Subject: " + m.Subject + "\r\n" +
			"\r\n" + m.Body + "\r\n")

		err := smtp.SendMail(s.Address(), auth, m.Sender, []string{e}, msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {

	smtpServer := smtpServer{
		Host:     "smtp.gmail.com",
		Port:     "587",
		Username: "your_email",
		Password: "your_password",
	}

	subject := "Email Subject Lines"
	body := "Do you want your email content opened, read, and clicked? It all starts with the \r\n" +
		"subject line. Read on for some tried-and-true tips to help jazz up your subject \r\n" +
        "lines and boost your emails' engagement."

	mailMessage := mailMessage{
		Sender:   "sender_email",
		Receiver: []string{"receiver_email_1", "receiver_email_2"},
		Subject:  subject,
		Body:     body,
	}

	err := smtpServer.SendMessage(mailMessage)
	if err != nil {
		log.Fatal(err)
	}

}
