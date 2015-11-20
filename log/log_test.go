package log

import "testing"

func TestLog(t *testing.T) {
	l := New("log")
	l.SetLevel(Warn)
	l.Trace("trace")
	l.Debug("debug")
	l.Info("info")
	l.Notice("notice")
	l.Warn("warn")
	l.Error("error")
	l.Fatal("fatal")
}
