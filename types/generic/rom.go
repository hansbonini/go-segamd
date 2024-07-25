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

func NewROM(filename string) (rom *ROM, err error) {
	rom = &ROM{
		Filename: filename,
	}
	err = rom.Init()
	return rom, err
}

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

func (rom *ROM) Read8() (value uint8, err error) {
	if rom.Offset < rom.Size {
		value = rom.Data[rom.Offset]
		rom.Offset += 1
	} else {
		err = io.EOF
	}
	return
}

func (rom *ROM) Read16() (value uint16, err error) {
	if rom.Offset+1 < rom.Size {
		value = uint16(rom.Data[rom.Offset])<<8 | uint16(rom.Data[rom.Offset+1])
		rom.Offset += 2
	} else {
		err = io.EOF
	}
	return
}

func (rom *ROM) Read32() (value uint32, err error) {
	if rom.Offset+3 < rom.Size {
		value = uint32(rom.Data[rom.Offset])<<24 | uint32(rom.Data[rom.Offset+1])<<16 | uint32(rom.Data[rom.Offset+2])<<8 | uint32(rom.Data[rom.Offset+3])
		rom.Offset += 4
	} else {
		err = io.EOF
	}
	return
}

func (rom *ROM) ReadString() (value string, err error) {
	if rom.Offset+1 < rom.Size {
		value = string(rom.Data[rom.Offset])
		rom.Offset += 1
	} else {
		err = io.EOF
	}
	return
}

func (rom *ROM) Seek(offset int) {
	rom.Offset = offset
}

func (rom *ROM) Tell() (offset int) {
	return rom.Offset
}
