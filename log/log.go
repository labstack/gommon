package log

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type (
	Log struct {
		level  Level
		out    io.Writer
		err    io.Writer
		prefix string
		sync.Mutex
		lock bool
	}
	Level uint8
)

const (
	Trace = iota
	Debug
	Info
	Notice
	Warn
	Error
	Fatal
	Off = 10
)

var (
	levels = []string{
		"trace",
		"debug",
		"info",
		"notice",
		"warn",
		"error",
		"fatal",
	}
)

func New(prefix string) *Log {
	return &Log{
		level:  Debug,
		out:    os.Stdout,
		err:    os.Stderr,
		prefix: prefix,
	}
}

func (l *Log) SetLevel(v Level) {
	l.level = v
}

func (l *Log) SetOutput(w io.Writer) {
	l.out = w
	l.err = w

	switch w.(type) {
	case *os.File:
		l.lock = true
	default:
		l.lock = false
	}
}

func (l *Log) Trace(i interface{}) {
	l.log(Trace, l.out, i)
}

func (l *Log) Debug(i interface{}) {
	l.log(Debug, l.out, i)
}

func (l *Log) Info(i interface{}) {
	l.log(Info, l.out, i)
}

func (l *Log) Notice(i interface{}) {
	l.log(Notice, l.out, i)
}

func (l *Log) Warn(i interface{}) {
	l.log(Warn, l.out, i)
}

func (l *Log) Error(i interface{}) {
	l.log(Error, l.err, i)
}

func (l *Log) Fatal(i interface{}) {
	l.log(Fatal, l.err, i)
}

func (l *Log) log(v Level, w io.Writer, i interface{}) {
	if l.lock {
		l.Lock()
		defer l.Unlock()
	}
	if v >= l.level {
		fmt.Fprintf(w, "%s|%v\n", levels[v], i)
	}
}
