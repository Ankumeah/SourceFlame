package email

import (
	"context"
	"net/smtp"
	"strings"
)

type NetSMTPEmailManager struct {
	addr string
	auth smtp.Auth
	user string
}

func GetNetSMTPEmailManager(user string, pass string, url string) EmailManager {
	return &NetSMTPEmailManager{
		addr: url,
		auth: smtp.PlainAuth("", user, pass, strings.Split(url, ":")[0]),
		user: user,
	}
}

func (e *NetSMTPEmailManager) SendEmail(ctx context.Context, email string, data []byte) error {
	return smtp.SendMail(e.addr, e.auth, e.user, []string{email}, data)
}
