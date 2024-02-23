package repository

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/xerrors"

	"emails/model/po"
	"emails/thirdparty/smtp"
)

const emailSender = "genning7@gmail.com"

func SendEmail(ctx context.Context, req po.Email) (err error) {
	password := os.Getenv("EMAIL_PASSWORD")
	fmt.Println("EMAIL_PASSWORD", password)
	if password == "" {
		err = xerrors.Errorf("EMAIL_PASSWORD is not set")
		return
	}

	if err = smtp.SendEmail(ctx, smtp.SendEmailReq{
		To:       []string{req.To},
		Subject:  req.Subject,
		Body:     req.Body,
		Password: password,
		From:     emailSender,
		SmtpHost: os.Getenv("EMAIL_HOST"),
		SmtpPort: os.Getenv("EMAIL_PORT"),
	}); err != nil {
		err = xerrors.Errorf("failed to send email: %w", err)
		return
	}

	return
}
