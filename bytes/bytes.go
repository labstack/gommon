package bytes

import (
	"fmt"
	"math"
)

// Format formats bytes to string with decimal prefix. For example, 1000 would returns
// 1 KB
func Format(b uint64) string {
	return format(float64(b), false)
}

// FormatB formats bytes to string with binary prefix. For example, 1024 would
// return 1 KiB.
func FormatB(b uint64) string {
	return format(float64(b), true)
}

func format(b float64, bin bool) string {
	unit := float64(1000)
	if bin {
		unit = 1024
	}
	if b < unit {
		return fmt.Sprintf("%.0f B", b)
	} else {
		x := math.Floor(math.Log(b) / math.Log(unit))
		pre := make([]byte, 1, 2)
		pre[0] = "KMGTPE"[uint8(x)-1]
		if bin {
			pre = pre[:2]
			pre[1] = 'i'
		}
		return fmt.Sprintf("%.02f %sB", b/math.Pow(unit, x), pre)
	}
}
