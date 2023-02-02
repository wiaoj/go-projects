package consoleColor

import (
	"fmt"
	"strings"
)

type PrintableText interface {
	String() string
}

type Color struct {
	color   int
	style   int
	bgcolor int
	text    []interface{}
}

func (color Color) String() string {
	var ext string
	var parts []string

	if color.bgcolor > 0 {
		parts = append(parts, fmt.Sprintf("%d", color.bgcolor))
	}
	if color.color > 0 {
		parts = append(parts, fmt.Sprintf("%d", color.color))
	}
	if color.style > 0 {
		parts = append(parts, fmt.Sprintf("%d", color.style))
	}

	ext = strings.Join(parts[:], ";")

	returnString := "\u001b[" + ext + "m"

	if len(color.text) > 0 {
		for _, item := range color.text {
			switch f := item.(type) {
			default:
				returnString += fmt.Sprint(f)
			}
		}
		returnString += Reset()
	}

	return returnString
}

func (color Color) Apply(messages ...interface{}) Color {
	color.SetText(messages)
	return color
}

func (color *Color) SetText(messages []interface{}) *Color {
	if len(messages) > 0 {
		if len(color.text) == 0 {
			color.text = messages
		} else {
			color.text = append(color.text, messages...)
		}
	}
	return color
}

// Colors
func (color *Color) Black(messages ...interface{}) *Color {
	color.color = 30
	color.SetText(messages)
	return color
}
func (color *Color) Red(messages ...interface{}) *Color {
	color.color = 31
	color.SetText(messages)
	return color
}
func (color *Color) Green(messages ...interface{}) *Color {
	color.color = 32
	color.SetText(messages)
	return color
}
func (color *Color) Yellow(messages ...interface{}) *Color {
	color.color = 33
	color.SetText(messages)
	return color
}
func (color *Color) Blue(messages ...interface{}) *Color {
	color.color = 34
	color.SetText(messages)
	return color
}
func (color *Color) Magenta(messages ...interface{}) *Color {
	color.color = 35
	color.SetText(messages)
	return color
}
func (color *Color) Cyan(messages ...interface{}) *Color {
	color.color = 36
	color.SetText(messages)
	return color
}
func (color *Color) White(messages ...interface{}) *Color {
	color.color = 37
	color.SetText(messages)
	return color
}

func (color *Color) BlackLight(messages ...interface{}) *Color {
	color.color = 90
	color.SetText(messages)
	return color
}
func (color *Color) RedLight(messages ...interface{}) *Color {
	color.color = 91
	color.SetText(messages)
	return color
}
func (color *Color) GreenLight(messages ...interface{}) *Color {
	color.color = 92
	color.SetText(messages)
	return color
}
func (color *Color) YellowLight(messages ...interface{}) *Color {
	color.color = 93
	color.SetText(messages)
	return color
}
func (color *Color) BlueLight(messages ...interface{}) *Color {
	color.color = 94
	color.SetText(messages)
	return color
}
func (color *Color) MagentaLight(messages ...interface{}) *Color {
	color.color = 95
	color.SetText(messages)
	return color
}
func (color *Color) CyanLight(messages ...interface{}) *Color {
	color.color = 96
	color.SetText(messages)
	return color
}
func (color *Color) WhiteLight(messages ...interface{}) *Color {
	color.color = 97
	color.SetText(messages)
	return color
}

// Text Styles
func (color *Color) Bold(messages ...interface{}) *Color {
	color.style = 1
	color.SetText(messages)
	return color
}
func (color *Color) Dim(messages ...interface{}) *Color {
	color.style = 2
	color.SetText(messages)
	return color
}
func (color *Color) Italic(messages ...interface{}) *Color {
	color.style = 3
	color.SetText(messages)
	return color
}
func (color *Color) Underline(messages ...interface{}) *Color {
	color.style = 4
	color.SetText(messages)
	return color
}
func (color *Color) Inverse(messages ...interface{}) *Color {
	color.style = 7
	color.SetText(messages)
	return color
}
func (color *Color) Hidden(messages ...interface{}) *Color {
	color.style = 8
	color.SetText(messages)
	return color
}
func (color *Color) Strikethrough(messages ...interface{}) *Color {
	color.style = 9
	color.SetText(messages)
	return color
}

// Background Colors
func (color *Color) BgBlack(messages ...interface{}) *Color {
	color.bgcolor = 40
	color.SetText(messages)
	return color
}
func (color *Color) BgRed(messages ...interface{}) *Color {
	color.bgcolor = 41
	color.SetText(messages)
	return color
}
func (color *Color) BgGreen(messages ...interface{}) *Color {
	color.bgcolor = 42
	color.SetText(messages)
	return color
}
func (color *Color) BgYellow(messages ...interface{}) *Color {
	color.bgcolor = 43
	color.SetText(messages)
	return color
}
func (color *Color) BgBlue(messages ...interface{}) *Color {
	color.bgcolor = 44
	color.SetText(messages)
	return color
}
func (color *Color) BgMagenta(messages ...interface{}) *Color {
	color.bgcolor = 45
	color.SetText(messages)
	return color
}
func (color *Color) BgCyan(messages ...interface{}) *Color {
	color.bgcolor = 46
	color.SetText(messages)
	return color
}
func (color *Color) BgWhite(messages ...interface{}) *Color {
	color.bgcolor = 47
	color.SetText(messages)
	return color
}

func (color *Color) BgBlackLight(messages ...interface{}) *Color {
	color.bgcolor = 100
	color.SetText(messages)
	return color
}
func (color *Color) BgRedLight(messages ...interface{}) *Color {
	color.bgcolor = 101
	color.SetText(messages)
	return color
}
func (color *Color) BgGreenLight(messages ...interface{}) *Color {
	color.bgcolor = 102
	color.SetText(messages)
	return color
}
func (color *Color) BgYellowLight(messages ...interface{}) *Color {
	color.bgcolor = 103
	color.SetText(messages)
	return color
}
func (color *Color) BgBlueLight(messages ...interface{}) *Color {
	color.bgcolor = 104
	color.SetText(messages)
	return color
}
func (color *Color) BgMagentaLight(messages ...interface{}) *Color {
	color.bgcolor = 105
	color.SetText(messages)
	return color
}
func (color *Color) BgCyanLight(messages ...interface{}) *Color {
	color.bgcolor = 106
	color.SetText(messages)
	return color
}
func (color *Color) BgWhiteLight(messages ...interface{}) *Color {
	color.bgcolor = 107
	color.SetText(messages)
	return color
}

// For all reset. BackgroundColor and TextColor
func Reset() string {
	return fmt.Sprintf("\033[%dm", 0)
}

func newColor() *Color {
	var color Color
	color.color = 0
	color.style = 0
	color.bgcolor = 0
	return &color
}

var (
	// Text Colors
	Black   = func(messages ...interface{}) *Color { return newColor().Black(messages...) }
	Red     = func(messages ...interface{}) *Color { return newColor().Red(messages...) }
	Green   = func(messages ...interface{}) *Color { return newColor().Green(messages...) }
	Yellow  = func(messages ...interface{}) *Color { return newColor().Yellow(messages...) }
	Blue    = func(messages ...interface{}) *Color { return newColor().Blue(messages...) }
	Magenta = func(messages ...interface{}) *Color { return newColor().Magenta(messages...) }
	Cyan    = func(messages ...interface{}) *Color { return newColor().Cyan(messages...) }
	White   = func(messages ...interface{}) *Color { return newColor().White(messages...) }

	// Text Light Colors
	BlackLight   = func(messages ...interface{}) *Color { return newColor().BlackLight(messages...) }
	RedLight     = func(messages ...interface{}) *Color { return newColor().RedLight(messages...) }
	GreenLight   = func(messages ...interface{}) *Color { return newColor().GreenLight(messages...) }
	YellowLight  = func(messages ...interface{}) *Color { return newColor().YellowLight(messages...) }
	BlueLight    = func(messages ...interface{}) *Color { return newColor().BlueLight(messages...) }
	MagentaLight = func(messages ...interface{}) *Color { return newColor().MagentaLight(messages...) }
	CyanLight    = func(messages ...interface{}) *Color { return newColor().CyanLight(messages...) }
	WhiteLight   = func(messages ...interface{}) *Color { return newColor().WhiteLight(messages...) }

	// Text Styles
	Bold          = func(messages ...interface{}) *Color { return newColor().Bold(messages...) }
	Dim           = func(messages ...interface{}) *Color { return newColor().Dim(messages...) }
	Italic        = func(messages ...interface{}) *Color { return newColor().Italic(messages...) }
	Underline     = func(messages ...interface{}) *Color { return newColor().Underline(messages...) }
	Inverse       = func(messages ...interface{}) *Color { return newColor().Inverse(messages...) }
	Hidden        = func(messages ...interface{}) *Color { return newColor().Hidden(messages...) }
	Strikethrough = func(messages ...interface{}) *Color { return newColor().Strikethrough(messages...) }

	// Background Colors
	BgBlack   = func(messages ...interface{}) *Color { return newColor().BgBlack(messages...) }
	BgRed     = func(messages ...interface{}) *Color { return newColor().BgRed(messages...) }
	BgGreen   = func(messages ...interface{}) *Color { return newColor().BgGreen(messages...) }
	BgYellow  = func(messages ...interface{}) *Color { return newColor().BgYellow(messages...) }
	BgBlue    = func(messages ...interface{}) *Color { return newColor().BgBlue(messages...) }
	BgMagenta = func(messages ...interface{}) *Color { return newColor().BgMagenta(messages...) }
	BgCyan    = func(messages ...interface{}) *Color { return newColor().BgCyan(messages...) }
	BgWhite   = func(messages ...interface{}) *Color { return newColor().BgWhite(messages...) }

	BgBlackLight   = func(messages ...interface{}) *Color { return newColor().BgBlackLight(messages...) }
	BgRedLight     = func(messages ...interface{}) *Color { return newColor().BgRedLight(messages...) }
	BgGreenLight   = func(messages ...interface{}) *Color { return newColor().BgGreenLight(messages...) }
	BgYellowLight  = func(messages ...interface{}) *Color { return newColor().BgYellowLight(messages...) }
	BgBlueLight    = func(messages ...interface{}) *Color { return newColor().BgBlueLight(messages...) }
	BgMagentaLight = func(messages ...interface{}) *Color { return newColor().BgMagentaLight(messages...) }
	BgCyanLight    = func(messages ...interface{}) *Color { return newColor().BgCyanLight(messages...) }
	BgWhiteLight   = func(messages ...interface{}) *Color { return newColor().BgWhiteLight(messages...) }
)
