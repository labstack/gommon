package gytes

import (
	"fmt"
	"math"
	"strconv"
)

var (
	global = New()
)

type (
	Gytes struct {
		iec bool
	}
)

// New creates a Gytes instance.
func New() *Gytes {
	return &Gytes{}
}

// SetBinaryPrefix sets binary prefix format.
func (g *Gytes) SetBinaryPrefix(on bool) {
	g.iec = on
}

// Format formats bytes to string. For example, 1323 bytes will return 1.32 KB.
// If binary prefix is set, it will return 1.29 KiB.
func (g *Gytes) Format(b uint64) string {
	unit := uint64(1000)
	if g.iec {
		unit = 1024
	}
	if b < unit {
		return strconv.FormatUint(b, 10) + " B"
	} else {
		b := float64(b)
		unit := float64(unit)
		x := math.Floor(math.Log(b) / math.Log(unit))
		pre := make([]byte, 1, 2)
		pre[0] = "KMGTPE"[uint8(x)-1]
		if g.iec {
			pre = pre[:2]
			pre[1] = 'i'
		}
		// TODO: Improve performance?
		return fmt.Sprintf("%.02f %sB", b/math.Pow(unit, x), pre)
	}
}

func BinaryPrefix(on bool) {
	global.SetBinaryPrefix(on)
}

// Format wraps default instance's Format function.
func Format(b uint64) string {
	return global.Format(b)
}
