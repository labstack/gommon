package log

import (
	"bytes"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	l := New("test")
	b := new(bytes.Buffer)
	l.SetOutput(b)
	test(l, TRACE, t)
	assert.Contains(t, b.String(), "\nTRACE|test|trace\n")
	assert.Contains(t, b.String(), "\nFATAL|test|fatal\n")

	b.Reset()
	SetOutput(b)
	test(global, NOTICE, t)
	assert.NotContains(t, b.String(), "\nINFO|-|info\n")
	assert.Contains(t, b.String(), "\nNOTICE|-|notice\n")
	assert.Contains(t, b.String(), "\nFATAL|-|fatal\n")
}

func TestLogConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			TestLog(t)
			wg.Done()
		}()
	}
	wg.Wait()
}

func test(l *Logger, v Level, t *testing.T) {
	l.SetLevel(v)
	l.Print("print")
	l.Println("println")
	l.Trace("trace")
	l.Debug("debug")
	l.Info("info")
	l.Notice("notice")
	l.Warn("warn")
	l.Error("error")
	l.Fatal("fatal")
}
