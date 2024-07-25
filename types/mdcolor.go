package types

import "image/color"

type MDColor struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewMDColor() *MDColor {
	return &MDColor{}
}

func (mdcolor *MDColor) FromValue(v uint16) {
	mdcolor.R = uint8((v & 0x000E) >> 1)
	mdcolor.G = uint8((v & 0x00E0) >> 5)
	mdcolor.B = uint8((v & 0x0E00) >> 9)
	mdcolor.A = 255
}

func (mdcolor *MDColor) ToRGBA() color.RGBA {
	return color.RGBA{
		R: uint8(mdcolor.R << 5),
		G: uint8(mdcolor.G << 5),
		B: uint8(mdcolor.B << 5),
		A: uint8(mdcolor.A),
	}
}

func (mdcolor *MDColor) ToValue() uint16 {
	return uint16(mdcolor.R<<1) | uint16(mdcolor.G<<5) | uint16(mdcolor.B<<9)
}
