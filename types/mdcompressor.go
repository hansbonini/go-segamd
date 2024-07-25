package types

import (
	"bytes"
	"encoding/binary"
	"go-segamd/types/generic"
	"log"
	"slices"
)

type MDCompressor interface {
	Marshal() []byte
	Unmarshal() []byte
}

type MDCompressor_SEGARD struct {
	ROM generic.ROM
}

func NewMDCompressor(algorithm string, rom generic.ROM) MDCompressor {
	switch algorithm {
	case "SEGARD":
		return &MDCompressor_SEGARD{
			ROM: rom,
		}
	}
	return nil
}

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
		for _, repeats := range chunk {
			stats[repeats]++
		}
		for k, v := range stats {
			if v > 5 {
				candidates[k] = v
			}
		}
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
