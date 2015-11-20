package log

import "testing"

func TestLog(t *testing.T) {
	l := New("test")
	test(l, trace, t)
	test(global, trace, t)
	test(l, notice, t)
	test(global, notice, t)
}

func test(l *Logger, v Level, t *testing.T) {
	l.SetLevel(trace)
	l.Trace("trace")
	l.Debug("debug")
	l.Info("info")
	l.Notice("notice")
	l.Warn("warn")
	l.Error("error")
	l.Fatal("fatal")
}
