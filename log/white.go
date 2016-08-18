// +build appengine

package log

import "io"

func output() io.Writer {
	return os.stdout
}
