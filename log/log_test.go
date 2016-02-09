package log

import (
	"bytes"
	"os"
	"os/exec"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestLog(t *testing.T) {
	l := New("test")
	b := new(bytes.Buffer)
	l.SetOutput(b)
	test(l, DEBUG, t)
	assert.Contains(t, b.String(), "\nDEBUG|test|debug\n")
	assert.Contains(t, b.String(), "\nDEBUG|test|debugf\n")
	assert.Contains(t, b.String(), "\nWARN|test|warn\n")
	assert.Contains(t, b.String(), "\nWARN|test|warnf\n")

	b.Reset()
	SetOutput(b)
	test(global, WARN, t)
	assert.NotContains(t, b.String(), "info")
	assert.Contains(t, b.String(), "\nWARN|-|warn\n")
	println(b.String())
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
