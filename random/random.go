package random

import (
	"bufio"
	"crypto/rand"
	"strings"
)

type (
	Random struct {
	}
)

// Charsets
const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic   = Uppercase + Lowercase
	Numeric      = "0123456789"
	Alphanumeric = Alphabetic + Numeric
	Symbols      = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex          = Numeric + "abcdef"
)

var (
	global = New()
)

func New() *Random {
	return new(Random)
}

func (r *Random) String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}
	reader := bufio.NewReaderSize(rand.Reader, int(length))
	buf := make([]byte, length)
	for i := range buf {
		b, err := reader.ReadByte()
		if err != nil {
			panic(err)
		}
		buf[i] = charset[int(b)%len(charset)]
	}
	return string(buf)
}

func String(length uint8, charsets ...string) string {
	return global.String(length, charsets...)
}
