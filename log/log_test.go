package log

import (
	"bytes"
	"testing"

	"os"
	"os/exec"

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

	b.Reset()
	SetOutput(b)
	test(global, WARN, t)
	assert.NotContains(t, b.String(), "info")
	assert.Contains(t, b.String(), "warn")
	println(b.String())
}

func TestFatal(t *testing.T) {
	l := New("test")
	switch os.Getenv("TEST_LOGGER_FATAL") {
	case "fatal":
		l.Fatal("fatal")
		return
	case "fatalf":
		l.Fatalf("fatal-%s", "f")
		return
	}

	loggerFatalTest(t, "fatal", "fatal")
	loggerFatalTest(t, "fatalf", "fatal-f")
}

func loggerFatalTest(t *testing.T, env string, contains string) {
	buf := new(bytes.Buffer)
	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	cmd.Env = append(os.Environ(), "TEST_LOGGER_FATAL="+env)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		assert.Contains(t, buf.String(), contains)
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
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
}
