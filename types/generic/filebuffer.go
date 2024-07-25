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

// NewFileBuffer creates a new instance of the FileBuffer struct.
//
// It returns a pointer to the newly created FileBuffer.
func NewFileBuffer() *FileBuffer {
	return &FileBuffer{}
}

// Write writes the contents of the byte slice p to the FileBuffer.
//
// It returns the number of bytes written and any error encountered.
// The function first checks if there is any extra space between the current
// offset and the end of the buffer. If so, it writes zeros to fill that space.
// Then it checks if the current offset is within the buffer's length. If so,
// it copies as much of the input as possible to the buffer starting from the
// current offset. The remaining input is then written to the buffer.
// Finally, the function updates the offset and returns the number of bytes
// written and any error encountered.
//
// Parameters:
// - p: a byte slice containing the data to be written to the FileBuffer.
//
// Returns:
// - n: the number of bytes written to the FileBuffer.
// - err: any error encountered during the write operation.
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

// Seek sets the offset for the next read or write on the FileBuffer.
//
// Parameters:
// - offset: the offset in bytes from the origin.
// - whence: the origin from which to calculate the offset. It can be one of:
//   - io.SeekStart: the offset is relative to the start of the buffer.
//   - io.SeekCurrent: the offset is relative to the current position in the buffer.
//   - io.SeekEnd: the offset is relative to the end of the buffer.
//
// Returns:
// - int64: the new offset.
// - error: an error if the new offset is negative.
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

// Reader returns an io.Reader that reads from the FileBuffer.
//
// It returns a bytes.Reader that reads from the bytes of the FileBuffer's Buffer.
// The bytes.Reader is created using the Bytes method of the FileBuffer's Buffer.
//
// Returns:
// - io.Reader: an io.Reader that reads from the FileBuffer's Buffer.
func (fb *FileBuffer) Reader() io.Reader {
	return bytes.NewReader(fb.Buffer.Bytes())
}

// Close closes the FileBuffer.
//
// It returns nil as there are no resources to close.
func (fb *FileBuffer) Close() error {
	return nil
}

// Tell returns the current offset of the FileBuffer.
//
// It returns the current offset of the FileBuffer.
//
// Returns:
// - int: the current offset of the FileBuffer.
func (fb *FileBuffer) Tell() int {
	return fb.Offset
}

// Size returns the size of the FileBuffer.
//
// It calculates the size of the buffer by calling the Len method of the underlying Buffer.
//
// Returns:
// - int: the size of the FileBuffer.
func (fb *FileBuffer) Size() int {
	return fb.Buffer.Len()
}

// Bytes returns the underlying byte slice of the FileBuffer.
//
// It returns a byte slice that represents the contents of the FileBuffer.
//
// Returns:
// - []byte: a byte slice containing the contents of the FileBuffer.
func (fb *FileBuffer) Bytes() []byte {
	return fb.Buffer.Bytes()
}
