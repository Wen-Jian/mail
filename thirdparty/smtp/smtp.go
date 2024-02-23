package smtp

import (
	"context"
	"fmt"
	"net/smtp"
)

func SendEmail(ctx context.Context, req SendEmailReq) (err error) {
	auth := smtp.PlainAuth("", req.From, req.Password, req.SmtpHost)
	err = smtp.SendMail(req.SmtpHost+":"+req.SmtpPort, auth, req.From, req.To, []byte(req.Body))
	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}
