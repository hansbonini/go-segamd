package types

import (
	"bytes"
	"encoding/binary"
	"go-segamd/types/generic"
	"io"

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

func (rom *MDROM) Init() {
	rom.Header.Unmarshal(rom.Data[0x100:])
}

func (rom *MDROM) UpdateHeader() {
	beforeHeader := rom.Data[:0x100]
	afterHeader := rom.Data[0x200:]
	rom.Data = beforeHeader
	rom.Data = append(rom.Data, rom.Header.Marshal()...)
	rom.Data = append(rom.Data, afterHeader...)
}

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

func (rom *MDROM) UpdateChecksum() {
	checksum := 0
	rom.Seek(0x200)
	for i := 0; i < len(rom.Data)-0x200; i++ {
		v, _ := rom.Read16()
		checksum += int(v)
	}
	rom.Header.Checksum = uint16(checksum) & 0xffff
}

func EncodeSJIS(s string) []byte {
	b, _ := japanese.ShiftJIS.NewEncoder().String(s)
	return []byte(b)
}

func DecodeSJIS(sjis []byte) string {
	t := transform.NewReader(bytes.NewReader(sjis), japanese.ShiftJIS.NewDecoder())
	d, _ := io.ReadAll(t)
	return string(d)
}
