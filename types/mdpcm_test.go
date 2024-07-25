package types_test

import (
	"testing"

	"github.com/hansbonini/go-segamd/types"
)

func TestToWAV(t *testing.T) {
	pcm := &types.MDPCM{
		Channels:   2,
		SampleRate: 44100,
	}

	// Test case: empty data
	data := []byte{}
	wavData, err := pcm.ToWAV(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(wavData) != 44 {
		t.Errorf("Expected empty WAV data, got %v", wavData)
	}

	// Test case: non-empty data
	data = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	wavData, err = pcm.ToWAV(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(wavData) == 0 {
		t.Errorf("Expected non-empty WAV data, got empty data")
	}
}

func TestFromWAV(t *testing.T) {
	pcm := &types.MDPCM{
		Channels:   1,
		SampleRate: 44100,
	}

	// Test case: invalid WAV data
	data := []byte{
		0x00, 0x01, 0x02,
	}

	_, err := pcm.FromWAV(data)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// Test case: valid WAV data with incorrect number of channels
	data = []byte{
		// WAV header
		0x52, 0x49, 0x46, 0x46, // "RIFF"
		0xFF, 0xFF, 0xFF, 0xFF, // Chunk size
		0x57, 0x41, 0x56, 0x45, // "WAVE"
		0x66, 0x6D, 0x74, 0x20, // "fmt "
		0x10, 0x00, 0x00, 0x00, // Subchunk size
		0x01, 0x00, // Audio format (PCM)
		0x02, 0x00, // Number of channels
		0x44, 0xAC, 0x00, 0x00, // Sample rate
		0x88, 0x58, 0x01, 0x00, // Byte rate
		0x02, 0x00, // Block align
		0x08, 0x00, // Bits per sample
		0x00, 0x00, // Extra params
		0x64, 0x61, 0x74, 0x61, // "data"
		0xFF, 0xFF, 0xFF, 0xFF, // Data size
		// Audio data
		0x01, 0x02, 0x03, 0x04,
		0x05, 0x06, 0x07, 0x08,
	}

	_, err = pcm.FromWAV(data)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// Test case: valid wav data
	data = []byte{
		// WAV header
		0x52, 0x49, 0x46, 0x46, // "RIFF"
		0x24, 0x08, 0x00, 0x00, // Chunk size
		0x57, 0x41, 0x56, 0x45, // "WAVE"
		0x66, 0x6D, 0x74, 0x20, // "fmt "
		0x10, 0x00, 0x00, 0x00, // Subchunk size
		0x01, 0x00, // Audio format (PCM)
		0x01, 0x00, // Number of channels
		0x44, 0xAC, 0x00, 0x00, // Sample rate
		0x44, 0xAC, 0x00, 0x00, // Byte rate
		0x01, 0x00, // Block align
		0x08, 0x00, // Bits per sample
		0x64, 0x61, 0x74, 0x61, // "data"
		0x00, 0x08, 0x00, 0x00, // Data size
	}
	for i := 0; i < 2048; i++ {
		data = append(data, byte(0x80))
	}
	wavData, err := pcm.FromWAV(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(wavData) == 0 {
		t.Errorf("Expected non-empty WAV data, got empty data")
	}
}
