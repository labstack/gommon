# Color
Style terminal text.

### Installation
```go get github.com/labstack/gommon/color```

### Usage
[labstack/gommon/color](https://github.com/labstack/gommon/blob/master/color/color_test.go)
```go
// Colored text
fmt.Println(color.Black("black"))
fmt.Println(color.Red("red"))
fmt.Println(color.Green("green"))
fmt.Println(color.Yellow("yellow"))
fmt.Println(color.Blue("blue"))
fmt.Println(color.Magenta("magenta"))
fmt.Println(color.Cyan("cyan"))
fmt.Println(color.White("white"))
fmt.Println(color.Grey("grey"))
```
![Colored Text](http://i.imgur.com/8RtY1QR.png)

```go
// Colored background
fmt.Println(color.BlackBg("black background", color.Wht))
fmt.Println(color.RedBg("red background"))
fmt.Println(color.GreenBg("green background"))
fmt.Println(color.YellowBg("yellow background"))
fmt.Println(color.BlueBg("blue background"))
fmt.Println(color.MagentaBg("magenta background"))
fmt.Println(color.CyanBg("cyan background"))
fmt.Println(color.WhiteBg("white background"))
```
![Colored Background](http://i.imgur.com/SrrS6lw.png)

```go
// Emphasis
fmt.Println(color.Bold("bold"))
fmt.Println(color.Dim("dim"))
fmt.Println(color.Italic("italic"))
fmt.Println(color.Underline("underline"))
fmt.Println(color.Inverse("inverse"))
fmt.Println(color.Hidden("hidden"))
fmt.Println(color.Strikeout("strikeout"))
```
![Emphasis](http://i.imgur.com/3RSJBbc.png)
```go
// Combo
fmt.Println(Green("bold green with white background", B, WhtBg))
fmt.Println(Red("underline red", U))
fmt.Println(Yellow("dim yellow", D))
fmt.Println(Cyan("inverse cyan", In))
fmt.Println(Blue("bold underline dim blue", B, U, D))
```
![Combo](http://i.imgur.com/jWGq9Ca.png)
