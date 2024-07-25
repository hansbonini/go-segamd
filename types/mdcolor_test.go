package types_test

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/hansbonini/go-segamd/types"
)

func TestNewMDColor(t *testing.T) {
	tests := []struct {
		name  string
		value uint16
		want  *types.MDColor
	}{
		{
			name:  "Test with max values",
			value: 0xFFFF,
			want:  &types.MDColor{R: 7, G: 7, B: 7, A: 255},
		},
		{
			name:  "Test with min values",
			value: 0x2222,
			want:  &types.MDColor{R: 1, G: 1, B: 1, A: 255},
		},
		{
			name:  "Test with zero values",
			value: 0x0000,
			want:  &types.MDColor{R: 0, G: 0, B: 0, A: 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := types.NewMDColor()
			got.FromValue(tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMDColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMDColor_FromValue(t *testing.T) {
	tests := []struct {
		name  string
		input uint16
		want  *types.MDColor
	}{
		{
			name:  "Test 1",
			input: 0x0000,
			want:  &types.MDColor{R: 0, G: 0, B: 0, A: 255},
		},
		{
			name:  "Test 2",
			input: 0x000E,
			want:  &types.MDColor{R: 7, G: 0, B: 0, A: 255},
		},
		{
			name:  "Test 3",
			input: 0x00E0,
			want:  &types.MDColor{R: 0, G: 7, B: 0, A: 255},
		},
		{
			name:  "Test 4",
			input: 0x0E00,
			want:  &types.MDColor{R: 0, G: 0, B: 7, A: 255},
		},
		{
			name:  "Test 5",
			input: 0x000F,
			want:  &types.MDColor{R: 7, G: 0, B: 0, A: 255},
		},
		{
			name:  "Test 6",
			input: 0x00EF,
			want:  &types.MDColor{R: 7, G: 7, B: 0, A: 255},
		},
		{
			name:  "Test 7",
			input: 0x0EF0,
			want:  &types.MDColor{R: 0, G: 7, B: 7, A: 255},
		},
		{
			name:  "Test 8",
			input: 0x0EFF,
			want:  &types.MDColor{R: 7, G: 7, B: 7, A: 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mdcolor := &types.MDColor{}
			mdcolor.FromValue(tt.input)
			if !reflect.DeepEqual(mdcolor, tt.want) {
				t.Errorf("FromValue() = %v, want %v", mdcolor, tt.want)
			}
		})
	}
}

func TestMDColor_ToRGBA(t *testing.T) {
	tests := []struct {
		name    string
		mdcolor types.MDColor
		want    color.RGBA
	}{
		{
			name: "Test with max values",
			mdcolor: types.MDColor{
				R: 0x1F,
				G: 0x1F,
				B: 0x1F,
				A: 0xFF,
			},
			want: color.RGBA{
				R: 224,
				G: 224,
				B: 224,
				A: 255,
			},
		},
		{
			name: "Test with min values",
			mdcolor: types.MDColor{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			},
			want: color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 0,
			},
		},
		{
			name: "Test with middle values",
			mdcolor: types.MDColor{
				R: 0x10,
				G: 0x10,
				B: 0x10,
				A: 0x80,
			},
			want: color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 128,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mdcolor.ToRGBA(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MDColor.ToRGBA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMDColor_ToValue(t *testing.T) {
	tests := []struct {
		name    string
		mdcolor types.MDColor
		want    uint16
	}{
		{
			name: "Test with max values",
			mdcolor: types.MDColor{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			},
			want: 0xFFFE,
		},
		{
			name: "Test with min values",
			mdcolor: types.MDColor{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0x00,
			},
			want: 0x0000,
		},
		{
			name: "Test with middle values",
			mdcolor: types.MDColor{
				R: 15,
				G: 15,
				B: 15,
				A: 255,
			},
			want: 0x1FFE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mdcolor.ToValue(); got != tt.want {
				t.Errorf("MDColor.ToValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
