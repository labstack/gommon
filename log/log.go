package log

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/labstack/gommon/color"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

type (
	Logger struct {
		level  Level
		out    io.Writer
		err    io.Writer
		prefix string
		sync.Mutex
	}
	Level uint8
)

const (
	TRACE = iota
	DEBUG
	INFO
	NOTICE
	WARN
	ERROR
	FATAL
	OFF
)

var (
	global = New("-")
	levels []string
)

func New(prefix string) (l *Logger) {
	l = &Logger{
		level:  INFO,
		prefix: prefix,
	}
	l.SetOutput(colorable.NewColorableStdout())
	l.err = colorable.NewColorableStderr()
	return
}

func (l *Logger) SetPrefix(p string) {
	l.prefix = p
}

func (l *Logger) SetLevel(v Level) {
	l.level = v
}

func (l *Logger) SetOutput(w io.Writer) {
	l.out = w
	l.err = w
	color.Disable()

	switch w := w.(type) {
	case *os.File:
		if isatty.IsTerminal(w.Fd()) {
			color.Enable()
		}
		levels = []string{
			color.Cyan("TRACE"),
			color.Blue("DEBUG"),
			color.Green("INFO"),
			color.Magenta("NOTICE"),
			color.Yellow("WARN"),
			color.Red("ERROR"),
			color.RedBg("FATAL"),
		}
	}
}

func (l *Logger) Trace(msg interface{}, args ...interface{}) {
	l.log(TRACE, l.out, msg, args...)
}

func (l *Logger) Debug(msg interface{}, args ...interface{}) {
	l.log(DEBUG, l.out, msg, args...)
}

func (l *Logger) Info(msg interface{}, args ...interface{}) {
	l.log(INFO, l.out, msg, args...)
}

func (l *Logger) Notice(msg interface{}, args ...interface{}) {
	l.log(NOTICE, l.out, msg, args...)
}

func (l *Logger) Warn(msg interface{}, args ...interface{}) {
	l.log(WARN, l.out, msg, args...)
}

func (l *Logger) Error(msg interface{}, args ...interface{}) {
	l.log(ERROR, l.err, msg, args...)
}

func (l *Logger) Fatal(msg interface{}, args ...interface{}) {
	l.log(FATAL, l.err, msg, args...)
}

func SetPrefix(p string) {
	global.SetPrefix(p)
}

func SetLevel(v Level) {
	global.SetLevel(v)
}

func SetOutput(w io.Writer) {
	global.SetOutput(w)
}

func Trace(msg interface{}, args ...interface{}) {
	global.Trace(msg, args...)
}

func Debug(msg interface{}, args ...interface{}) {
	global.Debug(msg, args...)
}

func Info(msg interface{}, args ...interface{}) {
	global.Info(msg, args...)
}

func Notice(msg interface{}, args ...interface{}) {
	global.Notice(msg, args...)
}

func Warn(msg interface{}, args ...interface{}) {
	global.Warn(msg, args...)
}

func Error(msg interface{}, args ...interface{}) {
	global.Error(msg, args...)
}

func Fatal(msg interface{}, args ...interface{}) {
	global.Fatal(msg, args...)
}

func (l *Logger) log(v Level, w io.Writer, msg interface{}, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if v >= l.level {
		// TODO: Improve performance
		f := fmt.Sprintf("%s|%s|%s\n", levels[v], l.prefix, msg)
		fmt.Fprintf(w, f, args...)
	}
}
