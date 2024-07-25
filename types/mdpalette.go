package types

import (
	"bytes"
	"encoding/binary"
)

type MDPalette struct {
	Colors []MDColor
}

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

func (palette *MDPalette) Size() int {
	return len(palette.Colors)
}
