package utils

import "fmt"

var (
	Red     = outer(31)
	Green   = outer(32)
	Yellow  = outer(33)
	Blue    = outer(34)
	Magenta = outer(35)
	Cyan    = outer(36)
	White   = outer(37)
)

type (
	inner func(m interface{}) string
)

func outer(n int) inner {
	return func(m interface{}) string {
		return fmt.Sprintf("\033[%dm%v\033[0m", n, m)
	}
}
