package color

import "image/color"

var (
	Black      = color.RGBA{0x33, 0x33, 0x33, 0}
	LightBlack = color.RGBA{0x44, 0x44, 0x44, 0}
	White      = color.RGBA{0xf0, 0xf0, 0xf0, 0}
	LightWhite = color.RGBA{0xff, 0xff, 0xff, 0}

	// dark background
	n0           uint8 = 0xff
	n1           uint8 = 0xe0
	n2           uint8 = 0xf0
	Red                = color.RGBA{n0, n1, n1, 0}
	LightRed           = color.RGBA{n0, n2, n2, 0}
	Green              = color.RGBA{n1, n0, n1, 0}
	LightGreen         = color.RGBA{n2, n0, n2, 0}
	Yellow             = color.RGBA{n0, n0, n1, 0}
	LightYellow        = color.RGBA{n0, n0, n2, 0}
	Blue               = color.RGBA{n1, n1, n0, 0}
	LightBlue          = color.RGBA{n2, n2, n0, 0}
	Magenta            = color.RGBA{n0, n1, n0, 0}
	LightMagenta       = color.RGBA{n0, n2, n0, 0}
	Cyan               = color.RGBA{n1, n0, n0, 0}
	LightCyan          = color.RGBA{n2, n0, n0, 0}

	// light background
	//n0           uint8 = 0x00
	//n1           uint8 = 0x50
	//n2           uint8 = 0x60
	//Red                = color.RGBA{n1, n0, n0, 0}
	//LightRed           = color.RGBA{n2, n0, n0, 0}
	//Green              = color.RGBA{n0, n1, n0, 0}
	//LightGreen         = color.RGBA{n0, n2, n0, 0}
	//Yellow             = color.RGBA{n1, n1, n0, 0}
	//LightYellow        = color.RGBA{n2, n2, n0, 0}
	//Blue               = color.RGBA{n0, n0, n1, 0}
	//LightBlue          = color.RGBA{n0, n0, n2, 0}
	//Magenta            = color.RGBA{n1, n0, n1, 0}
	//LightMagenta       = color.RGBA{n2, n0, n2, 0}
	//Cyan               = color.RGBA{n0, n1, n1, 0}
	//LightCyan          = color.RGBA{n0, n2, n2, 0}

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
