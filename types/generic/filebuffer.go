package generic

import (
	"bytes"
	"errors"
	"io"
)

type FileBuffer struct {
	Buffer bytes.Buffer
	Offset int
}

func NewFileBuffer() *FileBuffer {
	return &FileBuffer{}
}

func (fb *FileBuffer) Write(p []byte) (n int, err error) {
	if extra := fb.Offset - fb.Buffer.Len(); extra > 0 {
		if _, err := fb.Buffer.Write(make([]byte, extra)); err != nil {
			return n, err
		}
	}
	if fb.Offset < fb.Buffer.Len() {
		n = copy(fb.Buffer.Bytes()[fb.Offset:], p)
		p = p[n:]
	}
	if len(p) > 0 {
		var bn int
		bn, err = fb.Buffer.Write(p)
		n += bn
	}
	fb.Offset += n
	return n, err
}

func (fb *FileBuffer) Seek(offset int64, whence int) (int64, error) {
	newPos, offs := 0, int(offset)
	switch whence {
	case io.SeekStart:
		newPos = offs
	case io.SeekCurrent:
		newPos = fb.Offset + offs
	case io.SeekEnd:
		newPos = fb.Buffer.Len() + offs
	}
	if newPos < 0 {
		return 0, errors.New("negative position")
	}
	fb.Offset = newPos
	return int64(newPos), nil
}

func (fb *FileBuffer) Reader() io.Reader {
	return bytes.NewReader(fb.Buffer.Bytes())
}

func (fb *FileBuffer) Close() error {
	return nil
}

func (fb *FileBuffer) Tell() int {
	return fb.Offset
}

func (fb *FileBuffer) Size() int {
	return fb.Buffer.Len()
}

func (fb *FileBuffer) Bytes() []byte {
	return fb.Buffer.Bytes()
}
