package color

import (
	"bytes"
	"fmt"
)

type (
	inner func(interface{}, []string) string
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

	// Rst reset emphasis style
	Rst = "0"
	// B bold emphasis style
	B = "1"
	// Dm dim emphasis style
	Dm = "2"
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
	global = New()
)

func outer(n string) inner {
	return func(msg interface{}, styles []string) string {
		b := new(bytes.Buffer)
		b.WriteString("\x1b[")
		b.WriteString(n)
		for _, s := range styles {
			b.WriteString(";")
			b.WriteString(s)
		}
		b.WriteString("m")
		// TODO: Replace fmt for performance
		return fmt.Sprintf("%s%v\x1b[0m", b.String(), msg)
	}
}

type (
	Color struct {
	}
)

func New() *Color {
	return &Color{}
}

func (c *Color) Black(msg interface{}, styles ...string) string {
	return outer(Blk)(msg, styles)
}

func (c *Color) Red(msg interface{}, styles ...string) string {
	return outer(Rd)(msg, styles)
}

func (c *Color) Green(msg interface{}, styles ...string) string {
	return outer(Grn)(msg, styles)
}

func (c *Color) Yellow(msg interface{}, styles ...string) string {
	return outer(Yel)(msg, styles)
}

func (c *Color) Blue(msg interface{}, styles ...string) string {
	return outer(Blu)(msg, styles)
}

func (c *Color) Magenta(msg interface{}, styles ...string) string {
	return outer(Mgn)(msg, styles)
}

func (c *Color) Cyan(msg interface{}, styles ...string) string {
	return outer(Cyn)(msg, styles)
}

func (c *Color) White(msg interface{}, styles ...string) string {
	return outer(Wht)(msg, styles)
}

func (c *Color) Grey(msg interface{}, styles ...string) string {
	return outer(Gry)(msg, styles)
}

func (c *Color) BlackBg(msg interface{}, styles ...string) string {
	return outer(BlkBg)(msg, styles)
}

func (c *Color) RedBg(msg interface{}, styles ...string) string {
	return outer(RdBg)(msg, styles)
}

func (c *Color) GreenBg(msg interface{}, styles ...string) string {
	return outer(GrnBg)(msg, styles)
}

func (c *Color) YellowBg(msg interface{}, styles ...string) string {
	return outer(YelBg)(msg, styles)
}

func (c *Color) BlueBg(msg interface{}, styles ...string) string {
	return outer(BluBg)(msg, styles)
}

func (c *Color) MagentaBg(msg interface{}, styles ...string) string {
	return outer(MgnBg)(msg, styles)
}

func (c *Color) CyanBg(msg interface{}, styles ...string) string {
	return outer(CynBg)(msg, styles)
}

func (c *Color) WhiteBg(msg interface{}, styles ...string) string {
	return outer(WhtBg)(msg, styles)
}

func (c *Color) Reset(msg interface{}, styles ...string) string {
	return outer(Rst)(msg, styles)
}

func (c *Color) Bold(msg interface{}, styles ...string) string {
	return outer(B)(msg, styles)
}

func (c *Color) Dim(msg interface{}, styles ...string) string {
	return outer(Dm)(msg, styles)
}

func (c *Color) Italic(msg interface{}, styles ...string) string {
	return outer(I)(msg, styles)
}

func (c *Color) Underline(msg interface{}, styles ...string) string {
	return outer(U)(msg, styles)
}

func (c *Color) Inverse(msg interface{}, styles ...string) string {
	return outer(In)(msg, styles)
}

func (c *Color) Hidden(msg interface{}, styles ...string) string {
	return outer(Hd)(msg, styles)
}

func (c *Color) Strikeout(msg interface{}, styles ...string) string {
	return outer(So)(msg, styles)
}

func Black(msg interface{}, styles ...string) string {
	return global.Black(msg, styles...)
}

func Red(msg interface{}, styles ...string) string {
	return global.Red(msg, styles...)
}

func Green(msg interface{}, styles ...string) string {
	return global.Green(msg, styles...)
}

func Yellow(msg interface{}, styles ...string) string {
	return global.Yellow(msg, styles...)
}

func Blue(msg interface{}, styles ...string) string {
	return global.Blue(msg, styles...)
}

func Magenta(msg interface{}, styles ...string) string {
	return global.Magenta(msg, styles...)
}

func Cyan(msg interface{}, styles ...string) string {
	return global.Cyan(msg, styles...)
}

func White(msg interface{}, styles ...string) string {
	return global.White(msg, styles...)
}

func Grey(msg interface{}, styles ...string) string {
	return global.Grey(msg, styles...)
}

func BlackBg(msg interface{}, styles ...string) string {
	return global.BlackBg(msg, styles...)
}

func RedBg(msg interface{}, styles ...string) string {
	return global.RedBg(msg, styles...)
}

func GreenBg(msg interface{}, styles ...string) string {
	return global.GreenBg(msg, styles...)
}

func YellowBg(msg interface{}, styles ...string) string {
	return global.YellowBg(msg, styles...)
}

func BlueBg(msg interface{}, styles ...string) string {
	return global.BlueBg(msg, styles...)
}

func MagentaBg(msg interface{}, styles ...string) string {
	return global.MagentaBg(msg, styles...)
}

func CyanBg(msg interface{}, styles ...string) string {
	return global.CyanBg(msg, styles...)
}

func WhiteBg(msg interface{}, styles ...string) string {
	return global.WhiteBg(msg, styles...)
}

func Reset(msg interface{}, styles ...string) string {
	return global.Reset(msg, styles...)
}

func Bold(msg interface{}, styles ...string) string {
	return global.Bold(msg, styles...)
}

func Dim(msg interface{}, styles ...string) string {
	return global.Dim(msg, styles...)
}

func Italic(msg interface{}, styles ...string) string {
	return global.Italic(msg, styles...)
}

func Underline(msg interface{}, styles ...string) string {
	return global.Underline(msg, styles...)
}

func Inverse(msg interface{}, styles ...string) string {
	return global.Underline(msg, styles...)
}

func Hidden(msg interface{}, styles ...string) string {
	return global.Hidden(msg, styles...)
}

func Strikeout(msg interface{}, styles ...string) string {
	return global.Strikeout(msg, styles...)
}
