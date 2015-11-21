package log

import "testing"

func TestLog(t *testing.T) {
	l := New("test")
	test(l, TRACE, t)
	test(global, TRACE, t)
	test(l, NOTICE, t)
	test(global, NOTICE, t)
}

func test(l *Logger, v level, t *testing.T) {
	l.SetLevel(TRACE)
	l.Trace("trace")
	l.Debug("debug")
	l.Info("info")
	l.Notice("notice")
	l.Warn("warn")
	l.Error("error")
	l.Fatal("fatal")
}
