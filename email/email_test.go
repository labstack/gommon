package email

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestDial_PlainAccepts verifies non-465 ports dial plaintext and
// complete the SMTP handshake against a minimal fake server.
func TestDial_PlainAccepts(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)
	defer ln.Close()

	done := runFakeSMTP(ln)

	e := New(ln.Addr().String())
	e.DialTimeout = 2 * time.Second

	c, err := e.dial()
	assert.NoError(t, err)
	if c != nil {
		// Our fake server never advertises STARTTLS, so dial must have
		// stayed plaintext — document that intent.
		ok, _ := c.Extension("STARTTLS")
		assert.False(t, ok, "fake server should not advertise STARTTLS")
		_ = c.Quit()
	}
	<-done
}

// TestDial_Timeout verifies DialTimeout caps dial wait on unreachable
// hosts. Uses 203.0.113.1 (TEST-NET-3, RFC 5737) which is unroutable.
// Timeout is 500ms to absorb CI scheduler jitter; the upper bound is
// still well under a round-trip expectation.
func TestDial_Timeout(t *testing.T) {
	e := New("203.0.113.1:465")
	e.DialTimeout = 500 * time.Millisecond

	start := time.Now()
	_, err := e.dial()
	elapsed := time.Since(start)

	assert.Error(t, err)
	assert.Less(t, elapsed, 3*time.Second)
}

// TestDial_InvalidAddress verifies malformed addresses fail fast.
func TestDial_InvalidAddress(t *testing.T) {
	e := New("not-a-valid-address")
	_, err := e.dial()
	assert.Error(t, err)
}

// runFakeSMTP accepts one connection and drives a minimal SMTP dialog.
// Returns a channel that closes when the connection is done.
func runFakeSMTP(ln net.Listener) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		conn.SetDeadline(time.Now().Add(5 * time.Second))

		_, _ = conn.Write([]byte("220 fake.local ESMTP ready\r\n"))
		r := bufio.NewReader(conn)
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				return
			}
			upper := strings.ToUpper(strings.TrimSpace(line))
			switch {
			case strings.HasPrefix(upper, "EHLO"), strings.HasPrefix(upper, "HELO"):
				_, _ = conn.Write([]byte("250-fake.local\r\n250 OK\r\n"))
			case strings.HasPrefix(upper, "QUIT"):
				_, _ = conn.Write([]byte("221 bye\r\n"))
				return
			default:
				_, _ = conn.Write([]byte("502 not implemented\r\n"))
			}
		}
	}()
	return done
}

// TestSend_MessageFormatting verifies MIME formatting for text-only, HTML-only, and multipart/alternative emails.
func TestSend_MessageFormatting(t *testing.T) {
	// We test message buffer construction by running a fake SMTP server that captures the DATA command.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)
	defer ln.Close()

	dataChan := make(chan string, 1)
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		conn.SetDeadline(time.Now().Add(5 * time.Second))

		_, _ = conn.Write([]byte("220 fake.local ESMTP ready\r\n"))
		r := bufio.NewReader(conn)
		recording := false
		var sb strings.Builder
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				return
			}
			upper := strings.ToUpper(strings.TrimSpace(line))
			switch {
			case strings.HasPrefix(upper, "EHLO"), strings.HasPrefix(upper, "HELO"):
				_, _ = conn.Write([]byte("250-fake.local\r\n250 OK\r\n"))
			case strings.HasPrefix(upper, "MAIL FROM"):
				_, _ = conn.Write([]byte("250 OK\r\n"))
			case strings.HasPrefix(upper, "RCPT TO"):
				_, _ = conn.Write([]byte("250 OK\r\n"))
			case strings.HasPrefix(upper, "DATA"):
				_, _ = conn.Write([]byte("354 Start mail input; end with <CRLF>.<CRLF>\r\n"))
				recording = true
			case recording && strings.TrimSpace(line) == ".":
				_, _ = conn.Write([]byte("250 OK message accepted\r\n"))
				dataChan <- sb.String()
				recording = false
			case recording:
				sb.WriteString(line)
			case strings.HasPrefix(upper, "QUIT"):
				_, _ = conn.Write([]byte("221 bye\r\n"))
				return
			default:
				_, _ = conn.Write([]byte("502 not implemented\r\n"))
			}
		}
	}()

	t.Run("Text only", func(t *testing.T) {
		e := New(ln.Addr().String())
		msg := &Message{
			From:     "sender@example.com",
			To:       "recipient@example.com",
			Subject:  "Test Text",
			BodyText: "Hello Text",
		}
		err := e.Send(msg)
		assert.NoError(t, err)

		select {
		case data := <-dataChan:
			assert.Contains(t, data, "Content-Type: text/plain; charset=UTF-8")
			assert.Contains(t, data, "Hello Text")
			assert.NotContains(t, data, "multipart/alternative")
		case <-time.After(2 * time.Second):
			t.Fatal("timeout waiting for SMTP data")
		}
	})

	t.Run("HTML only", func(t *testing.T) {
		e := New(ln.Addr().String())
		msg := &Message{
			From:     "sender@example.com",
			To:       "recipient@example.com",
			Subject:  "Test HTML",
			BodyHTML: "<p>Hello HTML</p>",
		}
		err := e.Send(msg)
		assert.NoError(t, err)

		select {
		case data := <-dataChan:
			assert.Contains(t, data, "Content-Type: text/html; charset=UTF-8")
			assert.Contains(t, data, "<p>Hello HTML</p>")
			assert.NotContains(t, data, "multipart/alternative")
		case <-time.After(2 * time.Second):
			t.Fatal("timeout waiting for SMTP data")
		}
	})

	t.Run("Both Text and HTML (Multipart Alternative)", func(t *testing.T) {
		e := New(ln.Addr().String())
		msg := &Message{
			From:     "sender@example.com",
			To:       "recipient@example.com",
			Subject:  "Test Both",
			BodyText: "Hello Text",
			BodyHTML: "<p>Hello HTML</p>",
		}
		err := e.Send(msg)
		assert.NoError(t, err)

		select {
		case data := <-dataChan:
			assert.Contains(t, data, "Content-Type: multipart/alternative;")
			assert.Contains(t, data, "text/plain; charset=UTF-8")
			assert.Contains(t, data, "Hello Text")
			assert.Contains(t, data, "text/html; charset=UTF-8")
			assert.Contains(t, data, "<p>Hello HTML</p>")
		case <-time.After(2 * time.Second):
			t.Fatal("timeout waiting for SMTP data")
		}
	})
}
