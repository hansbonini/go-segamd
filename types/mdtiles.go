package types

import (
	"image"
)

type MDTiles struct {
	Raw    []byte
	Width  int
	Height int
	Bpp    int
}

func NewMDTiles(data []byte, width int, bpp int) *MDTiles {
	tiles := &MDTiles{
		Width: width,
		Bpp:   bpp,
	}
	min := 0
	if len(data)%(tiles.Width*tiles.Bpp*8) > 0 {
		min = 1
	}
	tiles.Height = min + len(data)/(tiles.Width*tiles.Bpp*8)
	tiles.FromData(data)
	return tiles
}

func (tiles *MDTiles) FromData(data []byte) {
	switch tiles.Bpp {
	case 1:
		tiles.Raw = make([]byte, len(data)*8)
		for k, v := range data {
			tiles.Raw[k*8] = v >> 7
			tiles.Raw[k*8+1] = (v >> 6) & 0x1
			tiles.Raw[k*8+2] = (v >> 5) & 0x1
			tiles.Raw[k*8+3] = (v >> 4) & 0x1
			tiles.Raw[k*8+4] = (v >> 3) & 0x1
			tiles.Raw[k*8+5] = (v >> 2) & 0x1
			tiles.Raw[k*8+6] = (v >> 1) & 0x1
			tiles.Raw[k*8+7] = v & 0x1
		}
	case 2:
		tiles.Raw = make([]byte, len(data)*4)
		for k, v := range data {
			tiles.Raw[k*4] = v >> 6
			tiles.Raw[k*4+1] = (v >> 4) & 0x3
			tiles.Raw[k*4+2] = (v >> 2) & 0x3
			tiles.Raw[k*4+3] = v & 0x3
		}
	case 4:
		tiles.Raw = make([]byte, len(data)*2)
		for k, v := range data {
			tiles.Raw[k*2] = v >> 4
			tiles.Raw[k*2+1] = v & 0xF
		}
	default:
		tiles.Raw = data
	}
}

func (tiles *MDTiles) ReadPixel(x, y int) (value byte) {
	tx := (x%8 + ((x / 8) * (tiles.Bpp * 8 * (64 / (tiles.Bpp * 8)))))
	ty := ((y % 8) * 8) + ((y / 8) * (tiles.Width * tiles.Bpp * 8 * (64 / (tiles.Bpp * 8))))
	if tx+ty < len(tiles.Raw) {
		value = tiles.Raw[tx+ty]
	}
	return
}

func (tiles *MDTiles) ToPNG(mdpalette MDPalette) (img *image.RGBA) {
	rect := image.Rect(0, 0, tiles.Width*8, tiles.Height*8)
	img = image.NewRGBA(rect)

	for y := 0; y < tiles.Height*8; y++ {
		for x := 0; x < tiles.Width*8; x++ {
			pixel := tiles.ReadPixel(x, y)
			rgba := mdpalette.Colors[pixel].ToRGBA()
			img.Set(x, y, rgba)
		}
	}
	return
}
