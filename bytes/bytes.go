package bytes

import (
	"bytes"
	"fmt"
	"math"
)

// Format formats bytes to string. For example, 1000 would returns 1 KB
func Format(b uint64) string {
	return format(float64(b), false)
}

// FormatBin formats bytes to string as specified by ICE standard. For example,
// 1024 would return 1 KiB.
func FormatBin(b uint64) string {
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
		exp := math.Floor(math.Log(b) / math.Log(unit))
		pfx := new(bytes.Buffer)
		pfx.WriteByte("KMGTPE"[uint8(exp)-1])
		if bin {
			pfx.WriteString("i")
		}
		return fmt.Sprintf("%.02f %sB", b/math.Pow(unit, exp), pfx)
	}
}
