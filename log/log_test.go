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
	b := new(bytes.Buffer)
	l.SetOutput(b)
	l.DisableColor()
	l.SetLevel(WARN)

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

	assert.Contains(t, b.String(), "print\n")
	assert.Contains(t, b.String(), "\nprintf\n")
	assert.NotContains(t, b.String(), "debug")
	assert.NotContains(t, b.String(), "debugf")
	assert.NotContains(t, b.String(), "info")
	assert.NotContains(t, b.String(), "infof")
	assert.Contains(t, b.String(), "\nWARN|"+l.prefix+"|warn\n")
	assert.Contains(t, b.String(), "\nWARN|"+l.prefix+"|warnf\n")
	assert.Contains(t, b.String(), "\nERROR|"+l.prefix+"|error\n")
	assert.Contains(t, b.String(), "\nERROR|"+l.prefix+"|errorf\n")
}

func TestLog(t *testing.T) {
	l := New("test")
	test(l, t)
}

func TestGlobal(t *testing.T) {
	test(global, t)
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
