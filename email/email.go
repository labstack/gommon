package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"

	"net"

	"time"

	"github.com/labstack/gommon/random"
)

type (
	Email struct {
		Auth        smtp.Auth
		Header      map[string]string
		Template    *template.Template
		smtpAddress string
	}

	Message struct {
		ID          string
		From        string
		To          string
		CC          string
		Subject     string
		Text        string
		HTML        string
		Inlines     []*File
		Attachments []*File
		buffer      *bytes.Buffer
		boundary    string
	}

	File struct {
		Name    string
		Type    string
		Content string
	}
)

func New(smtpAddress string) *Email {
	return &Email{
		smtpAddress: smtpAddress,
		Header:      map[string]string{},
	}
}

func (m *Message) writeText(content string, contentType string) {
	m.buffer.WriteString(fmt.Sprintf("--%s\n", m.boundary))
	m.buffer.WriteString(fmt.Sprintf("Content-Type: %s; charset=UTF-8\n", contentType))
	m.buffer.WriteString("Content-Transfer-Encoding: quoted-printable\n")
	m.buffer.WriteString("\n")
	m.buffer.WriteString(content + "\n")
}

func (m *Message) writeFile(f *File, disposition string) {
	m.buffer.WriteString(fmt.Sprintf("--%s\n", m.boundary))
	m.buffer.WriteString(fmt.Sprintf("Content-Type: %s; name=%s\n", f.Type, f.Name))
	m.buffer.WriteString(fmt.Sprintf("Content-Disposition: %s; filename=%s\n", disposition, f.Name))
	m.buffer.WriteString("Content-Transfer-Encoding: base64\n")
	m.buffer.WriteString("\n")
	m.buffer.WriteString(f.Content + "\n")
}

func (e *Email) Send(m *Message) (err error) {
	// Message header
	m.buffer = new(bytes.Buffer)
	m.boundary = random.String(16)
	m.buffer.WriteString("MIME-Version: 1.0\n")
	m.buffer.WriteString(fmt.Sprintf("Message-ID: %s\n", m.ID))
	m.buffer.WriteString(fmt.Sprintf("Date: %s\n", time.Now().Format(time.RFC1123Z)))
	m.buffer.WriteString(fmt.Sprintf("From: %s\n", m.From))
	m.buffer.WriteString(fmt.Sprintf("To: %s\n", m.To))
	if m.CC != "" {
		m.buffer.WriteString(fmt.Sprintf("CC: %s\n", m.CC))
	}
	if m.Subject != "" {
		m.buffer.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	}
	// Extra
	for k, v := range e.Header {
		m.buffer.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
	m.buffer.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", m.boundary))
	m.buffer.WriteString("\n")

	// Message body
	if m.Text != "" {
		m.writeText(m.Text, "text/plain")
	} else if m.HTML != "" {
		m.writeText(m.HTML, "text/html")
	} else {
		// TODO:
	}

	// Attachments / inlines
	for _, f := range m.Inlines {
		m.writeFile(f, "inline")
	}
	for _, f := range m.Attachments {
		m.writeFile(f, "disposition")
	}
	m.buffer.WriteString("\n")
	m.buffer.WriteString("--" + m.boundary + "--")

	// Dial
	c, err := smtp.Dial(e.smtpAddress)
	if err != nil {
		return
	}
	defer c.Close()

	// Check if TLS is required
	if ok, _ := c.Extension("STARTTLS"); ok {
		host, _, _ := net.SplitHostPort(e.smtpAddress)
		config := &tls.Config{ServerName: host}
		if err = c.StartTLS(config); err != nil {
			return err
		}
	}

	// Authenticate
	if e.Auth != nil {
		if err = c.Auth(e.Auth); err != nil {
			return
		}
	}

	// Send message
	from, err := mail.ParseAddress(m.From)
	if err != nil {
		return
	}
	if err = c.Mail(from.Address); err != nil {
		return
	}
	to, err := mail.ParseAddressList(m.To)
	if err != nil {
		return
	}
	for _, a := range to {
		if err = c.Rcpt(a.Address); err != nil {
			return
		}
	}
	wc, err := c.Data()
	if err != nil {
		return
	}
	defer wc.Close()
	_, err = m.buffer.WriteTo(wc)
	return
}
