package main

type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f uint8
	h uint8
	l uint8
}

func (r *Registers) getBC() uint16 {
	return uint16(r.b)<<8 | uint16(r.c)
}
func (r *Registers) setBC(value uint16) {
	r.b = uint8((value & 0xFF00) >> 8)
	r.c = uint8((value & 0xFF) >> 8)
}

func main() {
}
