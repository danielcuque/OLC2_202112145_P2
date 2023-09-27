package compiler

type Heap struct {
	Pointer int
}

func NewHeap() *Heap {
	return &Heap{
		Pointer: 0,
	}
}

func (h *Heap) GetPointer() int {
	return h.Pointer
}

func (h *Heap) AddPointer() int {
	h.Pointer++
	return h.Pointer
}

func (h *Heap) SubPointer() int {
	h.Pointer--
	return h.Pointer
}
