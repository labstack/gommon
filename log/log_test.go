package log

import "testing"

func TestLog(t *testing.T) {
	// l := New("test")
	// b := new(bytes.Buffer)
	// l.SetOutput(b)
	// test(l, TRACE, t)
	// assert.Contains(t, b.String(), "trace")
	// assert.Contains(t, b.String(), "fatal")
	//
	// b.Reset()
	// SetOutput(b)
	// test(global, NOTICE, t)
	// assert.NotContains(t, b.String(), "info")
	// assert.Contains(t, b.String(), "notice")
	// assert.Contains(t, b.String(), "fatal")

	test(global, 8, t)
}

func test(l *Logger, v Level, t *testing.T) {
	l.SetLevel(v)
	l.Trace("trace")
	l.Debug("debug")
	l.Info("info")
	l.Notice("notice")
	l.Warn("warn")
	l.Error("error")
	l.Fatal("fatal")
}
