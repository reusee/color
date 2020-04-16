package color

import "image/color"

var (
	Black   = color.RGBA{0x55, 0x55, 0x55, 0}
	Red     = color.RGBA{0xff, 0x82, 0x72, 0}
	Green   = color.RGBA{0xb4, 0xfa, 0x72, 0}
	Yellow  = color.RGBA{0xfe, 0xfd, 0xc2, 0}
	Blue    = color.RGBA{0xad, 0xd5, 0xfe, 0}
	Magenta = color.RGBA{0xff, 0x8f, 0xfd, 0}
	Cyan    = color.RGBA{0xd0, 0xd1, 0xfe, 0}
	White   = color.RGBA{0xf3, 0xf3, 0xf3, 0}

	LightBlack   = color.RGBA{0x66, 0x66, 0x66, 0}
	LightRed     = color.RGBA{0xff, 0xc4, 0xbd, 0}
	LightGreen   = color.RGBA{0xd6, 0xfc, 0xb9, 0}
	LightYellow  = color.RGBA{0xfe, 0xfd, 0xd5, 0}
	LightBlue    = color.RGBA{0xc1, 0xe3, 0xfe, 0}
	LightMagenta = color.RGBA{0xff, 0xb1, 0xfe, 0}
	LightCyan    = color.RGBA{0xe5, 0xe6, 0xfe, 0}
	LightWhite   = color.RGBA{0xfe, 0xff, 0xff, 0}

	Foreground = White
	Background = Black

	Ansi16 = [...]color.Color{
		0:  Black,
		1:  Red,
		2:  Green,
		3:  Yellow,
		4:  Blue,
		5:  Magenta,
		6:  Cyan,
		7:  White,
		8:  LightBlack,
		9:  LightRed,
		10: LightGreen,
		11: LightYellow,
		12: LightBlue,
		13: LightMagenta,
		14: LightCyan,
		15: LightWhite,
	}
)
