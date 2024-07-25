package types

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/hansbonini/go-segamd/types/generic"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

type MDPCM struct {
	Channels   int
	SampleRate int
}

func (pcm *MDPCM) ToWAV(data []byte) ([]byte, error) {
	wBuf := generic.NewFileBuffer()
	pBuf := audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: pcm.Channels,
			SampleRate:  pcm.SampleRate,
		},
		SourceBitDepth: 8,
	}
	w := wav.NewEncoder(wBuf, pcm.SampleRate, 8, pcm.Channels, 1)
	buf := bytes.NewBuffer(data)
	for {
		var r int8
		if err := binary.Read(buf, binary.LittleEndian, &r); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		pBuf.Data = append(pBuf.Data, int(r))
	}
	if err := w.Write(&pBuf); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return wBuf.Bytes(), nil
}

func (pcm *MDPCM) FromWAV(data []byte) ([]byte, error) {
	wBuf := wav.NewDecoder(bytes.NewReader(data))
	pBuf := audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: pcm.Channels,
			SampleRate:  pcm.SampleRate,
		},
		Data:           make([]int, 2048),
		SourceBitDepth: 8,
	}
	chunk := make([]byte, 2048)
	buf := new(bytes.Buffer)
	err := wBuf.FwdToPCM()
	if err != nil {
		return nil, err
	}
	n, err := wBuf.PCMBuffer(&pBuf)
	if err != nil {
		return nil, err
	}
	for n > 0 {
		for i := 0; i < n; i++ {
			chunk[i] = byte(pBuf.Data[i])
		}
		_, err = buf.Write(chunk[0:n])
		if err != nil {
			return nil, err
		}
		n, err = wBuf.PCMBuffer(&pBuf)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
