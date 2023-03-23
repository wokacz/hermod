package smtp

import (
	"github.com/wokacz/hermod/pkg/env"
	"gopkg.in/mail.v2"
)

var Dialer *mail.Dialer

type Config struct {
	UserName string
	Password string
	Host     string
	Port     int
}

// Init initializes the SMTP dialer with the given configuration.
func Init() {
	userName := env.Get("SMTP_USERNAME")
	password := env.Get("SMTP_PASSWORD")
	host := env.Get("SMTP_HOST")
	port := env.GetInt("SMTP_PORT")
	Dialer = mail.NewDialer(host, port, userName, password)
}
