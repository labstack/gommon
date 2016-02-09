package log

import (
	"bytes"
	"os"
	"os/exec"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func test(l *Logger, t *testing.T) {
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

func TestLog(t *testing.T) {
	l := New("test")
	b := new(bytes.Buffer)
	l.SetOutput(b)
	// l.DisableColor()
	l.SetLevel(INFO)
	test(l, t)
	// assert.Contains(t, b.String(), "print\n")
	assert.Contains(t, b.String(), "printf\n")
	assert.NotContains(t, b.String(), "debug")
	assert.NotContains(t, b.String(), "debugf")
	assert.Contains(t, b.String(), "\nWARN|test|warn\n")
	assert.Contains(t, b.String(), "\nWARN|test|warnf\n")
}

func TestGlobal(t *testing.T) {
	b := new(bytes.Buffer)
	SetOutput(b)
	// DisableColor()
	SetLevel(INFO)
	test(global, t)
	// assert.Contains(t, b.String(), "print\n")
	assert.Contains(t, b.String(), "printf\n")
	assert.NotContains(t, b.String(), "debug")
	assert.NotContains(t, b.String(), "debugf")
	assert.Contains(t, b.String(), "\nWARN|-|warn\n")
	assert.Contains(t, b.String(), "\nWARN|-|warnf\n")
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
