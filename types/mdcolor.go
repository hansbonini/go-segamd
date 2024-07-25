package types

import "image/color"

type MDColor struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// NewMDColor creates a new instance of MDColor.
//
// It returns a pointer to the newly created MDColor.
func NewMDColor() *MDColor {
	return &MDColor{}
}

// FromValue sets the R, G, B, and A values of an MDColor instance based on a given uint16 value.
//
// Parameters:
// - v: The uint16 value to extract the R, G, B, and A values from.
//
// Return type: None.
func (mdcolor *MDColor) FromValue(v uint16) {
	mdcolor.R = uint8((v & 0x000E) >> 1)
	mdcolor.G = uint8((v & 0x00E0) >> 5)
	mdcolor.B = uint8((v & 0x0E00) >> 9)
	mdcolor.A = 255
}

// ToRGBA converts an MDColor object to an RGBA color object.
//
// It shifts the R, G, and B values of the MDColor object by 5 bits to the left
// and assigns the result to the R, G, and B fields of the RGBA color object.
// The A value of the MDColor object is assigned directly to the A field of the
// RGBA color object.
//
// Returns:
// - An RGBA color object with the converted R, G, B, and A values.
func (mdcolor *MDColor) ToRGBA() color.RGBA {
	return color.RGBA{
		R: uint8(mdcolor.R << 5),
		G: uint8(mdcolor.G << 5),
		B: uint8(mdcolor.B << 5),
		A: uint8(mdcolor.A),
	}
}

// ToValue converts an MDColor object to a uint16 value.
//
// It combines the R, G, and B values of the MDColor object into a single uint16
// value by shifting the R value by 1 bit to the left, the G value by 5 bits to
// the left, and the B value by 9 bits to the left. The A value is not included
// in the resulting uint16 value.
//
// Returns:
// - A uint16 value representing the R, G, and B values of the MDColor object.
func (mdcolor *MDColor) ToValue() uint16 {
	return uint16(mdcolor.R)<<1 | uint16(mdcolor.G)<<5 | uint16(mdcolor.B)<<9
}
