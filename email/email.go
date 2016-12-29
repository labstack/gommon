package email

import (
	"bytes"
	"fmt"
	"net/mail"
	"net/smtp"

	"github.com/labstack/gommon/random"
)

type (
	Email struct {
		Auth        smtp.Auth
		smtpAddress string
		// TextTemplates text.Template
	}

	Message struct {
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
		Name    string `json:"name"`
		Type    string `json:"type"`
		Content string `json:"content"`
	}
)

func New(smtpAddress string) *Email {
	return &Email{
		smtpAddress: smtpAddress,
	}
}

func (m *Message) writeText(content string, contentType string) {
	m.buffer.WriteString(fmt.Sprintf("--%s\r\n", m.boundary))
	m.buffer.WriteString(fmt.Sprintf("Content-Type: %s; charset=UTF-8\r\n", contentType))
	m.buffer.WriteString("Content-Transfer-Encoding: quoted-printable\r\n")
	m.buffer.WriteString("\r\n")
	m.buffer.WriteString(content + "\r\n")
}

func (m *Message) writeFile(f *File, disposition string) {
	m.buffer.WriteString(fmt.Sprintf("--%s\r\n", m.boundary))
	m.buffer.WriteString(fmt.Sprintf("Content-Type: %s; name=%s\r\n", f.Type, f.Name))
	m.buffer.WriteString(fmt.Sprintf("Content-Disposition: %s; filename=%s\r\n", disposition, f.Name))
	m.buffer.WriteString("Content-Transfer-Encoding: base64\r\n")
	m.buffer.WriteString("\r\n")
	m.buffer.WriteString(f.Content + "\r\n")
}

func (e *Email) Send(m *Message) error {
	// Construct message
	m.buffer = new(bytes.Buffer)
	m.boundary = random.String(16)
	m.buffer.WriteString("MIME-Version: 1.0\r\n")
	m.buffer.WriteString(fmt.Sprintf("CC: %s\r\n", m.CC))
	m.buffer.WriteString(fmt.Sprintf("Subject: %s\r\n", m.Subject))
	m.buffer.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n", m.boundary))
	m.buffer.WriteString("\r\n")
	m.writeText(m.Text, "text/plain")
	m.writeText(m.HTML, "text/html")
	for _, f := range m.Inlines {
		m.writeFile(f, "inline")
	}
	for _, f := range m.Attachments {
		m.writeFile(f, "disposition")
	}
	m.buffer.WriteString("\r\n")
	m.buffer.WriteString("--" + m.boundary + "--")

	// Send message
	c, err := smtp.Dial(e.smtpAddress)
	if e.Auth != nil {
		// Authenticate
		if err := c.Auth(e.Auth); err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	defer c.Close()
	from, err := mail.ParseAddress(m.From)
	if err != nil {
		return err
	}
	if err = c.Mail(from.Address); err != nil {
		return err
	}
	to, err := mail.ParseAddressList(m.To)
	if err != nil {
		return err
	}
	for _, a := range to {
		if err = c.Rcpt(a.Address); err != nil {
			return err
		}
	}
	wc, err := c.Data()
	if err != nil {
		return err
	}
	defer wc.Close()
	_, err = m.buffer.WriteTo(wc)
	return err
}
