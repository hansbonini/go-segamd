package generic

type RingBuffer struct {
	Data   []any
	Offset uint
	Size   int
	Fill   any
}

func NewRingBuffer(size int, fill any) *RingBuffer {
	rb := &RingBuffer{
		Data:   make([]any, size),
		Offset: 0,
		Size:   size,
		Fill:   fill,
	}
	rb.FillData(fill, size)
	return rb
}

func (rb *RingBuffer) Push(data any) {
	rb.Data[rb.Offset] = data
	rb.Offset = (rb.Offset + 1) % uint(rb.Size)
}

func (rb *RingBuffer) Pop() any {
	data := rb.Data[rb.Offset]
	rb.Offset = (rb.Offset - 1) % uint(rb.Size)
	return data
}

func (rb *RingBuffer) Set(data any, offset int) {
	rb.Data[uint(offset)%uint(rb.Size)] = data
}

func (rb *RingBuffer) Get(offset int) any {
	return rb.Data[uint(offset)%uint(rb.Size)]
}

func (rb *RingBuffer) FillData(data any, size int) {
	for i := 0; i < size; i++ {
		rb.Data[i] = data
	}
}
