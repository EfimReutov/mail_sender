# Mail Sender
![https://img.shields.io/github/v/tag/EfimReutov/mail_sender](https://img.shields.io/github/v/tag/EfimReutov/mail_sender)
![https://img.shields.io/github/license/EfimReutov/mail_sender](https://img.shields.io/github/license/EfimReutov/mail_sender)

This library is for sending emails from your mail

## Installation

mail_sender can be installed like any other Go library through go get:

```console
$ go get github.com/EfimReutov/mail_sender
```

Or, if you are already using
[Go Modules](https://github.com/golang/go/wiki/Modules), you may specify a version number as well:

```console
$ go get github.com/EfimReutov/mail_sender@latest
```

## Getting Started
```go
package main

import "github.com/EfimReutov/mail_sender"

func main() {
	sender := mail_sender.NewSender(
		mail_sender.Configuration{
			SMTPServer:   "smtp.gmail.com",
			SMTPPort:     587,
			MailUser:     "from@example.com",
			MailPassword: "examplePassword",
		},
	)
	dest := []string{
		"to@example.com",
	}

	data := struct {
		ExampleVariable string
	}{
		ExampleVariable: "example",
	}
	err := sender.WriteHTMLEmail(dest, "your subject", "templates/example.html", data)
	if err != nil {
		panic(err)
	}
	err = sender.WritePlainEmail(dest, "your subject", "text message")
	if err != nil {
		panic(err)
	}
}
```