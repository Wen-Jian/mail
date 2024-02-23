package bo

type EmailReq struct {
	To      string
	Subject string
	Body    string
}

type EmailTemplate struct {
	Url     string
	Body    string
	To      string
	Subject string
}
