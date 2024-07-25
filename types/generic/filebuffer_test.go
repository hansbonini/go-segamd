package generic_test

import (
	"io"
	"testing"

	"github.com/hansbonini/go-segamd/types/generic"
)

func TestFileBuffer(t *testing.T) {
	fb := &generic.FileBuffer{}

	if _, err := fb.Write([]byte{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}
}

func TestFileBuffer_Bytes(t *testing.T) {
	fb := &generic.FileBuffer{}

	if _, err := fb.Write([]byte{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}

	if bytes := fb.Bytes(); len(bytes) != 4 {
		t.Fatal(bytes)
	}
}

func TestFileBuffer_Reader(t *testing.T) {
	fb := &generic.FileBuffer{}

	if _, err := fb.Write([]byte{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}

	if reader := fb.Reader(); reader == nil {
		t.Fatal(reader)
	}
}

func TestFileBuffer_Close(t *testing.T) {
	fb := &generic.FileBuffer{}

	if _, err := fb.Write([]byte{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}

	if err := fb.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestFileBuffer_Size(t *testing.T) {
	fb := &generic.FileBuffer{}

	if _, err := fb.Write([]byte{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}

	if size := fb.Size(); size != 4 {
		t.Fatal(size)
	}
}

func TestFileBuffer_Tell(t *testing.T) {
	fb := &generic.FileBuffer{}

	if _, err := fb.Write([]byte{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}

	if offset := fb.Tell(); offset != 4 {
		t.Fatal(offset)
	}
}

func TestFileBuffer_Seek(t *testing.T) {
	fb := &generic.FileBuffer{}

	if _, err := fb.Write([]byte{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}

	if _, err := fb.Seek(2, io.SeekStart); err != nil {
		t.Fatal(err)
	}
}

func TestFileBuffer_Write(t *testing.T) {
	fb := &generic.FileBuffer{}

	if _, err := fb.Write([]byte{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}
}
