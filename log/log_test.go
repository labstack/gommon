package log

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	l := New("test")
	b := new(bytes.Buffer)
	l.SetOutput(b)
	test(l, DEBUG, t)
	assert.Contains(t, b.String(), "debug")
	assert.Contains(t, b.String(), "debugf")
	assert.Contains(t, b.String(), "warn")
	assert.Contains(t, b.String(), "warnf")
	// assert.Contains(t, b.String(), "fatal")

	b.Reset()
	SetOutput(b)
	test(global, WARN, t)
	assert.NotContains(t, b.String(), "info")
	assert.Contains(t, b.String(), "warn")
	// assert.Contains(t, b.String(), "fatal")
}

func test(l *Logger, v Level, t *testing.T) {
	l.SetLevel(v)
	l.Print("print")
	l.Printf("print%s", "f")
	l.Debug("debug")
	l.Debugf("debug%s", "f")
	l.Info("info")
	l.Infof("info%s", "f")
	l.Warn("warn")
	l.Warnf("warn%s", "f")
	l.Error("error")
	l.Errorf("error%s", "f")
	// l.Fatal("fatal")
}
