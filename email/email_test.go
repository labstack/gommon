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

// TestSend_CCRecipients verifies CC addresses receive RCPT commands, not
// just the CC header: SMTP only delivers to RCPT recipients.
func TestSend_CCRecipients(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)
	defer ln.Close()

	rcpts := make(chan string, 8)
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
			case strings.HasPrefix(upper, "MAIL"):
				_, _ = conn.Write([]byte("250 OK\r\n"))
			case strings.HasPrefix(upper, "RCPT"):
				rcpts <- line
				_, _ = conn.Write([]byte("250 OK\r\n"))
			case strings.HasPrefix(upper, "DATA"):
				_, _ = conn.Write([]byte("354 go\r\n"))
			case upper == ".":
				_, _ = conn.Write([]byte("250 queued\r\n"))
			case strings.HasPrefix(upper, "QUIT"):
				_, _ = conn.Write([]byte("221 bye\r\n"))
				return
			default:
				_, _ = conn.Write([]byte("250 OK\r\n"))
			}
		}
	}()

	e := New(ln.Addr().String())
	e.DialTimeout = 2 * time.Second
	m := &Message{
		ID:       "<test@fake.local>",
		From:     "sender@fake.local",
		To:       "to@fake.local",
		CC:       "cc1@fake.local, cc2@fake.local",
		Subject:  "hi",
		BodyText: "body",
	}
	err = e.Send(m)
	assert.NoError(t, err)

	got := []string{}
	for len(rcpts) > 0 {
		got = append(got, <-rcpts)
	}
	assert.Len(t, got, 3)
	assert.Contains(t, got[0], "to@fake.local")
	assert.Contains(t, got[1], "cc1@fake.local")
	assert.Contains(t, got[2], "cc2@fake.local")
	<-done
}
