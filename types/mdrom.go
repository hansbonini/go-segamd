package types

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/hansbonini/go-segamd/types/generic"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type MDRawHeader struct {
	Type               [16]byte
	Copyright          [16]byte
	DomesticTitle      [48]byte
	InternationalTitle [48]byte
	SerialNumber       [14]byte
	Checksum           [2]byte
	Devices            [16]byte
	ROMStartAddress    [4]byte
	ROMEndAddress      [4]byte
	RAMStartAddress    [4]byte
	RAMEndAddress      [4]byte
	SRAMType           [4]byte
	SRAMStartAddress   [4]byte
	SRAMEndAddress     [4]byte
	Modem              [12]byte
	Reserved1          [40]byte
	Region             [3]byte
	Reserved2          [13]byte
}

type MDHeader struct {
	Type               string
	Copyright          string
	DomesticTitle      string
	InternationalTitle string
	SerialNumber       string
	Checksum           uint16
	Devices            string
	ROMStartAddress    uint32
	ROMEndAddress      uint32
	RAMStartAddress    uint32
	RAMEndAddress      uint32
	SRAMType           [4]byte
	SRAMStartAddress   uint32
	SRAMEndAddress     uint32
	Modem              string
	Reserved1          [40]byte
	Region             string
	Reserved2          [13]byte
}

type MDROM struct {
	generic.ROM
	Header MDHeader
}

// NewMDROM creates a new MDROM object with the given filename.
//
// Parameters:
// - filename: the name of the file to be opened.
//
// Returns:
// - mdrom: a pointer to the newly created MDrom object.
// - err: an error if the file could not be opened.
func NewMDROM(filename string) (mdrom *MDROM, err error) {
	rom, err := generic.NewROM(filename)
	if err != nil {
		return nil, err
	}
	mdrom = &MDROM{
		Header: MDHeader{},
		ROM:    *rom,
	}
	mdrom.Init()
	return mdrom, err
}

// Init initializes the MDROM object by unmarshaling the data in the specified range into the Header field.
//
// Parameters:
// - None
//
// Returns:
// - None
func (rom *MDROM) Init() {
	rom.Header.Unmarshal(rom.Data[0x100:])
}

// UpdateHeader updates the header of the MDROM object by replacing the existing header with a new one.
//
// It does this by extracting the data before and after the header, then replacing the header with the marshaled version of the Header field.
//
// Parameters:
// - None
//
// Returns:
// - None
func (rom *MDROM) UpdateHeader() {
	beforeHeader := rom.Data[:0x100]
	afterHeader := rom.Data[0x200:]
	rom.Data = beforeHeader
	rom.Data = append(rom.Data, rom.Header.Marshal()...)
	rom.Data = append(rom.Data, afterHeader...)
}

// Unmarshal unmarshals the given byte slice into the MDHeader struct.
//
// It extracts the values from the byte slice and assigns them to the corresponding fields of the MDHeader struct.
// The byte slice is expected to contain the following data in the following order:
// - Type (10 bytes)
// - Copyright (20 bytes)
// - DomesticTitle (30 bytes)
// - InternationalTitle (40 bytes)
// - SerialNumber (14 bytes)
// - Checksum (2 bytes)
// - Devices (20 bytes)
// - ROMStartAddress (4 bytes)
// - ROMEndAddress (4 bytes)
// - RAMStartAddress (4 bytes)
// - RAMEndAddress (4 bytes)
// - SRAMType (4 bytes)
// - SRAMStartAddress (4 bytes)
// - SRAMEndAddress (4 bytes)
// - Modem (32 bytes)
// - Reserved1 (40 bytes)
// - Region (3 bytes)
// - Reserved2 (13 bytes)
//
// Parameters:
// - data: The byte slice containing the data to be unmarshaled.
//
// Returns:
// - None
func (header *MDHeader) Unmarshal(data []byte) {
	header.Type = DecodeSJIS(data[0x0:0x10])
	header.Copyright = DecodeSJIS(data[0x10:0x20])
	header.DomesticTitle = DecodeSJIS(data[0x20:0x50])
	header.InternationalTitle = DecodeSJIS(data[0x50:0x80])
	header.SerialNumber = DecodeSJIS(data[0x80:0x8e])
	header.Checksum = binary.BigEndian.Uint16(data[0x8e:0x90])
	header.Devices = DecodeSJIS(data[0x90:0xa0])
	header.ROMStartAddress = binary.BigEndian.Uint32(data[0xa0:0xa4])
	header.ROMEndAddress = binary.BigEndian.Uint32(data[0xa4:0xa8])
	header.RAMStartAddress = binary.BigEndian.Uint32(data[0xa8:0xac])
	header.RAMEndAddress = binary.BigEndian.Uint32(data[0xac:0xb0])
	copy(header.SRAMType[:], data[0xb0:0xb4])
	header.SRAMStartAddress = binary.BigEndian.Uint32(data[0xb4:0xb8])
	header.SRAMEndAddress = binary.BigEndian.Uint32(data[0xb8:0xbc])
	header.Modem = DecodeSJIS(data[0xbc:0xc8])
	copy(header.Reserved1[:], data[0xc8:0xf0])
	header.Region = DecodeSJIS(data[0xf0:0xf3])
	copy(header.Reserved2[:], data[0xf3:0x100])
}

// Marshal serializes the MDHeader struct into a byte slice.
//
// It writes the values of the MDHeader struct fields into the byte slice using the BigEndian byte order.
// The byte slice contains the following data in the following order:
// - Type (10 bytes)
// - Copyright (20 bytes)
// - DomesticTitle (30 bytes)
// - InternationalTitle (40 bytes)
// - SerialNumber (14 bytes)
// - Checksum (2 bytes)
// - Devices (20 bytes)
// - ROMStartAddress (4 bytes)
// - ROMEndAddress (4 bytes)
// - RAMStartAddress (4 bytes)
// - RAMEndAddress (4 bytes)
// - SRAMType (4 bytes)
// - SRAMStartAddress (4 bytes)
// - SRAMEndAddress (4 bytes)
// - Modem (32 bytes)
// - Reserved1 (40 bytes)
// - Region (3 bytes)
// - Reserved2 (13 bytes)
//
// Returns:
// - []byte: The serialized byte slice.
func (header *MDHeader) Marshal() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, EncodeSJIS(header.Type))
	binary.Write(&buf, binary.BigEndian, EncodeSJIS(header.Copyright))
	binary.Write(&buf, binary.BigEndian, EncodeSJIS(header.DomesticTitle))
	binary.Write(&buf, binary.BigEndian, EncodeSJIS(header.InternationalTitle))
	binary.Write(&buf, binary.BigEndian, EncodeSJIS(header.SerialNumber))
	binary.Write(&buf, binary.BigEndian, header.Checksum)
	binary.Write(&buf, binary.BigEndian, EncodeSJIS(header.Devices))
	binary.Write(&buf, binary.BigEndian, header.ROMStartAddress)
	binary.Write(&buf, binary.BigEndian, header.ROMEndAddress)
	binary.Write(&buf, binary.BigEndian, header.RAMStartAddress)
	binary.Write(&buf, binary.BigEndian, header.RAMEndAddress)
	binary.Write(&buf, binary.BigEndian, header.SRAMType)
	binary.Write(&buf, binary.BigEndian, header.SRAMStartAddress)
	binary.Write(&buf, binary.BigEndian, header.SRAMEndAddress)
	binary.Write(&buf, binary.BigEndian, EncodeSJIS(header.Modem))
	binary.Write(&buf, binary.BigEndian, header.Reserved1)
	binary.Write(&buf, binary.BigEndian, EncodeSJIS(header.Region))
	binary.Write(&buf, binary.BigEndian, header.Reserved2)
	return buf.Bytes()
}

// UpdateChecksum updates the checksum of the MDROM by iterating over the ROM data
// starting from offset 0x200 and summing up the values of each 16-bit word.
// The updated checksum is then stored in the Header.Checksum field.
//
// No parameters.
// No return values.
func (rom *MDROM) UpdateChecksum() {
	checksum := 0
	rom.Seek(0x200)
	for i := 0; i < len(rom.Data)-0x200; i++ {
		v, _ := rom.Read16()
		checksum += int(v)
	}
	rom.Header.Checksum = uint16(checksum) & 0xffff
}

// EncodeSJIS encodes a string into Shift-JIS encoding and returns the byte slice representation.
//
// Parameters:
// - s: The string to be encoded.
//
// Returns:
// - []byte: The byte slice representation of the encoded string.
func EncodeSJIS(s string) []byte {
	b, _ := japanese.ShiftJIS.NewEncoder().String(s)
	return []byte(b)
}

// DecodeSJIS decodes a byte slice containing Shift-JIS encoded data into a string.
//
// Parameters:
// - sjis: The byte slice containing the Shift-JIS encoded data.
//
// Returns:
// - string: The decoded string.
func DecodeSJIS(sjis []byte) string {
	t := transform.NewReader(bytes.NewReader(sjis), japanese.ShiftJIS.NewDecoder())
	d, _ := io.ReadAll(t)
	return string(d)
}
