package util

import (
	"bytes"
	"net/smtp"
)

// you need run postfix on local server on default port 25

func Sendmail(src, dest, content string) {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("127.0.0.1:25")
	if err != nil {
		LogError().Println(err)
		panic(err)
	}
	defer c.Close()
	// Set the sender and recipient.
	c.Mail(src)
	c.Rcpt(dest)
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		LogError().Println(err)
		panic(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString(content)
	if _, err = buf.WriteTo(wc); err != nil {
		LogError().Println(err)
		panic(err)
	}
}
