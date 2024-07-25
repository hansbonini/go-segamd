package types

import (
	"bytes"
	"encoding/binary"
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
	}
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
