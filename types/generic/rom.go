package generic

import (
	"fmt"
	"io"
	"os"
)

type ROM struct {
	Filename string
	Data     []byte
	Size     int
	Offset   int
}

// NewROM creates a new ROM object with the given filename.
//
// Parameters:
// - filename: the name of the file to be opened.
//
// Returns:
// - rom: a pointer to the newly created ROM object.
// - err: an error if the file could not be opened.
func NewROM(filename string) (rom *ROM, err error) {
	rom = &ROM{
		Filename: filename,
	}
	err = rom.Init()
	return rom, err
}

// Init initializes the ROM object by opening the ROM file, getting the file info,
// reading the ROM data, and storing it in the object.
//
// Returns an error if any of the operations fail.
func (rom *ROM) Init() error {
	// Open ROM file
	file, err := os.Open(rom.Filename)
	if err != nil {
		return fmt.Errorf("unable to open rom: %s", rom.Filename)
	}
	defer file.Close()

	// Get file info from ROM file
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("unable to get rom size: %s", rom.Filename)
	}
	rom.Size = int(info.Size())

	// Read ROM data
	data := make([]byte, rom.Size)
	_, err = file.Read(data)
	if err != nil {
		return fmt.Errorf("unable to read data from rom: %s", rom.Filename)
	}
	rom.Data = data

	return nil
}

// Read8 reads an 8-bit value from the ROM.
//
// It returns the read value and an error if the end of the ROM is reached.
func (rom *ROM) Read8() (value uint8, err error) {
	if rom.Offset < rom.Size {
		value = rom.Data[rom.Offset]
		rom.Offset += 1
	} else {
		err = io.EOF
	}
	return
}

// Read16 reads a 16-bit value from the ROM.
//
// It returns the read value and an error if the end of the ROM is reached.
//
// Parameters:
// - rom: a pointer to a ROM object.
//
// Returns:
// - value: a uint16 representing the read value.
// - err: an error if the end of the ROM is reached.
func (rom *ROM) Read16() (value uint16, err error) {
	if rom.Offset+1 < rom.Size {
		value = uint16(rom.Data[rom.Offset])<<8 | uint16(rom.Data[rom.Offset+1])
		rom.Offset += 2
	} else {
		err = io.EOF
	}
	return
}

// Read32 reads a 32-bit value from the ROM.
//
// It returns the read value and an error if the end of the ROM is reached.
//
// Parameters:
// - rom: a pointer to a ROM object.
//
// Returns:
// - value: a uint32 representing the read value.
// - err: an error if the end of the ROM is reached.
func (rom *ROM) Read32() (value uint32, err error) {
	if rom.Offset+3 < rom.Size {
		value = uint32(rom.Data[rom.Offset])<<24 | uint32(rom.Data[rom.Offset+1])<<16 | uint32(rom.Data[rom.Offset+2])<<8 | uint32(rom.Data[rom.Offset+3])
		rom.Offset += 4
	} else {
		err = io.EOF
	}
	return
}

// ReadString reads a string from the ROM starting at the current offset.
//
// It returns the read string and an error if the end of the ROM is reached.
//
// Parameters:
// - rom: a pointer to a ROM object.
//
// Returns:
// - value: a string representing the read value.
// - err: an error if the end of the ROM is reached.
func (rom *ROM) ReadString() (value string, err error) {
	if rom.Offset+1 < rom.Size {
		value = string(rom.Data[rom.Offset])
		rom.Offset += 1
	} else {
		err = io.EOF
	}
	return
}

// Seek sets the offset for the next read or write on the ROM.
//
// Parameters:
// - offset: the offset in bytes from the start of the ROM.
//
// Returns:
// - None.
func (rom *ROM) Seek(offset int) {
	rom.Offset = offset
}

// Tell returns the current offset of the ROM.
//
// No parameters.
// Returns an integer representing the current offset.
func (rom *ROM) Tell() (offset int) {
	return rom.Offset
}
