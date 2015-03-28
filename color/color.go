package color

import (
	"bytes"
	"fmt"
)

type (
	inner func(interface{}, ...string) string
)

// Color styles
const (
	// Blk Black text style
	Blk = "30"
	// Rd red text style
	Rd = "31"
	// Grn green text style
	Grn = "32"
	// Yel yellow text style
	Yel = "33"
	// Blu blue text style
	Blu = "34"
	// Mgn magenta text style
	Mgn = "35"
	// Cyn cyan text style
	Cyn = "36"
	// Wht white text style
	Wht = "37"
	// Gry grey text style
	Gry = "90"

	// BlkBg black background style
	BlkBg = "40"
	// RdBg red background style
	RdBg = "41"
	// GrnBg green background style
	GrnBg = "42"
	// YelBg yellow background style
	YelBg = "43"
	// BluBg blue background style
	BluBg = "44"
	// MgnBg magenta background style
	MgnBg = "45"
	// CynBg cyan background style
	CynBg = "46"
	// WhtBg white background style
	WhtBg = "47"

	// R reset emphasis style
	R = "0"
	// B bold emphasis style
	B = "1"
	// D dim emphasis style
	D = "2"
	// I italic emphasis style
	I = "3"
	// U underline emphasis style
	U = "4"
	// In inverse emphasis style
	In = "7"
	// Hd hidden emphasis style
	Hd = "8"
	// So strikeout emphasis style
	So = "9"
)

// Color functions
var (
	// Text color
	Black   = outer(Blk)
	Red     = outer(Rd)
	Green   = outer(Grn)
	Yellow  = outer(Yel)
	Blue    = outer(Blu)
	Magenta = outer(Mgn)
	Cyan    = outer(Cyn)
	White   = outer(Wht)
	Grey    = outer(Gry)

	// Background color
	BlackBg   = outer(BlkBg)
	RedBg     = outer(RdBg)
	GreenBg   = outer(GrnBg)
	YellowBg  = outer(YelBg)
	BlueBg    = outer(BluBg)
	MagentaBg = outer(MgnBg)
	CyanBg    = outer(CynBg)
	WhiteBg   = outer(WhtBg)

	// Emphasis
	Reset     = outer(R)
	Bold      = outer(B)
	Dim       = outer(D)
	Italic    = outer(I)
	Underline = outer(U)
	Inverse   = outer(In)
	Hidden    = outer(Hd)
	Strikeout = outer(So)
)

func outer(n string) inner {
	return func(m interface{}, style ...string) string {
		b := new(bytes.Buffer)
		b.WriteString("\x1b[")
		b.WriteString(n)
		for _, s := range style {
			b.WriteString(";")
			b.WriteString(s)
		}
		b.WriteString("m")
		// TODO: Replace fmt for performance
		return fmt.Sprintf("%s%v\x1b[0m", b.String(), m)
	}
}
