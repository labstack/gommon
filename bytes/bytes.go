package bytes

import (
	"fmt"
	"math"
)

var (
	sfx = "B"
)

// BinaryPrefix sets binary prefix as specified by ICE standard. For example, 13.23 MiB.
func BinaryPrefix(on bool) {
	if on {
		sfx = "iB"
	}
}

// Bytes formats bytes into string. For example, 1024 returns 1 KB
func Format(b uint64) string {
	n := float64(b)
	unit := float64(1024)

	if n == 0 {
		return "--"
	} else if n < unit {
		return fmt.Sprintf("%.0f B", n)
	} else {
		e := math.Floor(math.Log(n) / math.Log(unit))
		pfx := "KMGTPE"[uint8(e)-1]
		return fmt.Sprintf("%.02f %c%s", n/math.Pow(unit, e), pfx, sfx)
	}
}
