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

// NewMDTiles creates a new MDTiles object with the given data, width, and bits per pixel.
//
// Parameters:
// - data: a byte slice containing the tile data.
// - width: the width of each tile in pixels.
// - bpp: the number of bits per pixel.
//
// Returns:
// - a pointer to the newly created MDTiles object.
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

// FromData converts the given byte slice data into the format required by the MDTiles struct.
//
// Parameters:
// - data: a byte slice containing the data to be converted.
//
// Returns: None.
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

// ReadPixel returns the value of a pixel at the given coordinates (x, y) from the MDTiles object.
//
// Parameters:
// - x: the x-coordinate of the pixel.
// - y: the y-coordinate of the pixel.
//
// Returns:
// - value: the value of the pixel as a byte.
func (tiles *MDTiles) ReadPixel(x, y int) (value byte) {
	tx := (x%8 + ((x / 8) * (tiles.Bpp * 8 * (64 / (tiles.Bpp * 8)))))
	ty := ((y % 8) * 8) + ((y / 8) * (tiles.Width * tiles.Bpp * 8 * (64 / (tiles.Bpp * 8))))
	if tx+ty < len(tiles.Raw) {
		value = tiles.Raw[tx+ty]
	}
	return
}

// ToPNG generates an image.RGBA object from the given MDTiles object and MDPalette.
//
// Parameters:
// - mdpalette: The MDPalette object containing the colors to be used in the generated image.
//
// Returns:
// - img: The generated image.RGBA object.
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
