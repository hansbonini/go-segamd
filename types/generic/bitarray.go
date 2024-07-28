package generic

type BitArray interface {
	SetBit(bit int)
	ClearBit(bit int)
	GetBit(bit int) bool
	SetNextBit(reverse bool)
	ClearNextBit(reverse bool)
	GetValue() (value any)
	SetValue(value any)
}

type BitArray8 struct {
	Data [8]int
}

type BitArray16 struct {
	Data [16]int
}

type BitArray32 struct {
	Data [32]int
}

func NewBitArray8() *BitArray8 {
	return &BitArray8{}
}

func NewBitArray16() *BitArray16 {
	return &BitArray16{}
}

func NewBitArray32() *BitArray32 {
	return &BitArray32{}
}

func (a *BitArray8) SetBit(bit int) {
	a.Data[bit%8] = 1
}

func (a *BitArray8) ClearBit(bit int) {
	a.Data[bit%8] = 0
}

func (a *BitArray8) GetBit(bit int) int {
	return a.Data[bit%8]
}

func (a *BitArray8) SetNextBit(reverse bool) {
	if reverse {
		for i := 7; i > 0; i-- {
			a.Data[i] = a.Data[i-1]
		}
		a.Data[0] = 1
	} else {
		for i := 0; i < 7; i++ {
			a.Data[i] = a.Data[i+1]
		}
		a.Data[7] = 1
	}
}

func (a *BitArray8) ClearNextBit(reverse bool) {
	if reverse {
		for i := 7; i > 0; i-- {
			a.Data[i] = a.Data[i-1]
		}
		a.Data[0] = 0
	} else {
		for i := 0; i < 7; i++ {
			a.Data[i] = a.Data[i+1]
		}
		a.Data[7] &= ^(1)
	}
}

func (a *BitArray8) GetValue() (value uint8) {
	for i := 0; i < 8; i++ {
		if a.GetBit(i) == 1 {
			value |= 1 << i
		}
	}
	return
}

func (a *BitArray8) SetValue(value uint8) {
	for i := 0; i < 8; i++ {
		if value&(1<<i) != 0 {
			a.SetBit(i)
		} else {
			a.ClearBit(i)
		}
	}
}

func (a *BitArray16) SetBit(bit int) {
	a.Data[bit%16] = 1
}

func (a *BitArray16) ClearBit(bit int) {
	a.Data[bit%16] = 0
}

func (a *BitArray16) GetBit(bit int) int {
	return a.Data[bit%16]
}

func (a *BitArray16) SetNextBit(reverse bool) {
	if reverse {
		for i := 15; i > 0; i-- {
			a.Data[i] = a.Data[i-1]
		}
		a.Data[0] = 1
	} else {
		for i := 0; i < 15; i++ {
			a.Data[i] = a.Data[i+1]
		}
		a.Data[15] = 1
	}
}

func (a *BitArray16) ClearNextBit(reverse bool) {
	if reverse {
		for i := 15; i > 0; i-- {
			a.Data[i] = a.Data[i-1]
		}
		a.Data[0] = 0
	} else {
		for i := 0; i < 15; i++ {
			a.Data[i] = a.Data[i+1]
		}
		a.Data[15] &= ^(1)
	}
}

func (a *BitArray16) GetValue() (value uint16) {
	for i := 0; i < 16; i++ {
		if a.GetBit(i) == 1 {
			value |= 1 << i
		}
	}
	return
}

func (a *BitArray16) SetValue(value uint16) {
	for i := 0; i < 16; i++ {
		if value&(1<<i) != 0 {
			a.SetBit(i)
		} else {
			a.ClearBit(i)
		}
	}
}

func (a *BitArray32) SetBit(bit int) {
	a.Data[bit%32] = 1
}

func (a *BitArray32) ClearBit(bit int) {
	a.Data[bit%32] = 0
}

func (a *BitArray32) GetBit(bit int) int {
	return a.Data[bit%32]
}

func (a *BitArray32) SetNextBit(reverse bool) {
	if reverse {
		for i := 31; i > 0; i-- {
			a.Data[i] = a.Data[i-1]
		}
		a.Data[0] = 1
	} else {
		for i := 0; i < 31; i++ {
			a.Data[i] = a.Data[i+1]
		}
		a.Data[31] = 1
	}
}

func (a *BitArray32) ClearNextBit(reverse bool) {
	if reverse {
		for i := 31; i > 0; i-- {
			a.Data[i] = a.Data[i-1]
		}
		a.Data[0] = 0
	} else {
		for i := 0; i < 31; i++ {
			a.Data[i] = a.Data[i+1]
		}
		a.Data[31] &= ^(1)
	}
}

func (a *BitArray32) GetValue() (value uint32) {
	for i := 0; i < 32; i++ {
		if a.GetBit(i) == 1 {
			value |= 1 << i
		}
	}
	return
}

func (a *BitArray32) SetValue(value uint32) {
	for i := 0; i < 32; i++ {
		if value&(1<<i) != 0 {
			a.SetBit(i)
		} else {
			a.ClearBit(i)
		}
	}
}
