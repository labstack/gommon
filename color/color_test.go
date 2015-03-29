package color

import (
	"fmt"
	"testing"

	"github.com/labstack/gommon/color"
)

func TestText(t *testing.T) {
	fmt.Println("*** colored text ***")
	fmt.Println(color.Black("black"))
	fmt.Println(color.Red("red"))
	fmt.Println(color.Green("green"))
	fmt.Println(color.Yellow("yellow"))
	fmt.Println(color.Blue("blue"))
	fmt.Println(color.Magenta("magenta"))
	fmt.Println(color.Cyan("cyan"))
	fmt.Println(color.White("white"))
	fmt.Println(color.Grey("grey"))
}

func TestBackground(t *testing.T) {
	fmt.Println("*** colored background ***")
	fmt.Println(color.BlackBg("black background", color.Wht))
	fmt.Println(color.RedBg("red background"))
	fmt.Println(color.GreenBg("green background"))
	fmt.Println(color.YellowBg("yellow background"))
	fmt.Println(color.BlueBg("blue background"))
	fmt.Println(color.MagentaBg("magenta background"))
	fmt.Println(color.CyanBg("cyan background"))
	fmt.Println(color.WhiteBg("white background"))
}

func TestEmphasis(t *testing.T) {
	fmt.Println("*** emphasis ***")
	fmt.Println(color.Bold("bold"))
	fmt.Println(color.Dim("dim"))
	fmt.Println(color.Italic("italic"))
	fmt.Println(color.Underline("underline"))
	fmt.Println(color.Inverse("inverse"))
	fmt.Println(color.Hidden("hidden"))
	fmt.Println(color.Strikeout("strikeout"))
}

func TestMixMatch(t *testing.T) {
	fmt.Println("*** mix and match ***")
	fmt.Println(Green("bold green with white background", B, WhtBg))
	fmt.Println(Red("underline red", U))
	fmt.Println(Yellow("dim yellow", D))
	fmt.Println(Cyan("inverse cyan", In))
	fmt.Println(Blue("bold underline dim blue", B, U, D))
}
