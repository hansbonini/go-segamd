package types

import (
	"bytes"
	"encoding/binary"
)

type MDPalette struct {
	Colors []MDColor
}

// NewMDPalette creates a new MDPalette from the provided byte slice.
//
// The data parameter is a byte slice containing the raw data for the palette.
// The function reads 16 color values from the data using big-endian byte order,
// creates a new MDColor for each color, and appends it to the palette's Colors slice.
// The first color's alpha value is set to 0.
// The function returns a pointer to the newly created MDPalette.
//
// Parameters:
// - data: a byte slice containing the raw data for the palette.
//
// Returns:
// - palette: a pointer to the newly created MDPalette.
func NewMDPalette(data []byte) (palette *MDPalette) {
	palette = &MDPalette{}
	buf := bytes.NewBuffer(data)
	for i := 0; i < 16; i++ {
		var rawcolor uint16
		color := NewMDColor()
		if err := binary.Read(buf, binary.BigEndian, &rawcolor); err != nil {
			rawcolor = 0
		}
		color.FromValue(rawcolor)
		if i == 0 {
			color.A = 0
		}
		palette.Colors = append(palette.Colors, *color)
	}
	return palette
}

// Size returns the number of colors in the MDPalette.
//
// It takes no parameters.
// It returns an integer representing the number of colors in the palette.
func (palette *MDPalette) Size() int {
	return len(palette.Colors)
}
