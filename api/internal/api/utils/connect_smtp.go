package utils

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func ConnectSMTP(host string, port int, user string, password string) *gomail.Dialer {
	dialer := gomail.NewDialer(host, port, user, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return dialer
}
