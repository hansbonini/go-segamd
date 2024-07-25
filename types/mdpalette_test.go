package types_test

import (
	"testing"

	"github.com/hansbonini/go-segamd/types"
)

func TestNewMDPalette(t *testing.T) {
	// Test that NewMDPalette returns a non-nil pointer to MDPalette
	palette := types.NewMDPalette([]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF})
	if palette == nil {
		t.Errorf("NewMDPalette returned nil")
	}

	// Test that the returned MDPalette has the correct number of colors
	if len(palette.Colors) != 16 {
		t.Errorf("NewMDPalette returned wrong number of colors: %d", len(palette.Colors))
	}

	// Test that the colors in the palette are correct
	expectedColors := []types.MDColor{
		{R: 0, G: 0, B: 0, A: 0},
		{R: 1, G: 1, B: 1, A: 255},
		{R: 2, G: 2, B: 2, A: 255},
		{R: 3, G: 3, B: 3, A: 255},
		{R: 4, G: 4, B: 4, A: 255},
		{R: 5, G: 5, B: 5, A: 255},
		{R: 6, G: 6, B: 6, A: 255},
		{R: 7, G: 7, B: 7, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
		{R: 0, G: 0, B: 0, A: 255},
	}
	for i, color := range palette.Colors {
		if color != expectedColors[i] {
			t.Errorf("NewMDPalette returned wrong color at index %d: %v", i, color)
		}
	}
}
