package generic

type RingBuffer struct {
	Data   []any
	Cursor int
	Size   int
	Mask   int
	Fill   any
}

func NewRingBuffer(size int, fill any) *RingBuffer {
	rb := &RingBuffer{
		Data:   make([]any, size),
		Cursor: 0,
		Size:   size,
		Mask:   size - 1,
		Fill:   fill,
	}
	rb.FillData(fill, size)
	return rb
}

func (rb *RingBuffer) Push(data any) {
	rb.Data[rb.Cursor] = data
	rb.Cursor = (rb.Cursor + 1) & rb.Mask
}

func (rb *RingBuffer) Pop() any {
	data := rb.Data[rb.Cursor]
	rb.Cursor = (rb.Cursor - 1) & rb.Mask
	return data
}

func (rb *RingBuffer) Set(data any, cursor int) {
	rb.Cursor = cursor & rb.Mask
	rb.Data[rb.Cursor] = data
	rb.Cursor = (rb.Cursor + 1) & rb.Mask
}

func (rb *RingBuffer) Get(cursor int) any {
	return rb.Data[cursor&rb.Mask]
}

func (rb *RingBuffer) FillData(data any, size int) {
	for i := 0; i < size; i++ {
		rb.Data[i] = data
	}
}
