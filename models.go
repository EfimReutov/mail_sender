package mail_sender

const (
	mimeVersion      = "1.0"
	contentTypePlain = "text/plain"
	contentTypeHTML  = "text/html"
)

type Configuration struct {
	SMTPServer   string
	SMTPPort     int
	MailUser     string
	MailPassword string
}

type Sender struct {
	smtpServer string
	smtpPort   int
	user       string
	password   string
}

type Message struct {
	ContentType string
	Recipients  []string
	Subject     string
	Body        []byte
}
