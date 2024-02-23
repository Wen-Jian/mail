package usecase

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"golang.org/x/xerrors"

	"emails/model/bo"
	"emails/model/po"
	"emails/repository"
)

func SendEmails(ctx context.Context, req bo.EmailReq) (err error) {
	// build template
	// t, err := template.ParseFiles("template.html")
	// if err != nil {
	// 	return
	// }

	// var tpl bytes.Buffer
	// if err = t.Execute(&tpl, bo.EmailTemplate{
	// 	Url:     os.Getenv("APP_URL"),
	// 	To:      req.To,
	// 	Body:    req.Body,
	// 	Subject: req.Subject,
	// }); err != nil {
	// 	return
	// }

	var body bytes.Buffer

	msg := fmt.Sprintf(`
		<html>
		<body>
		<div>POST <span>%s/emails</span></div>
		<div>with JSON body:</div>
		<div>
			{
			<p>to: %s</p>
			<p>subject: %s</p>
			<p>body: %s</p>
			}
		</div>
		</body>
		</html>
		`, os.Getenv("APP_URL"), req.To, req.Subject, req.Body)

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n%s", req.Subject, mimeHeaders, msg)))

	if err = repository.SendEmail(ctx, po.Email{
		To:   req.To,
		Body: body.String(),
	}); err != nil {
		err = xerrors.Errorf("failed to send email: %w", err)
		return
	}

	return
}
