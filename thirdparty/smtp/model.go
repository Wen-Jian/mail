package smtp

type SendEmailReq struct {
	To       []string
	Subject  string
	Body     string
	Password string
	From     string
	SmtpHost string
	SmtpPort string
}
