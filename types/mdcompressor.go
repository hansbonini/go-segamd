package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"slices"

	"github.com/hansbonini/go-segamd/types/generic"
)

type MDCompressor interface {
	Marshal() []byte
	Unmarshal() []byte
}

type MDCompressor_SEGARD struct {
	ROM generic.ROM
}

type MDCompressor_NEMESIS struct {
	ROM generic.ROM
}

type MDCompressor_KOZINSKI struct {
	ROM generic.ROM
}

type MDCompressor_ENIGMA struct {
	ROM generic.ROM
}

type MDCompressor_SAXMAN struct {
	ROM generic.ROM
}

type MDCompressor_STI struct {
	ROM generic.ROM
}

type MDCompressor_STI2 struct {
	ROM generic.ROM
}

type MDCompressor_WESTONE struct {
	ROM generic.ROM
}

type MDCompressor_SILICONSYNAPSE struct {
	ROM generic.ROM
}

type MDCompressor_NAMCO struct {
	ROM generic.ROM
}

type MDCompressor_TECHNOSOFT struct {
	ROM generic.ROM
}

type MDCompressor_KONAMI1 struct {
	ROM generic.ROM
}

type MDCompressor_KONAMI2 struct {
	ROM generic.ROM
}

type MDCompressor_KONAMI3 struct {
	ROM generic.ROM
}

type MDCompressor_TOSE struct {
	ROM generic.ROM
}

type MDCompressor_EASTRIKE struct {
	ROM generic.ROM
}

type MDCompressor_NEXTECH struct {
	ROM generic.ROM
}

type MDCompressor_WOLFTEAM struct {
	ROM generic.ROM
}

type MDCompressor_ANCIENT struct {
	ROM generic.ROM
}

type MDCompressor_SOFTWARECREATIONS struct {
	ROM generic.ROM
}

type MDCompressor_KOEI struct {
	ROM generic.ROM
}

type MDCompressor_FACTOR5 struct {
	ROM generic.ROM
}

type MDCompressor_TECMO struct {
	ROM generic.ROM
}

type MDCompressor_SNK struct {
	ROM generic.ROM
}

type MDCompressor_ITL struct {
	ROM generic.ROM
}

// NewMDCompressor creates a new instance of MDCompressor based on the given algorithm and ROM.
//
// Parameters:
// - algorithm: a string representing the algorithm to use for compression.
// - rom: a generic.ROM object representing the ROM data.
//
// Returns:
// - MDCompressor: a pointer to the newly created MDCompressor object, or nil if the algorithm is not recognized.
func NewMDCompressor(algorithm string, rom generic.ROM) MDCompressor {
	switch algorithm {
	case "SEGARD":
		return &MDCompressor_SEGARD{
			ROM: rom,
		}
	case "NEMESIS":
		return &MDCompressor_NEMESIS{
			ROM: rom,
		}
	case "KOZINSKI":
		return &MDCompressor_KOZINSKI{
			ROM: rom,
		}
	case "ENIGMA":
		return &MDCompressor_ENIGMA{
			ROM: rom,
		}
	case "SAXMAN":
		return &MDCompressor_SAXMAN{
			ROM: rom,
		}
	case "STI":
		return &MDCompressor_STI{
			ROM: rom,
		}
	case "STI2":
		return &MDCompressor_STI2{
			ROM: rom,
		}
	case "WESTONE":
		return &MDCompressor_WESTONE{
			ROM: rom,
		}
	case "SILICONSYNAPSE":
		return &MDCompressor_SILICONSYNAPSE{
			ROM: rom,
		}
	case "NAMCO":
		return &MDCompressor_NAMCO{
			ROM: rom,
		}
	case "TECHNOSOFT":
		return &MDCompressor_TECHNOSOFT{
			ROM: rom,
		}
	case "KONAMI1":
		return &MDCompressor_KONAMI1{
			ROM: rom,
		}
	case "KONAMI2":
		return &MDCompressor_KONAMI2{
			ROM: rom,
		}
	case "KONAMI3":
		return &MDCompressor_KONAMI3{
			ROM: rom,
		}
	case "TOSE":
		return &MDCompressor_TOSE{
			ROM: rom,
		}
	case "EASTRIKE":
		return &MDCompressor_EASTRIKE{
			ROM: rom,
		}
	case "NEXTECH":
		return &MDCompressor_NEXTECH{
			ROM: rom,
		}
	case "WOLFTEAM":
		return &MDCompressor_WOLFTEAM{
			ROM: rom,
		}
	case "ANCIENT":
		return &MDCompressor_ANCIENT{
			ROM: rom,
		}
	case "SOFTWARECREATIONS":
		return &MDCompressor_SOFTWARECREATIONS{
			ROM: rom,
		}
	case "KOEI":
		return &MDCompressor_KOEI{
			ROM: rom,
		}
	case "FACTOR5":
		return &MDCompressor_FACTOR5{
			ROM: rom,
		}
	case "TECMO":
		return &MDCompressor_TECMO{
			ROM: rom,
		}
	case "SNK":
		return &MDCompressor_SNK{
			ROM: rom,
		}
	case "ITL":
		return &MDCompressor_ITL{
			ROM: rom,
		}
	}

	fmt.Printf("Unknown algorithm: %s\n", algorithm)
	return nil
}

// Marshal compresses the ROM data using the SEGARD compression algorithm and returns the compressed data as a byte slice.
//
// It reads the ROM data in chunks of 0x20 bytes and performs the following steps for each chunk:
// - Counts the occurrences of each byte in the chunk and stores them in the 'stats' map.
// - Identifies the candidates for compression based on the number of occurrences and stores them in the 'candidates' map.
// - Calculates the masks for each candidate byte and stores them in the 'masks' map.
// - Determines the order of occurrence for the candidate bytes and stores it in the 'ocurrenceOrder' byte slice.
// - Constructs a compression chain by appending the number of candidates, followed by the candidate bytes and their corresponding masks.
// - Applies the non-repeating mask to the chunk based on the 'nonrepeat' value.
// - Writes the compression chain to the output buffer.
//
// After processing all the chunks, it appends two 0xFF bytes to the output buffer to ensure even length.
// Finally, it returns the compressed data as a byte slice.
//
// Returns:
// - []byte: the compressed data as a byte slice.
func (segard *MDCompressor_SEGARD) Marshal() []byte {
	var err error
	data := bytes.NewBuffer(segard.ROM.Data)
	out := new(bytes.Buffer)
	chunk := make([]byte, 0x20)
	for {
		stats := make(map[byte]int)
		candidates := make(map[byte]int)
		masks := make(map[byte]uint32)
		if err = binary.Read(data, binary.BigEndian, &chunk); err != nil {
			break
		}
		segard.getRepeats(chunk, stats, candidates)
		segard.getMasks(chunk, candidates, masks)
		ocurrenceOrder := segard.getOcurrenceOrder(chunk, candidates)
		chain := make([]byte, 0)
		chain = append(chain, byte(len(candidates)))
		nonrepeat := uint32(0)
		for cv := range ocurrenceOrder {
			if v, ok := masks[ocurrenceOrder[cv]]; ok {
				chain = append(chain, ocurrenceOrder[cv])
				chain = binary.BigEndian.AppendUint32(chain, v)
				nonrepeat |= v
			}
		}
		if nonrepeat != 0xFFFFFFFF {
			for k, v := range chunk {
				bit := ((nonrepeat << k) >> 31) & 0x01
				if bit == 0 {
					chain = append(chain, v)
				}
			}
		}
		if err = binary.Write(out, binary.BigEndian, chain); err != nil {
			break
		}
	}
	if err = binary.Write(out, binary.BigEndian, uint8(0xFF)); err != nil {
		log.Fatal(err)
	}
	if len(out.Bytes())%2 == 0 {
		if err = binary.Write(out, binary.BigEndian, uint8(0xFF)); err != nil {
			log.Fatal(err)
		}
	}
	return out.Bytes()
}

// getOcurrenceOrder returns the occurrence order of candidate bytes in the given chunk.
//
// Parameters:
// - chunk: a byte slice representing the chunk to search for candidate bytes.
// - candidates: a map of byte to int representing the candidate bytes and their occurrences.
//
// Returns:
// - []byte: a byte slice representing the occurrence order of candidate bytes.
func (segard *MDCompressor_SEGARD) getOcurrenceOrder(chunk []byte, candidates map[byte]int) []byte {
	ocurrenceOrder := make([]byte, 0)
	for _, v := range chunk {
		for cv := range candidates {
			if v == cv {
				if !slices.Contains(ocurrenceOrder, v) {
					ocurrenceOrder = append(ocurrenceOrder, v)
				}
			}
		}
		if len(ocurrenceOrder) == len(candidates) {
			break
		}
	}
	return ocurrenceOrder
}

// getMasks calculates the masks for each candidate byte in the given chunk.
//
// Parameters:
// - chunk: a byte slice representing the chunk to calculate masks for.
// - candidates: a map of byte to int representing the candidate bytes and their occurrences.
// - masks: a map of byte to uint32 representing the masks for each candidate byte.
func (segard *MDCompressor_SEGARD) getMasks(chunk []byte, candidates map[byte]int, masks map[byte]uint32) {
	for _, v := range chunk {
		for cv := range candidates {
			masks[cv] <<= 1
			if v == cv {
				masks[cv] |= 1
			} else {
				masks[cv] |= 0
			}
		}
	}
}

// getRepeats calculates the number of occurrences of each byte in the given chunk and stores the results in the stats map.
// It also identifies candidate bytes that occur more than 5 times and stores them in the candidates map.
//
// Parameters:
// - chunk: a byte slice representing the chunk to search for byte occurrences.
// - stats: a map of byte to int representing the number of occurrences of each byte.
// - candidates: a map of byte to int representing the candidate bytes and their occurrences.
func (segard *MDCompressor_SEGARD) getRepeats(chunk []byte, stats map[byte]int, candidates map[byte]int) {
	for _, repeats := range chunk {
		stats[repeats]++
	}
	for k, v := range stats {
		if v > 5 {
			candidates[k] = v
		}
	}
}

// Unmarshal decodes the SEGARD compression format from the ROM and returns the decompressed data.
//
// It reads the ROM byte by byte, decoding the compressed data. The format consists of a series of chunks,
// each chunk containing a variable number of repeated bytes. The chunks are terminated by a FF byte.
// Each chunk starts with a byte indicating the number of repeated bytes in the chunk.
// After that, for each repeated byte, there is a byte indicating the value of the repeated byte,
// followed by a 32-bit mask indicating which bits of the repeated byte are set.
// The mask is used to determine which bits of the repeated byte are set in each occurrence of the byte.
// The mask is constructed by shifting the bits of the repeated byte to the left and setting the bits
// according to the mask.
// If the mask is not all ones, there are additional bytes in the chunk that are not repeated.
// These bytes are read from the ROM and stored in the chunk.
// The chunks are concatenated into the decompressed data.
//
// Returns:
// - []byte: the decompressed data.
func (segard *MDCompressor_SEGARD) Unmarshal() []byte {
	var repeats uint8
	var err error
	buffer := new(bytes.Buffer)
	chunk := make([]byte, 0x20)
	if repeats, err = segard.ROM.Read8(); err != nil {
		log.Fatal(err)
	}
	for repeats != uint8(0xFF) {
		var pattern uint32
		for x := uint8(0); x < repeats; x++ {
			var value uint8
			var mask uint32
			if value, err = segard.ROM.Read8(); err != nil {
				break
			}
			if mask, err = segard.ROM.Read32(); err != nil {
				break
			}
			pattern |= mask
			i := 0
			for y := 0x1F; y >= 0; y-- {
				bit := (mask >> y) & 0x01
				if bit == 0x01 {
					chunk[i] = value
				}
				i++
			}
		}
		i := 0
		if pattern != 0xFFFFFFFF {
			for x := 0x1F; x >= 0; x-- {
				bit := (pattern >> x) & 0x01
				if bit == 0 {
					if chunk[i], err = segard.ROM.Read8(); err != nil {
						break
					}
				}
				i++
			}
		}
		buffer.Write(chunk)
		if repeats, err = segard.ROM.Read8(); err != nil {
			break
		}
	}
	return buffer.Bytes()
}

func (nemesis *MDCompressor_NEMESIS) Marshal() []byte {
	return []byte{}
}

func (nemesis *MDCompressor_NEMESIS) Unmarshal() []byte {
	return []byte{}
}

func (kozinski *MDCompressor_KOZINSKI) Marshal() []byte {
	return []byte{}
}

func (kozinski *MDCompressor_KOZINSKI) Unmarshal() []byte {
	return []byte{}
}

func (enigma *MDCompressor_ENIGMA) Marshal() []byte {
	return []byte{}
}

func (enigma *MDCompressor_ENIGMA) Unmarshal() []byte {
	return []byte{}
}

func (saxman *MDCompressor_SAXMAN) Marshal() []byte {
	return []byte{}
}

func (saxman *MDCompressor_SAXMAN) Unmarshal() []byte {
	return []byte{}
}

func (sti *MDCompressor_STI) Marshal() []byte {
	return []byte{}
}

func (sti *MDCompressor_STI) Unmarshal() []byte {
	return []byte{}
}

func (sti2 *MDCompressor_STI2) Marshal() []byte {
	return []byte{}
}

func (sti2 *MDCompressor_STI2) Unmarshal() []byte {
	return []byte{}
}

func (westone *MDCompressor_WESTONE) Marshal() []byte {
	return []byte{}
}

func (westone *MDCompressor_WESTONE) Unmarshal() []byte {
	return []byte{}
}

func (siliconsynapse *MDCompressor_SILICONSYNAPSE) Marshal() []byte {
	return []byte{}
}

func (siliconsynapse *MDCompressor_SILICONSYNAPSE) Unmarshal() []byte {
	return []byte{}
}

func (namco *MDCompressor_NAMCO) Marshal() []byte {
	var buffer, temp bytes.Buffer
	var encoded int
	minLength := 3
	maxLength := (1 << 4) + minLength
	bitCount := 0
	bitFlag := generic.NewBitArray8()
	windowSize := 0x1000
	window := generic.NewRingBuffer(windowSize, byte(0x00))
	window.Offset = 0xFEE

	buffer.WriteByte(byte(namco.ROM.Size >> 8))
	buffer.WriteByte(byte(namco.ROM.Size & 0xFF))

	for namco.ROM.Offset < namco.ROM.Size {
		if bitCount > 7 {
			buffer.WriteByte(bitFlag.GetValue())
			for _, v := range temp.Bytes() {
				buffer.WriteByte(v)
			}
			temp.Reset()
			bitCount = 0
			bitFlag = generic.NewBitArray8()
		}
		offset, length := namco.FindMatch(window, minLength, maxLength)
		if length >= minLength {
			bitFlag.ClearBit(bitCount)
			lzpair := uint16((offset << 8) | ((offset >> 4) & 0xF0))
			lzpair &= 0xFFF0
			lzpair |= uint16(length - minLength)
			temp.WriteByte(uint8(lzpair >> 8))
			temp.WriteByte(uint8(lzpair & 0xFF))
			for i := 0; i < length; i++ {
				window.Push(namco.ROM.Data[namco.ROM.Offset])
				namco.ROM.Offset++
				encoded++
			}
		} else {
			bitFlag.SetBit(bitCount)
			temp.WriteByte(namco.ROM.Data[namco.ROM.Offset])
			window.Push(namco.ROM.Data[namco.ROM.Offset])
			namco.ROM.Offset++
		}
		bitCount++
	}
	if bitCount > 0 {
		buffer.WriteByte(bitFlag.GetValue())
		for _, v := range temp.Bytes() {
			buffer.WriteByte(v)
		}
	}
	return buffer.Bytes()
}

func (namco *MDCompressor_NAMCO) FindMatch(window *generic.RingBuffer, minLength int, maxLength int) (offset int, length int) {
	if namco.ROM.Offset+minLength >= namco.ROM.Size {
		return
	}

	pos := int((window.Offset) & uint(window.Size-1))
	for i := 0; i <= (window.Size + maxLength - 1); i++ {
		size := 0
		for ; size+namco.ROM.Offset < namco.ROM.Size; size++ {
			wo := pos - i + size
			if wo >= pos || (window.Get(wo) != namco.ROM.Data[namco.ROM.Offset+size]) {
				break
			}
			if size >= maxLength-1 {
				break
			}
		}
		if size >= length {
			length = size
			offset = pos - i
		}
	}

	return
}

// Unmarshal decodes the compressed data stored in the ROM using the Lempel-Ziv algorithm
// and returns the uncompressed data as a byte slice.
//
// Window has a fixed size of 4096 (0x1000) bytes.
// Window starts to write at position 4078 (0xFEE).
// Byte pair is composed by length and offset in the window
// Length is stored in the least significant 4 bits
// Offset is stored in the most significant 12 bits
//
// # OOOOOOOO OOOOLLLL
//
// Function read the first 2 bytes from the ROM which determine the uncompressed size.
// Next byte, will be read bit a bit to determine which bytes from the sequence should
// be read as a single byte or a LZ pair (composed by length and offset).
// If bit is 1, the next byte will be read and added to the window
// If bit is 0, the LZ pair will be parsed in length and offset. Then the sequence will
// be read from the window according to the offset, added to the buffer and added to the
// window at current position.
// After the entire sequence has been read, the buffer will be returned.
//
// Parameters:
//
// - None
//
// Returns:
//
// - []byte: The uncompressed data as a byte slice.
func (namco *MDCompressor_NAMCO) Unmarshal() []byte {
	var buffer bytes.Buffer
	var decoded int
	var b uint8
	var err error
	window := generic.NewRingBuffer(0x1000, uint8(0x00))
	window.Offset = 0xFEE
	uncompressedSize, err := namco.ROM.Read16()
	if err != nil {
		log.Fatal(err)
	}
	b, err = namco.ROM.Read8()
	if err != nil {
		log.Fatal(err)
	}
	for decoded < int(uncompressedSize) {
		pattern := generic.BitArray8{}
		pattern.SetValue(b)
		for i := 0; i < 8; i++ {
			if pattern.GetBit(i) == 1 {
				if b, err = namco.ROM.Read8(); err != nil {
					break
				}
				buffer.WriteByte(b)
				window.Push(b)
				decoded += 1
			} else {
				var r uint16
				if r, err = namco.ROM.Read16(); err != nil {
					break
				}
				length := (r & 0x0F) + 3
				offset := ((r & 0xF0) << 4) | (r >> 8)
				for j := 0; j < int(length); j++ {
					buffer.WriteByte(window.Get(int(offset) + j).(byte))
					window.Push(window.Get(int(offset) + j).(byte))
				}
				decoded += int(length)
			}
		}
		if b, err = namco.ROM.Read8(); err != nil {
			break
		}
	}
	return buffer.Bytes()
}

func (technosoft *MDCompressor_TECHNOSOFT) Marshal() []byte {
	c := NewMDCompressor("NAMCO", technosoft.ROM)
	return c.Marshal()
}

// Exactly the same algorithm as NAMCO
// Unmarshal decompresses the ROM data using the NAMCO algorithm
// and returns the decompressed data.
//
// Parameters:
//
// - None
//
// Returns:
//
// - []byte: The uncompressed data as a byte slice.
func (technosoft *MDCompressor_TECHNOSOFT) Unmarshal() []byte {
	c := NewMDCompressor("NAMCO", technosoft.ROM)
	return c.Unmarshal()
}

func (konami1 *MDCompressor_KONAMI1) Marshal() []byte {
	return []byte{}
}

func (konami1 *MDCompressor_KONAMI1) Unmarshal() []byte {
	return []byte{}
}

func (konami2 *MDCompressor_KONAMI2) Marshal() []byte {
	return []byte{}
}

func (konami2 *MDCompressor_KONAMI2) Unmarshal() []byte {
	return []byte{}
}

func (konami3 *MDCompressor_KONAMI3) Marshal() []byte {
	return []byte{}
}

func (konami3 *MDCompressor_KONAMI3) Unmarshal() []byte {
	return []byte{}
}

func (tose *MDCompressor_TOSE) Marshal() []byte {
	return []byte{}
}

func (tose *MDCompressor_TOSE) Unmarshal() []byte {
	return []byte{}
}

func (eastrike *MDCompressor_EASTRIKE) Marshal() []byte {
	return []byte{}
}

func (eastrike *MDCompressor_EASTRIKE) Unmarshal() []byte {
	return []byte{}
}

func (nextech *MDCompressor_NEXTECH) Marshal() []byte {
	return []byte{}
}

func (nextech *MDCompressor_NEXTECH) Unmarshal() []byte {
	return []byte{}
}

func (wolfteam *MDCompressor_WOLFTEAM) Marshal() []byte {
	return []byte{}
}

func (wolfteam *MDCompressor_WOLFTEAM) Unmarshal() []byte {
	return []byte{}
}

func (ancient *MDCompressor_ANCIENT) Marshal() []byte {
	return []byte{}
}

func (ancient *MDCompressor_ANCIENT) Unmarshal() []byte {
	return []byte{}
}

func (softwarecreations *MDCompressor_SOFTWARECREATIONS) Marshal() []byte {
	return []byte{}
}

func (softwarecreations *MDCompressor_SOFTWARECREATIONS) Unmarshal() []byte {
	return []byte{}
}

func (koei *MDCompressor_KOEI) Marshal() []byte {
	return []byte{}
}

func (koei *MDCompressor_KOEI) Unmarshal() []byte {
	return []byte{}
}

func (factor5 *MDCompressor_FACTOR5) Marshal() []byte {
	return []byte{}
}

func (factor5 *MDCompressor_FACTOR5) Unmarshal() []byte {
	return []byte{}
}

func (tecmo *MDCompressor_TECMO) Marshal() []byte {
	return []byte{}
}

func (tecmo *MDCompressor_TECMO) Unmarshal() []byte {
	return []byte{}
}

func (snk *MDCompressor_SNK) Marshal() []byte {
	return []byte{}
}

func (snk *MDCompressor_SNK) Unmarshal() []byte {
	return []byte{}
}

func (itl *MDCompressor_ITL) Marshal() []byte {
	return []byte{}
}

func (itl *MDCompressor_ITL) Unmarshal() []byte {
	return []byte{}
}
