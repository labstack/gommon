package random

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

type (
	Random struct {
		lock sync.Mutex
		src  rand.Source
	}
)

func (r *Random) Int63() (n int64) {
	r.lock.Lock()
	n = r.src.Int63()
	r.lock.Unlock()
	return
}

func (r *Random) Seed(seed int64) {
	r.lock.Lock()
	r.src.Seed(seed)
	r.lock.Unlock()
}

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
	src := rand.NewSource(time.Now().UnixNano())
	return &Random{src: src}
}

func (r *Random) String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Int63()%int64(len(charset))]
	}
	return string(b)
}

func String(length uint8, charsets ...string) string {
	return global.String(length, charsets...)
}
